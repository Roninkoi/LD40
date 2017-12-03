package main

import "github.com/gopherjs/webgl"

/*################################
			LEVEL 1
 ################################*/
 
type Level1 struct {
	level []Obj
	running bool

	player *Entity

	room Obj
	ramp Obj
	ramp2 Obj
	test Obj
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

	l.level = append(l.level, Obj{})
	l.level[0].loadObjH(gl, "gfx/models/room0.obj", "0", false, true, "gfx/checker.png")

	l.level = append(l.level, Obj{})
	l.level[1].loadObjH(gl, "gfx/models/room1.obj", "0", false, true, "gfx/checker.png")
}

func (l *Level1) draw(r *Renderer) {
	if l.running {
		for i := 0; i < len(l.level); i++ {
			l.level[i].draw(r)
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
	level []Obj
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
	level []Obj
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
	level []Obj
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
 