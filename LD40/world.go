package main

import (
	"github.com/gopherjs/webgl"
	_ "github.com/go-gl/mathgl/mgl32"
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

type World struct {
	ticks float64

	physSys PhysSys

	testmesh Obj
	room Obj
}

func (w *World) draw(r *Renderer) {
	w.testmesh.draw(r)
	w.room.draw(r)
}

func (w *World) tick() {
	w.physSys.ticks = w.ticks
	w.physSys.update()

	w.testmesh.phys.pos[0] = (float32)(math.Sin(w.ticks/50.0)*5.0)
	w.testmesh.phys.rot[1] = (float32)(w.ticks)/10.0
	w.testmesh.phys.rot[0] = (float32)(w.ticks)/100.0
	w.testmesh.phys.s = mgl32.Vec3{3.0, 3.0, 3.0}
	w.testmesh.update()
}

func (w *World) loadWorld(gl *webgl.Context) {
	w.room.loadObj(gl, "gfx/models/boxi.obj", "gfx/checker.png")
	//w.room.phys.s = mgl32.Vec3{1.0, 1.0, 1.0}
	w.room.update()

	w.testmesh.loadObj(gl, "gfx/models/ico.obj", "gfx/test.png")

	w.physSys.clearPhysObjs()
	w.physSys.addPhysObj(&w.testmesh)
	w.physSys.addPhysObj(&w.room)
}
