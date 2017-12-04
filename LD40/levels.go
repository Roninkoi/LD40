package main

import (
	"github.com/gopherjs/webgl"
	"strconv"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Level struct {
	level []Obj
	env   []Obj

	loot []Obj

	enemies []Entity

	loot0 []Obj

	enemies0 []Entity

	running bool

	player *Entity

	ticks int

	exit mgl32.Vec3

	rage float64
	ragelimit float64
}

func (l *Level) start(w *World) {
	l.running = true
	l.player = &w.player

	w.physSys.addPhysObj(&l.player.obj)

	for i := 0; i < len(l.level); i++ {
		w.physSys.addPhysObj(&l.level[i])
	}

	l.resetEnemies()
	l.resetLoot()
}

func (l *Level) stop() {
	l.running = false
}

func (l *Level) addEnemy(t int, x float32, y float32, z float32) {
	l.enemies = append(l.enemies, addEnemy(t, x, y, z))
	l.enemies0 = append(l.enemies0, addEnemy(t, x, y, z))
}

func (l *Level) addLoot(t int, x float32, y float32, z float32) {
	l.loot = append(l.loot, addLoot(t, x, y-0.3, z))
	l.loot0 = append(l.loot0, addLoot(t, x, y-0.3, z))
}

func (l *Level) removeEnemy(i int) {
	l.enemies[i].removed = true
	//l.enemies = append(l.enemies[:i], l.enemies[i+1:]...)
}

func (l *Level) removeLoot(i int) {
	l.loot[i].removed = true
	//l.loot = append(l.loot[:i], l.loot[i+1:]...)
}

func (l *Level) resetEnemies() {
	for i := 0; i < len(l.enemies); i++ {
		l.enemies[i] = l.enemies0[i]
		l.enemies[i].removed = false
		l.enemies[i].tickEnemy(l.player.obj.phys.pos, l.player.obj.phys.rot[1], false)
	}
}

func (l *Level) resetLoot() {
	for i := 0; i < len(l.loot); i++ {
		l.loot[i] = l.loot0[i]
		l.loot[i].removed = false
		l.loot[i].tickLoot(l.player.obj.phys.pos)
	}
}

func (l *Level) draw(r *Renderer) {
	if l.running {
		for i := 0; i < len(l.level); i++ {
			if l.player.obj.mesh.bsc.Sub(l.level[i].mesh.bsc).Len() <= l.player.obj.mesh.bsr+l.level[i].mesh.bsr+10.0 {
				l.level[i].draw(r)
			}
		}
		r.cflush()
		for i := 0; i < len(l.env); i++ {
			r.tex = l.env[i].mesh.tex // make sure tex found
			if l.player.obj.mesh.bsc.Sub(l.env[i].mesh.bsc).Len() <= l.player.obj.mesh.bsr+l.env[i].mesh.bsr+10.0 {
				l.env[i].draw(r)
			}
		}
		for i := 0; i < len(l.enemies); i++ {
			if !l.enemies[i].removed {
				if l.player.obj.mesh.bsc.Sub(l.enemies[i].obj.phys.pos).Len() <= 20.0 {
					l.enemies[i].sprite.animDraw(r)
				}
			}
		}
		for i := 0; i < len(l.loot); i++ {
			if !l.loot[i].removed {
				if l.player.obj.mesh.bsc.Sub(l.loot[i].phys.pos).Len() <= 20.0 {
					l.loot[i].draw(r)
				}
			}
		}
	}
}

func (l *Level) tick(t float64) {
	l.ticks += 1
	if l.running {
		if l.ticks%30 == 0 {
			for i := 0; i < len(l.enemies); i++ {
				l.enemies[i].obj.isects = false
				for j := i + 1; j < len(l.enemies); j++ {
					if !l.enemies[i].removed && !l.enemies[j].removed {
						c := l.enemies[i].obj.phys.pos.Sub(l.enemies[j].obj.phys.pos)
						if c.Len() <= 0.5 {
							l.enemies[i].obj.phys.pos = l.enemies[i].obj.phys.pos.Add(c.Normalize().Mul(0.04))
							l.enemies[j].obj.phys.pos = l.enemies[j].obj.phys.pos.Sub(c.Normalize().Mul(0.04))
						}
					}
				}
			}
		}
		ragedamage := 1.0/(l.rage/l.ragelimit+0.2)
		if l.ticks%2 == 0 {
			for i := 0; i < len(l.enemies); i++ {
				if !l.enemies[i].removed {
					if l.player.obj.mesh.bsc.Sub(l.enemies[i].obj.phys.pos).Len() <= 7.0 {
						l.enemies[i].tickEnemy(l.player.obj.phys.pos, l.player.obj.phys.rot[1], l.player.attacking)
					}
					if l.enemies[i].attacking && l.ticks%30 == 0 {
						l.player.attack(l.enemies[i].damage*ragedamage)
					}
				}
			}
			for i := 0; i < len(l.enemies); i++ {
				if l.enemies[i].health == 0.0 && !l.enemies[i].removed {
					l.removeEnemy(i)
				}
			}
			l.player.attacking = false
			for i := 0; i < len(l.loot); i++ {
				if !l.loot[i].removed {
					if l.player.obj.mesh.bsc.Sub(l.loot[i].phys.pos).Len() <= 20.0 {
						l.loot[i].tickLoot(l.player.obj.phys.pos)
					}
					if l.loot[i].collides {
						l.loot[i].lootPickup(l.player)

						l.removeLoot(i)
					}
				}
			}
		}
	}
}

func addEnemy(t int, x float32, y float32, z float32) Entity {
	newEnemy := Entity{}
	newEnemy.removed = false
	if t == 1 {
		newEnemy.loadEnemy(nil, "nil")
		newEnemy.sprite.animLoad([]int{0, 1, 0, 2}, 250.0, []mgl32.Vec4{
			{1.0, 133.0 + 40.0*0.0, 21.0, 38.0},
			{24.0, 133.0 + 40.0*0.0, 21.0, 38.0},
			{47.0, 133.0 + 40.0*0.0, 21.0, 38.0},
			{70.0, 133.0 + 40.0*0.0, 21.0, 38.0}})
		newEnemy.setEnemy(x, y, z)
		newEnemy.damage = 20.0
		newEnemy.spd = 0.025
		newEnemy.health = 120.0
	}
	if t == 0 {
		newEnemy.loadEnemy(nil, "nil")
		newEnemy.sprite.animLoad([]int{0, 1, 0, 2}, 250.0, []mgl32.Vec4{
			{1.0, 133.0 + 40.0*1.0, 21.0, 38.0},
			{24.0, 133.0 + 40.0*1.0, 21.0, 38.0},
			{47.0, 133.0 + 40.0*1.0, 21.0, 38.0},
			{70.0, 133.0 + 40.0*1.0, 21.0, 38.0}})
		newEnemy.setEnemy(x, y, z)
		newEnemy.damage = 10.0
		newEnemy.spd = 0.04
		newEnemy.health = 60.0
	}
	if t == 2 {
		newEnemy.loadEnemy(nil, "nil")
		newEnemy.sprite.animLoad([]int{0, 1, 0, 2}, 250.0, []mgl32.Vec4{
			{1.0, 133.0 + 40.0*2.0, 21.0, 38.0},
			{24.0, 133.0 + 40.0*2.0, 21.0, 38.0},
			{47.0, 133.0 + 40.0*2.0, 21.0, 38.0},
			{70.0, 133.0 + 40.0*2.0, 21.0, 38.0}})
		newEnemy.setEnemy(x, y, z)
		newEnemy.damage = 40.0
		newEnemy.spd = 0.015
		newEnemy.health = 300.0
	}
	return newEnemy
}

func addLoot(t int, x float32, y float32, z float32) Obj {
	newObj := Obj{}
	if t == 0 {
		newObj.loadObj(nil, "gfx/models/coin.obj", "nil")
	}
	if t == 1 {
		newObj.loadObj(nil, "gfx/models/gem.obj", "nil")
	}
	if t == 2 {
		newObj.loadObj(nil, "gfx/models/beetle.obj", "nil")
	}
	newObj.phys.pos = mgl32.Vec3{x, y, z}
	newObj.obj_type = t
	return newObj
}

func (o *Obj) tickLoot(playerPos mgl32.Vec3) {
	o.phys.rot[1] += 0.04
	o.phys.pos[1] += (float32)(math.Sin((float64)(o.phys.rot[1]))) * 0.005

	o.update()

	c := o.phys.pos.Sub(playerPos).Len()

	if c < 0.8 {
		o.collides = true
	}
}

func (o *Obj) lootPickup(p *Entity) {
	if o.obj_type == 0 {
		p.coins += 1
	}
	if o.obj_type == 1 {
		p.gems += 1
	}
	if o.obj_type == 2 {
		p.beetles += 1
	}
}

func (l *Level) load1(gl *webgl.Context) {
	l.level = nil

	l.exit = mgl32.Vec3{0.0, -1.0, 25.0}

	for i := 0; i < 14; i++ {
		l.level = append(l.level, Obj{})
		p := "gfx/models/level1/level_" + strconv.Itoa(i) + ".obj"
		l.level[i].loadObjH(gl, p, "0", false, true, "gfx/textures.png")
	}

	for i := 0; i < 5; i++ {
		l.env = append(l.env, Obj{})
		p := "gfx/models/level1/level_env_" + strconv.Itoa(i) + ".obj"
		l.env[i].loadObjH(gl, p, "0", false, true, "gfx/sprites.png")
	}

	l.addEnemy(0, 0.0, 1.0, 13.8)
	l.addEnemy(0, 15.7, 1.0, -0.8)
	l.addEnemy(0, 16.9, 1.0, 0.8)
	l.addEnemy(0, 20.7, 1.0, -0.8)

	l.addEnemy(1, 20.7, 1.0, 13.8)
	l.addEnemy(0, 22.7, 1.0, 12.8)

	l.addEnemy(1, 19.7, 1.0, -12.8)
	l.addEnemy(1, 20.7, 1.0, -15.8)
	l.addEnemy(1, 22.7, 1.0, -14.8)

	l.addLoot(0, 2.3, -0.3, 0.0)
	l.addLoot(1, 1.2, -0.3, 2.0)
	l.addLoot(2, 2.0, -0.3, 4.0)
	l.addLoot(0, 15.0, 1.1, -0.2)
	l.addLoot(1, 17.7, 1.1, 0.7)
	l.addLoot(2, 20.2,1.1, -15.0)

	l.addLoot(1, 18.2,1.1, -13.0)

	l.addLoot(0, 22.2,1.1, -12.0)
	l.addLoot(2, 20.2,1.1, 15.0)

	l.addLoot(0, 2.2,-0.8, 20.0)
	l.addLoot(0, -0.5,-0.8, 18.5)
	l.addLoot(0, -2.0,-0.8, 20.7)
}

func (l *Level) load2(gl *webgl.Context) {
	l.level = nil

	for i := 0; i < 0; i++ {
		l.level = append(l.level, Obj{})
		p := "gfx/models/levels/level2_" + strconv.Itoa(i) + ".obj"
		l.level[i].loadObjH(gl, p, "0", false, true, "gfx/textures.png")
	}

	for i := 0; i < 0; i++ {
		l.env = append(l.env, Obj{})
		p := "gfx/models/levels/level2_env_" + strconv.Itoa(i) + ".obj"
		l.env[i].loadObjH(gl, p, "0", false, true, "gfx/sprites.png")
	}

	l.addEnemy(1, 0.0, 0.0, 0.0)
	l.addEnemy(0, 0.1, 0.0, 0.0)
	l.addEnemy(2, 2.0, 0.0, 0.0)

	l.addLoot(0, 2.0, -0.3, 0.0)
	l.addLoot(1, 2.0, -0.3, 2.0)
	l.addLoot(2, 2.0, -0.3, 4.0)
}

func (l *Level) load3(gl *webgl.Context) {
	l.level = nil

	for i := 0; i < 0; i++ {
		l.level = append(l.level, Obj{})
		p := "gfx/models/levels/level3_" + strconv.Itoa(i) + ".obj"
		l.level[i].loadObjH(gl, p, "0", false, true, "gfx/textures.png")
	}

	for i := 0; i < 0; i++ {
		l.env = append(l.env, Obj{})
		p := "gfx/models/levels/level3_env_" + strconv.Itoa(i) + ".obj"
		l.env[i].loadObjH(gl, p, "0", false, true, "gfx/sprites.png")
	}

	l.addEnemy(1, 0.0, 0.0, 0.0)
	l.addEnemy(0, 0.1, 0.0, 0.0)
	l.addEnemy(2, 2.0, 0.0, 0.0)

	l.addLoot(0, 2.0, -0.3, 0.0)
	l.addLoot(1, 2.0, -0.3, 2.0)
	l.addLoot(2, 2.0, -0.3, 4.0)
}
