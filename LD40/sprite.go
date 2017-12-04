package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/webgl"
)

type Sprite struct {
	mesh Mesh

	anim int
	anim_time float64
	time_old float64
	anim_cycle []int
	anim_sprites []mgl32.Vec4
}

func (s *Sprite) animTick() {
	s.anim += 1

	if s.anim >= len(s.anim_cycle) {
		s.anim = 0
	}
}

func (s *Sprite) animUpdate() {
	if timeNow() - s.time_old >= s.anim_time {
		s.animTick()
		s.time_old = timeNow()
	}

	x := s.anim_sprites[s.anim_cycle[s.anim]][0] / (float32)(256.0)
	y := -s.anim_sprites[s.anim_cycle[s.anim]][1] / (float32)(256.0)
	w := s.anim_sprites[s.anim_cycle[s.anim]][2] / (float32)(256.0)
	h := s.anim_sprites[s.anim_cycle[s.anim]][3] / (float32)(256.0)

	t0 := mgl32.Vec4{x/w, y/h, w, h}
	t1 := mgl32.Vec4{x/w + 1.0, y/h, w, h}
	t2 := mgl32.Vec4{x/w + 1.0, y/h - 1.0, w, h}
	t3 := mgl32.Vec4{x/w, y/h - 1.0, w, h}

	s.mesh.texData = []float32{
		t2[0], t2[1], t2[2], t2[3],
		t0[0], t0[1], t0[2], t0[3],
		t1[0], t1[1], t1[2], t1[3],
		t3[0], t3[1], t3[2], t3[3]}
}

func (s *Sprite) animLoad(ac []int, at float64, as []mgl32.Vec4) {
	s.anim_cycle = ac

	s.anim_time = at

	s.anim_sprites = as
}

func (s *Sprite) animDraw(r *Renderer) {
	s.mesh.draw(r)
}

func (s *Sprite) loadSprite(gl *webgl.Context, t string) {
	s.mesh.loadMesh(gl, "gfx/models/quad.obj", t)
}
