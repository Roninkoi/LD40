package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
)

type Obj struct {
	mesh Mesh
	phys Phys
}

func (o *Obj) loadObj(gl *webgl.Context, p string, t string) {
	o.mesh.loadMesh(gl, p, t)
	o.phys.init()
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
}
