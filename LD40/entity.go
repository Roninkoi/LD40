package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Entity struct {
	obj    Obj
	sprite Sprite

	health float64
	damage float64
	spd    float32

	ticks int

	dmgticks int

	coins int
	gems int
	beetles int

	attacking bool
}

func (e *Entity) loadEnemy(gl *webgl.Context, p string) {
	e.sprite.loadSprite(gl, p)
}

func (e *Entity) setEnemy(x float32, y float32, z float32) {
	e.health = 100.0
	e.obj.phys.pos = mgl32.Vec3{x, y, z}
}

func (e *Entity) tickEnemy(playerPos mgl32.Vec3, camY float32, attacking bool) {
	e.ticks += 1

	if e.ticks - e.dmgticks >= 15 {
		e.sprite.anim_cycle = []int{0, 1, 0, 2}
	}

	e.sprite.mesh.um = mgl32.Translate3D(e.obj.phys.pos[0], e.obj.phys.pos[1], e.obj.phys.pos[2])
	e.sprite.mesh.um = e.sprite.mesh.um.Mul4(mgl32.HomogRotate3DY(-camY + math.Pi*0.5))
	e.sprite.mesh.update()

	e.sprite.animUpdate()

	c := playerPos.Sub(e.obj.phys.pos.Add(mgl32.Vec3{0.0, -0.2, 0.0}))

	movVec := c.Normalize().Mul(e.spd)
	if c.Len() > 1.0 {
		e.obj.phys.pos = e.obj.phys.pos.Add(movVec)
	}
	e.obj.phys.pos[1] += movVec[1]

	lookVec := mgl32.Vec3{(float32)(math.Sin((float64)(-camY))), 0.0, (float32)(math.Cos((float64)(-camY)))}

	if attacking && lookVec.Dot(c.Normalize()) > 0.7 && c.Len() < 1.5 {
		e.attack(30.0)
	}
}

func (e *Entity) attack(dmg float64) {
	e.health -= dmg

	e.dmgticks = e.ticks

	e.sprite.anim_cycle = []int{3}
	e.sprite.anim = 0

	if e.health < 0.0 {
		e.health = 0.0
	}
}
