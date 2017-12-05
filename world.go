package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type World struct {
	ticks float64

	physSys PhysSys

	player Entity

	currentLevel int

	level1 Level
	level2 Level

	seconds int
	timer int
	timelimit int

	rage float64
	ragelimit float64

	score int
}

func (w *World) draw(r *Renderer) {
	w.level1.draw(r)
	w.level2.draw(r)
}

func (w *World) tick() {
	w.physSys.ticks = w.ticks
	w.physSys.update()

	w.player.obj.mesh.bsc = w.player.obj.phys.pos
	w.player.obj.mesh.bsr = 0.8
	w.player.ticks += 1

	w.level1.rage = w.rage
	w.level2.rage = w.rage

	w.level1.tick(w.ticks)
	w.level2.tick(w.ticks)

	w.rage = w.ragelimit - (float64)(w.player.coins*4) - (float64)(w.player.gems*8) - (float64)(w.player.beetles*16) + (float64)(w.seconds)*0.5

	if w.rage < 0 {
		w.rage = 0
	}
	if w.rage >= w.ragelimit {
		w.rage = w.ragelimit
	}
	w.player.spd = 0.05*(float32)(math.Sqrt(w.rage/100.0)) // make sure not < 0

	w.timer = w.timelimit - w.seconds
	if w.timer < 0 {
		w.timer = 0
	}
}

func (w *World) restart() {
	w.player.health = 100.0
	w.player.coins = 0
	w.player.gems = 0
	w.player.beetles = 0
	w.player.obj.phys.pos = mgl32.Vec3{0.0, 0.0, 0.0}
	w.player.obj.phys.rot = mgl32.Vec3{0.0, 0.0, 0.0}

	w.switchLevel(w.currentLevel)
}

func (w *World) loadWorld(gl *webgl.Context) {
	w.physSys.gravity = mgl32.Vec3{0.0, -0.005, 0.0}
	w.physSys.clearPhysObjs()

	w.player.obj.phys.init()
	w.player.obj.hasH = false
	w.player.obj.si = true
	w.player.health = 100.0
	w.player.dmgticks = -20
	w.player.spd = 0.05

	w.level1.load1(gl)
	w.level2.load2(gl)

	w.currentLevel = 1

	w.switchLevel(w.currentLevel)
}

func (w *World) getThisLevel() *Level {
	return w.getLevel(w.currentLevel)
}

func (w *World) getLevel(l int) *Level {
	if l == 1 {
		return &w.level1
	}
	if l == 2 {
		return &w.level2
	}
	return nil
}

func (w *World) switchLevel(l int) {
	w.physSys.clearPhysObjs()

	w.player.health = 100.0
	w.player.coins = 0
	w.player.gems = 0
	w.player.beetles = 0
	w.player.obj.phys.pos = mgl32.Vec3{0.0, 0.0, 0.0}
	w.player.obj.phys.rot = mgl32.Vec3{0.0, 0.0, 0.0}

	w.level1.stop()
	w.level2.stop()

	w.seconds = 0
	w.ragelimit = 100.0
	w.rage = w.ragelimit

	w.currentLevel = l

	if l == 1 {
		w.level1.start(w)
		w.level1.ticks = 29 // tick +1 stuff at start
		w.level1.ragelimit = w.ragelimit
		w.timelimit = 120
		w.getThisLevel().exit = mgl32.Vec3{0.0, -1.0, 25.0}
	}
	if l == 2 {
		w.level2.start(w)
		w.level2.ticks = 29
		w.level2.ragelimit = w.ragelimit
		w.timelimit = 180
		w.getThisLevel().exit = mgl32.Vec3{25.8, -1.3, -14.2}
	}

	w.tick()
}
