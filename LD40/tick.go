package main

import (
	"math"
)

func (g *Game) gameInput() {
	g.input.getKeys()

	if g.input.keys[A] {
		g.renderer.camPos[0] += 0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
		g.renderer.camPos[2] += 0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
	}
	if g.input.keys[D] {
		g.renderer.camPos[0] -= 0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
		g.renderer.camPos[2] -= 0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
	}
	if g.input.keys[W] {
		g.renderer.camPos[0] -= 0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
		g.renderer.camPos[2] += 0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
	}
	if g.input.keys[S] {
		g.renderer.camPos[0] += 0.05*(float32)(math.Sin((float64)(g.renderer.camRot[1])))
		g.renderer.camPos[2] -= 0.05*(float32)(math.Cos((float64)(g.renderer.camRot[1])))
	}

	if g.input.keys[R] {
		g.renderer.camPos[1] -= 0.05
	}
	if g.input.keys[F] {
		g.renderer.camPos[1] += 0.05
	}

	if g.input.keys[LEFT] {
		g.renderer.camRot[1] -= 0.05
	}
	if g.input.keys[RIGHT] {
		g.renderer.camRot[1] += 0.05
	}
}

func (g *Game) tick() {
	g.ticks += 1.0

	g.world.ticks = g.ticks
	g.world.tick()

	g.gameInput()
}
