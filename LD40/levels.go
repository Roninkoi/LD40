package main

import (
	"github.com/gopherjs/webgl"
	"strconv"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

/*################################
			LEVEL 1
 ################################*/

func addEnemy(t int, x float32, y float32, z float32) Entity {
	newEnemy := Entity{}
	if t == 0 {
		newEnemy.loadEnemy(nil, "nil")
		newEnemy.sprite.animLoad([]int{0, 1, 0, 2}, 250.0, []mgl32.Vec4{
			mgl32.Vec4{1.0, 133.0 + 40.0*0.0, 21.0, 38.0},
			mgl32.Vec4{24.0, 133.0 + 40.0*0.0, 21.0, 38.0},
			mgl32.Vec4{47.0, 133.0 + 40.0*0.0, 21.0, 38.0},
			mgl32.Vec4{70.0, 133.0 + 40.0*0.0, 21.0, 38.0}})
		newEnemy.setEnemy(x, y, z)
		newEnemy.damage = 20.0
		newEnemy.spd = 0.02
	}
	if t == 1 {
		newEnemy.loadEnemy(nil, "nil")
		newEnemy.sprite.animLoad([]int{0, 1, 0, 2}, 250.0, []mgl32.Vec4{
			mgl32.Vec4{1.0, 133.0 + 40.0*1.0, 21.0, 38.0},
			mgl32.Vec4{24.0, 133.0 + 40.0*1.0, 21.0, 38.0},
			mgl32.Vec4{47.0, 133.0 + 40.0*1.0, 21.0, 38.0},
			mgl32.Vec4{70.0, 133.0 + 40.0*1.0, 21.0, 38.0}})
		newEnemy.setEnemy(x, y, z)
		newEnemy.damage = 10.0
		newEnemy.spd = 0.04
	}
	if t == 2 {
		newEnemy.loadEnemy(nil, "nil")
		newEnemy.sprite.animLoad([]int{0, 1, 0, 2}, 250.0, []mgl32.Vec4{
			mgl32.Vec4{1.0, 133.0 + 40.0*2.0, 21.0, 38.0},
			mgl32.Vec4{24.0, 133.0 + 40.0*2.0, 21.0, 38.0},
			mgl32.Vec4{47.0, 133.0 + 40.0*2.0, 21.0, 38.0},
			mgl32.Vec4{70.0, 133.0 + 40.0*2.0, 21.0, 38.0}})
		newEnemy.setEnemy(x, y, z)
		newEnemy.damage = 40.0
		newEnemy.spd = 0.01
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
	return newObj
}

func (o *Obj) tickLoot(playerPos mgl32.Vec3) {
	o.phys.rot[1] += 0.04
	o.phys.pos[1] += (float32)(math.Sin((float64)(o.phys.rot[1]))) * 0.005

	o.update()
}

type Level1 struct {
	level []Obj
	env   []Obj

	loot []Obj

	enemies []Entity

	running bool

	player *Entity

	ticks int
}

func (l *Level1) start(w *World) {
	l.running = true
	l.player = &w.player

	w.physSys.addPhysObj(&l.player.obj)

	for i := 0; i < len(l.level); i++ {
		w.physSys.addPhysObj(&l.level[i])
	}
}

func (l *Level1) stop() {
	l.running = false
}

func (l *Level1) addEnemy(t int, x float32, y float32, z float32) {
	l.enemies = append(l.enemies, addEnemy(t, x, y, z))
}

func (l *Level1) addLoot(t int, x float32, y float32, z float32) {
	l.loot = append(l.loot, addLoot(t, x, y-0.3, z))
}

func (l *Level1) load(gl *webgl.Context) {
	l.level = nil

	for i := 0; i < 14; i++ {
		l.level = append(l.level, Obj{})
		p := "gfx/models/levels/level1_" + strconv.Itoa(i) + ".obj"
		l.level[i].loadObjH(gl, p, "0", false, true, "gfx/textures.png")
	}

	for i := 0; i < 2; i++ {
		l.env = append(l.env, Obj{})
		p := "gfx/models/levels/level1_env_" + strconv.Itoa(i) + ".obj"
		l.env[i].loadObjH(gl, p, "0", false, true, "gfx/sprites.png")
	}

	l.addEnemy(1, 0.0, 0.0, 0.0)
	l.addEnemy(0, 0.1, 0.0, 0.0)
	l.addEnemy(0, 0.2, 0.0, 0.0)
	l.addEnemy(0, 0.3, 0.0, 0.0)
	l.addEnemy(0, 0.4, 0.0, 0.0)
	l.addEnemy(0, 0.5, 0.0, 0.0)
	l.addEnemy(0, 0.6, 0.0, 0.0)
	l.addEnemy(0, 0.7, 0.0, 0.0)
	l.addEnemy(2, 2.0, 0.0, 0.0)
	l.addLoot(0, 2.0, -0.3, 0.0)
	l.addLoot(1, 2.0, -0.3, 2.0)
	l.addLoot(2, 2.0, -0.3, 4.0)
}

func (l *Level1) draw(r *Renderer) {
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
			if l.player.obj.mesh.bsc.Sub(l.enemies[i].obj.phys.pos).Len() <= 20.0 {
				l.enemies[i].sprite.animDraw(r)
			}
		}
		for i := 0; i < len(l.loot); i++ {
			if l.player.obj.mesh.bsc.Sub(l.loot[i].phys.pos).Len() <= 20.0 {
				l.loot[i].draw(r)
			}
		}
	}
}

func (l *Level1) tick(t float64) {
	l.ticks += 1
	if l.running {
		if l.ticks%30 == 0 {
			for i := 0; i < len(l.enemies); i++ {
				l.enemies[i].obj.isects = false
				for j := i + 1; j < len(l.enemies); j++ {
					c := l.enemies[i].obj.phys.pos.Sub(l.enemies[j].obj.phys.pos)
					if c.Len() <= 0.5 {
						l.enemies[i].obj.phys.pos = l.enemies[i].obj.phys.pos.Add(c.Normalize().Mul(0.04))
						l.enemies[j].obj.phys.pos = l.enemies[j].obj.phys.pos.Sub(c.Normalize().Mul(0.04))
					}
				}
			}
		}
		if (l.ticks%2 == 0) {
			for i := 0; i < len(l.enemies); i++ {
				if l.player.obj.mesh.bsc.Sub(l.enemies[i].obj.phys.pos).Len() <= 7.0 {
					l.enemies[i].tickEnemy(l.player.obj.phys.pos, l.player.obj.phys.rot[1])
				}
			}
			for i := 0; i < len(l.loot); i++ {
				if l.player.obj.mesh.bsc.Sub(l.loot[i].phys.pos).Len() <= 20.0 {
					l.loot[i].tickLoot(l.player.obj.phys.pos)
				}
			}
		}
	}
}

/*################################
			LEVEL 2
 ################################*/

type Level2 struct {
	level   []Obj
	running bool
}

func (l *Level2) start(w *World) {
	l.running = true
	for i := 0; i < len(l.level); i++ {
		w.physSys.addPhysObj(&l.level[i])
	}
}

func (l *Level2) stop() {
	l.running = false
}

func (l *Level2) load(gl *webgl.Context) {

}

func (l *Level2) draw(r *Renderer) {
	if l.running {

	}
}

func (l *Level2) tick(t float64) {
	if l.running {

	}
}

/*################################
			LEVEL 3
 ################################*/

type Level3 struct {
	level   []Obj
	running bool
}

func (l *Level3) start(w *World) {
	l.running = true
	for i := 0; i < len(l.level); i++ {
		w.physSys.addPhysObj(&l.level[i])
	}
}

func (l *Level3) stop() {
	l.running = false
}

func (l *Level3) load(gl *webgl.Context) {

}

func (l *Level3) draw(r *Renderer) {
	if l.running {

	}
}

func (l *Level3) tick(t float64) {
	if l.running {

	}
}
