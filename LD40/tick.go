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

	if g.input.keys[SPACE] {
		if g.input.space_pressed {
			g.input.space_pressed = false

			g.world.player.attacking = true
			g.gui.attacking = true

			if (g.win || g.lose) {
				g.world.restart()
			}
			if !g.start {
				g.win = false
				g.lose = false
				g.stopRender = false
			}
			g.start = false
		}
	} else {
		g.input.space_pressed = true
	}

	if g.input.keys[LEFT] {
		if g.input.left_pressed {
			g.input.left_pressed = false

			if (g.win || g.lose) {
				prev := g.world.currentLevel - 1

				if prev <= 0 {
					prev = 1
				}
				g.world.switchLevel(prev)
			}
			if !g.start {
				g.win = false
				g.lose = false
				g.stopRender = false
			}
			g.start = false
		}
	} else {
		g.input.left_pressed = true
	}
	if g.input.keys[RIGHT] {
		if g.input.right_pressed {
			g.input.right_pressed = false

			if (g.win || g.lose) {
				next := g.world.currentLevel + 1

				if next >= 3 {
					next = 3
				}
				g.world.switchLevel(next)
			}
			if !g.start {
				g.win = false
				g.lose = false
				g.stopRender = false
			}
			g.start = false
		}
	} else {
		g.input.right_pressed = true
	}

	if playerMov.Len() > 0.0 {
		g.world.player.obj.phys.v[0] = playerMov[0]
		g.world.player.obj.phys.v[2] = playerMov[2]
	} else if g.world.player.obj.isects {
		g.world.player.obj.phys.v = g.world.player.obj.phys.v.Mul(0.9)
	}

	if g.input.keys[LEFT] {
		g.world.player.obj.phys.rot[1] -= 0.05
	}
	if g.input.keys[RIGHT] {
		g.world.player.obj.phys.rot[1] += 0.05
	}
}

func (g *Game) tick() {
	tickdelta := timeNow()

	g.ticks += 1.0

	pcam := g.world.player.obj.phys.pos.Mul(-0.12)
	g.renderer.camPos = g.renderer.camPos.Mul(0.88)
	g.renderer.camPos = g.renderer.camPos.Add(pcam)

	g.gameInput()

	if g.start {
		g.renderer.camRot[1] = (float32)(math.Sin(g.ticks/100.0))
	} else {
		g.renderer.camRot = g.world.player.obj.phys.rot

		g.world.ticks = g.ticks
		g.world.tick()

		if g.world.getThisLevel().exit.Sub(g.world.player.obj.phys.pos).Len() < 3.0 {
			g.win = true
		}
	}

	if g.world.player.health <= 0.0 || g.world.rage <= 0.0 || g.world.timer <= 0 {
		g.lose = true
	}

	g.world.score = (int)((float32)(g.world.player.coins*4 + g.world.player.gems*8 + g.world.player.beetles*16)*((float32)(g.world.seconds)*0.1+1.0))

	g.gui.ticks = g.ticks
	g.gui.health = (float32)(g.world.player.health / 100.0)
	g.gui.rage = (float32)(g.world.rage) / 100.0
	g.gui.score = g.world.score

	g.gui.hurting = g.world.player.ticks - g.world.player.dmgticks < 15

	g.gui.lootPerc = (int)((float32)(g.world.player.coins + g.world.player.gems + g.world.player.beetles)/
		((float32)(len(g.world.getThisLevel().loot0)))*100.0)

	g.gui.coinnum = g.world.player.coins
	g.gui.gemnum = g.world.player.gems
	g.gui.beetlenum = g.world.player.beetles
	g.gui.seconds = g.world.timer

	g.gui.tickGUI((float32)(g.ticks))

	g.tick_time += timeNow() - tickdelta
}
