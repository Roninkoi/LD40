package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
)

type World struct {
	ticks float64

	physSys PhysSys

	player Entity

	currentLevel int

	level1 Level1
	level2 Level2
	level3 Level3
	level4 Level4
}

func (w *World) draw(r *Renderer) {
	w.level1.draw(r)
	w.level2.draw(r)
	w.level3.draw(r)
	w.level4.draw(r)
}

func (w *World) tick() {
	w.physSys.ticks = w.ticks
	w.physSys.update()

	w.player.obj.mesh.bsc = w.player.obj.phys.pos
	w.player.obj.mesh.bsr = 0.8

	w.level1.tick(w.ticks)
	w.level2.tick(w.ticks)
	w.level3.tick(w.ticks)
	w.level4.tick(w.ticks)
}

func (w *World) loadWorld(gl *webgl.Context) {
	w.physSys.gravity = mgl32.Vec3{0.0, -0.005, 0.0}
	w.physSys.clearPhysObjs()

	w.player.obj.phys.init()
	w.player.obj.hasH = false
	w.player.obj.si = true

	w.level1.load(gl)
	w.level2.load(gl)
	w.level3.load(gl)
	w.level4.load(gl)

	w.switchLevel(1)
}

func (w *World) switchLevel(l int) {
	w.physSys.clearPhysObjs()

	w.level1.stop()
	w.level2.stop()
	w.level3.stop()
	w.level4.stop()

	if l == 1 {
		w.level1.start(w)
	}
	if l == 2 {
		w.level2.start(w)
	}
	if l == 3 {
		w.level3.start(w)
	}
	if l == 4 {
		w.level4.start(w)
	}
}
