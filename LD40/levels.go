package main

import (
	"github.com/gopherjs/webgl"
	"strconv"
)

/*################################
			LEVEL 1
 ################################*/

type Level1 struct {
	level []Obj
	env   []Obj

	running bool

	player *Entity
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
}

func (l *Level1) draw(r *Renderer) {
	if l.running {
		for i := 0; i < len(l.level); i++ {
			if l.player.obj.mesh.bsc.Sub(l.level[i].mesh.bsc).Len() <= l.player.obj.mesh.bsr + l.level[i].mesh.bsr + 10.0 {
				l.level[i].draw(r)
			}
		}
		for i := 0; i < len(l.env); i++ {
			if l.player.obj.mesh.bsc.Sub(l.env[i].mesh.bsc).Len() <= l.player.obj.mesh.bsr + l.env[i].mesh.bsr + 10.0 {
				l.env[i].draw(r)
			}
		}
	}
}

func (l *Level1) tick(t float64) {
	if l.running {

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

/*################################
			LEVEL 4
 ################################*/

type Level4 struct {
	level   []Obj
	running bool
}

func (l *Level4) start(w *World) {
	l.running = true
	for i := 0; i < len(l.level); i++ {
		w.physSys.addPhysObj(&l.level[i])
	}
}

func (l *Level4) stop() {
	l.running = false
}

func (l *Level4) load(gl *webgl.Context) {

}

func (l *Level4) draw(r *Renderer) {
	if l.running {

	}
}

func (l *Level4) tick(t float64) {
	if l.running {

	}
}
