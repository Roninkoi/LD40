package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Phys struct {
	pos mgl32.Vec3
	rot mgl32.Vec3
	s   mgl32.Vec3

	v mgl32.Vec3
	a mgl32.Vec3

	isStatic  bool
}

func (p *Phys) init() {
	p.pos = mgl32.Vec3{0.0, 0.0, 0.0}
	p.rot = mgl32.Vec3{0.0, 0.0, 0.0}
	p.s = mgl32.Vec3{1.0, 1.0, 1.0}

	p.v = mgl32.Vec3{0.0, 0.0, 0.0}
	p.a = mgl32.Vec3{0.0, 0.0, 0.0}
}

func (p *Phys) update() {
	p.pos = p.pos.Add(p.v)
}

type PhysSys struct {
	objs []*Obj

	gravity mgl32.Vec3

	ticks float64
}

type Hull struct {
	mesh Mesh

	cn mgl32.Vec3 // normal center
	no bool // normals outside?
}

func (h *Hull) loadHull(p string, no bool) {
	h.mesh.loadMesh(nil, p, "nil")
	h.no = no
}

func (h *Hull) update(m *Mesh) {
	h.mesh.um = m.um
	h.mesh.update()
}

func (h *Hull) intersects(i *Mesh) bool {
	returns := false

	cni := 0

	for j := 0; j < (int)((float64)(len(h.mesh.faceNC))/3.0); j++ {
		for k := 0; k < (int)((float64)(len(i.faceNC))/3.0); k++ {
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
			}
		}
	}

	if returns {
		h.cn = h.cn.Mul(1.0/(float32)(cni))
	}

	return returns
}

func (p *PhysSys) physIsect(i int, j int) {
	if p.objs[i].hull.intersects(&p.objs[j].mesh) {
		collisionNormal := p.objs[i].hull.cn
		collisionNormal = collisionNormal.Mul(p.objs[i].phys.v.Len() + p.objs[j].phys.v.Len())

		if !p.objs[i].phys.isStatic {
			p.objs[i].phys.v = p.objs[i].phys.v.Add(collisionNormal)
		}
		if !p.objs[j].phys.isStatic {
			p.objs[j].phys.v = p.objs[j].phys.v.Sub(collisionNormal)
		}
	}
}

func (p *PhysSys) update() {
	for i := 0; i < len(p.objs); i++ {
		for j := i + 1; j < len(p.objs); j++ {
			if p.objs[i].mesh.bsi(&p.objs[j].mesh) &&
				(!p.objs[i].phys.isStatic || !p.objs[j].phys.isStatic) {
				if p.objs[i].hasH {
					p.physIsect(i, j)
				} else if p.objs[j].hasH {
					p.physIsect(j, i)
				}
			}
		}
	}
	for i := 0; i < len(p.objs); i++ {
		if !p.objs[i].phys.isStatic {
			p.objs[i].phys.v = p.objs[i].phys.v.Add(p.gravity)
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
