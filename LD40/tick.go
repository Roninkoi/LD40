package main

import (
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

func (g *Game) gameInput() {
	g.input.getKeys()

	playerMov := mgl32.Vec3{0.0, 0.0, 0.0}

	if g.input.keys[A] {
		playerMov[0] += -0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
		playerMov[2] += -0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
	}
	if g.input.keys[D] {
		playerMov[0] += 0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
		playerMov[2] += 0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
	}
	if g.input.keys[W] {
		playerMov[0] += 0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
		playerMov[2] += -0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
	}
	if g.input.keys[S] {
		playerMov[0] += -0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
		playerMov[2] += 0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
	}

	/*if g.input.keys[R] {
		playerMov[1] = 0.05
	}
	if g.input.keys[F] {
		playerMov[1] = -0.05
	}*/

	if playerMov.Len() > 0.0 {
		g.world.player.obj.phys.v[0] = playerMov[0]
		g.world.player.obj.phys.v[2] = playerMov[2]
	} else if g.world.player.obj.isects {
		g.world.player.obj.phys.v = g.world.player.obj.phys.v.Mul(0.9)
	}

	if g.input.keys[LEFT] {
		g.renderer.camRot[1] -= 0.05
	}
	if g.input.keys[RIGHT] {
		g.renderer.camRot[1] += 0.05
	}
}

func (g *Game) tick() {
	tickdelta := timeNow()

	g.ticks += 1.0

	pcam := g.world.player.obj.phys.pos.Mul(-0.12)
	g.renderer.camPos = g.renderer.camPos.Mul(0.88)
	g.renderer.camPos = g.renderer.camPos.Add(pcam)

	g.world.ticks = g.ticks
	g.world.tick()

	g.gameInput()

	g.tick_time += timeNow() - tickdelta
}
