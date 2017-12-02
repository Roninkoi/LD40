package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type World struct {
	ticks float64
	camPos mgl32.Vec3

	physSys PhysSys

	player Entity

	objs []Obj

	testmesh Obj
	testmesh2 Obj
	room Obj
}

func (w *World) draw(r *Renderer) {
	w.testmesh.draw(r)
	//w.testmesh2.draw(r)
	w.room.draw(r)

	for i := 0; i < 4 && w.objs != nil; i++ {
		w.objs[i].draw(r)
	}
}

func (w *World) tick() {
	w.physSys.ticks = w.ticks
	w.physSys.update()

	//w.testmesh.phys.pos = mgl32.Vec3{(float32)(math.Sin(w.ticks/100.0))*3.0, 0.0, (float32)(math.Cos(w.ticks/100.0))*7.0}
	//w.testmesh.update()
	w.room.phys.rot[0] = 0.5*(float32)(math.Sin(w.ticks/100.0))
	w.room.update()
}

func (w *World) loadWorld(gl *webgl.Context) {
	w.room.loadObjH(gl, "gfx/models/boxi.obj", "0", false, "gfx/checker.png")
	//w.room.phys.s = mgl32.Vec3{1.0, 1.0, 1.0}
	w.room.phys.isStatic = true

	w.testmesh.loadObjH(gl, "gfx/models/ico.obj", "0", true, "gfx/test.png")

	w.testmesh2.loadObjH(gl, "gfx/models/ico.obj", "0", true, "gfx/test.png")

	w.testmesh.phys.pos[2] = -2.0
	w.testmesh2.phys.pos[2] = -3.0
	w.testmesh2.phys.pos[1] = 2.0
	w.testmesh.phys.v[0] = 0.02

	w.physSys.gravity = mgl32.Vec3{0.0, -0.005, 0.0}
	w.physSys.clearPhysObjs()
	w.physSys.addPhysObj(&w.room)
	w.physSys.addPhysObj(&w.testmesh)

	w.objs = nil
/*
	for i := 0; i < 4; i++ {
		w.objs = append(w.objs, Obj{})
		w.objs[i].loadObjH(gl, "gfx/models/ico.obj", "0", true, "gfx/test.png")

		w.testmesh.phys.pos[2] = -3.0 + (float32)(i)

		w.physSys.addPhysObj(&w.objs[i])
	}*/
	//w.physSys.addPhysObj(&w.testmesh2)
}
