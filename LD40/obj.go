package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
)

type Obj struct {
	mesh Mesh
	hull Hull
	hasH bool

	phys Phys
}

func (o *Obj) loadObj(gl *webgl.Context, p string, t string) {
	o.mesh.loadMesh(gl, p, t)

	o.hasH = false

	o.phys.init()

	o.update()
}

func (o *Obj) loadObjH(gl *webgl.Context, p string, h string, no bool, t string) {
	o.mesh.loadMesh(gl, p, t)

	o.hasH = false

	if h == "0" {
		o.hull.loadHull(p, no)
		o.hasH = true
	} else {
		if h != "nil" {
			o.hull.loadHull(h, no)
			o.hasH = true
		}
	}

	o.phys.init()

	o.update()
}

func (o *Obj) draw(r *Renderer) {
	o.mesh.draw(r)
}

func (o *Obj) render(r *Renderer) {
	o.mesh.render(r)
}

func (o *Obj) update() {
	o.phys.update()

	o.mesh.um = mgl32.Ident4()

	o.mesh.um = o.mesh.um.Mul4(mgl32.Translate3D(o.phys.pos[0], o.phys.pos[1], o.phys.pos[2]))

	o.mesh.um = o.mesh.um.Mul4(mgl32.Scale3D(o.phys.s[0], o.phys.s[1], o.phys.s[2]))

	o.mesh.um = o.mesh.um.Mul4(mgl32.HomogRotate3DX(o.phys.rot[0]))
	o.mesh.um = o.mesh.um.Mul4(mgl32.HomogRotate3DY(o.phys.rot[1]))
	o.mesh.um = o.mesh.um.Mul4(mgl32.HomogRotate3DZ(o.phys.rot[2]))

	o.mesh.update()

	if o.hasH {
		o.hull.update(&o.mesh)
	}
}
