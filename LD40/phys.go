package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Phys struct {
	pos mgl32.Vec3
	rot mgl32.Vec3
	s   mgl32.Vec3

	rPos mgl32.Vec3

	v mgl32.Vec3
	a mgl32.Vec3

	isStatic bool
}

func (p *Phys) init() {
	p.pos = mgl32.Vec3{0.0, 0.0, 0.0}
	p.rot = mgl32.Vec3{0.0, 0.0, 0.0}
	p.s = mgl32.Vec3{1.0, 1.0, 1.0}

	p.rPos = mgl32.Vec3{0.0, 0.0, 0.0}

	p.v = mgl32.Vec3{0.0, 0.0, 0.0}
	p.a = mgl32.Vec3{0.0, 0.0, 0.0}
}

func (p *Phys) update() {
	p.pos = p.pos.Add(p.v)

	np := mgl32.Vec3{p.pos[0], p.pos[1], p.pos[2]}
	np = np.Mul(0.15)
	p.rPos = p.rPos.Mul(0.85)
	p.rPos = p.rPos.Add(np)
}

type PhysSys struct {
	objs []*Obj

	gravity mgl32.Vec3

	ticks float64
}

type Hull struct {
	mesh Mesh

	cn mgl32.Vec3 // normal center
	no bool       // normals outside?
}

func (h *Hull) loadHull(p string, no bool) {
	h.mesh.loadMesh(nil, p, "nil")
	h.no = no
}

func (h *Hull) update(u *mgl32.Mat4) {
	h.mesh.um = *u
	h.mesh.update()
	h.mesh.getTriSize()
}

// general intersection (slow)
func (h *Hull) intersects(i *Mesh) bool {
	returns := false

	cni := 0

	cnold := h.cn
	h.cn = mgl32.Vec3{0.0, 0.0, 0.0}

	for j := 0; j < (int)((float64)(len(h.mesh.faceNC))/3.0) && !returns; j++ {
		for k := 0; k < (int)((float64)(len(i.faceNC))/3.0) && !returns; k++ {
			a := mgl32.Vec3{h.mesh.faceNC[j*3+0], h.mesh.faceNC[j*3+1], h.mesh.faceNC[j*3+2]}
			b := mgl32.Vec3{i.faceNC[k*3+0], i.faceNC[k*3+1], i.faceNC[k*3+2]}

			c := a.Sub(b)
			c = c.Normalize()

			n := mgl32.Vec3{h.mesh.faceNormals[j*3+0], h.mesh.faceNormals[j*3+1], h.mesh.faceNormals[j*3+2]}

			cdn := c.Dot(n)

			is := cdn < 0.0

			if h.no {
				is = cdn > 0.0
			}

			if is {
				returns = true
				cni += 1

				h.cn = h.cn.Add(n)
				h.cn = h.cn.Mul(a.Sub(b).Len()*0.02 + 1.0)
			}
		}
	}

	if returns {
		h.cn = h.cn.Mul(1.0 / (float32)(cni))
	}
	if h.cn.Len() == 0.0 {
		h.cn = cnold
	}

	return returns
}

// sphere intersection
func (h *Hull) intersectsS(sc mgl32.Vec3, sr float32) bool {
	returns := false

	cni := 0

	h.cn = mgl32.Vec3{0.0, 0.0, 0.0}

	for j := 0; j < (int)((float64)(len(h.mesh.faceNC))/3.0); j++ {
		a := mgl32.Vec3{h.mesh.faceNC[j*3+0], h.mesh.faceNC[j*3+1], h.mesh.faceNC[j*3+2]}
		b := mgl32.Vec3{sc[0], sc[1], sc[2]}

		n := mgl32.Vec3{h.mesh.faceNormals[j*3+0], h.mesh.faceNormals[j*3+1], h.mesh.faceNormals[j*3+2]}

		b = b.Add(n.Mul(sr))

		c := a.Sub(b)

		if c.Len() < h.mesh.tri_size*0.5 {
			cdn := c.Dot(n)

			if cdn < 0.0 {
				returns = true
				cni += 1

				h.cn = h.cn.Add(n)
			}
		}
	}

	if returns {
		h.cn = h.cn.Mul(1.0/(float32)(cni))
	}

	return returns
}

func (p *PhysSys) physIsect(i int, j int) {
	isects := false

	if p.objs[j].si {
		isects = p.objs[i].hull.intersectsS(p.objs[j].mesh.bsc, p.objs[j].mesh.bsr)
	} else {
		isects = p.objs[i].hull.intersects(&p.objs[j].mesh)
	}

	if isects {
		p.objs[i].isects = true
		p.objs[j].isects = true

		nv := p.objs[i].hull.cn

		v := p.objs[i].phys.v.Add(p.objs[j].phys.v)
		/*
				// this is nicer but can be slow
				vecProj := mgl32.Ident4()

				// projection matrix
				vecProj.SetRow(0, mgl32.Vec4{nv[0] * nv[0], nv[0] * nv[1], nv[0] * nv[2], 0.0})
				vecProj.SetRow(1, mgl32.Vec4{nv[1] * nv[0], nv[1] * nv[1], nv[1] * nv[2], 0.0})
				vecProj.SetRow(2, mgl32.Vec4{nv[2] * nv[0], nv[2] * nv[1], nv[2] * nv[2], 0.0})

				v = mgl32.TransformCoordinate(v, vecProj)
		*/
		v = nv.Mul(v.Len()*0.95 + 0.001)

		if !p.objs[i].phys.isStatic {
			p.objs[i].phys.pos = p.objs[i].phys.pos.Add(v)
			p.objs[i].phys.v = p.objs[i].phys.v.Add(v)
		}
		if !p.objs[j].phys.isStatic {
			p.objs[j].phys.pos = p.objs[j].phys.pos.Sub(v)
			p.objs[j].phys.v = p.objs[j].phys.v.Sub(v)
		}
	}
}

var epsilon float32 = 0.006 //0.006

func (p *PhysSys) update() {

	for i := 0; i < len(p.objs); i++ {
		p.objs[i].isects = false
		for j := i + 1; j < len(p.objs); j++ {
			if !p.objs[i].phys.isStatic || !p.objs[j].phys.isStatic {
				if p.objs[i].mesh.bsi(&p.objs[j].mesh) {
					if p.objs[i].phys.v.Len() > epsilon || p.objs[j].phys.v.Len() > epsilon {
						if p.objs[i].hasH {
							p.physIsect(i, j)
						} else if p.objs[j].hasH {
							p.physIsect(j, i)
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(p.objs); i++ {
		if !p.objs[i].phys.isStatic && p.objs[i].phys.v.Len() > epsilon {
			if p.objs[i].phys.v.Len() <= epsilon {
				p.objs[i].phys.v = mgl32.Vec3{0.0, 0.0, 0.0}
			}

			p.objs[i].phys.v = p.objs[i].phys.v.Add(p.gravity)
			p.objs[i].phys.v = p.objs[i].phys.v.Mul(0.97)
			p.objs[i].update()
		}
	}
}

func (p *PhysSys) clearPhysObjs() {
	p.objs = nil
}

func (p *PhysSys) addPhysObj(o *Obj) {
	p.objs = append(p.objs, o)
}
