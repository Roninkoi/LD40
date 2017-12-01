package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
)

type World struct {
	ticks float64

	testmesh Mesh
}

func (w *World) draw(r *Renderer) {
	//for i := 0; i < 10000; i++ {
		w.testmesh.draw(r)
	//}
}

func (w *World) tick() {
	//w.testmesh.um = mgl32.Translate3D(3.0, 0.0, 0.0)
	w.testmesh.um = mgl32.Ident4()
	w.testmesh.um = mgl32.HomogRotate3DY(0.01)
	w.testmesh.update()
}

func (w *World) loadWorld(gl *webgl.Context) {
	w.testmesh.loadMesh(gl,"gfx/models/dragon.obj", "gfx/dragon.png")
}
