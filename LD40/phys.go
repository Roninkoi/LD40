package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Phys struct {
	pos mgl32.Vec3
	rot mgl32.Vec3
	s   mgl32.Vec3

	v mgl32.Vec3
	a mgl32.Vec3
}

func (p *Phys) init() {
	p.pos = mgl32.Vec3{0.0, 0.0, 0.0}
	p.rot = mgl32.Vec3{0.0, 0.0, 0.0}
	p.s = mgl32.Vec3{1.0, 1.0, 1.0}

	p.v = mgl32.Vec3{0.0, 0.0, 0.0}
	p.a = mgl32.Vec3{0.0, 0.0, 0.0}
}

func (p *Phys) update() {

}

type PhysSys struct {
	objs []*Obj

	ticks float64
}

func (p *PhysSys) update() {
	for i := 0; i < len(p.objs); i++ {
		for j := i + 1; j < len(p.objs); j++ {
			if p.objs[i].mesh.bsi(&p.objs[j].mesh) {
				/*if (int)(p.ticks)%20 == 0 || true {
					fmt.Println("intersecting")
				}*/
			}
		}
	}
}

func (p *PhysSys) clearPhysObjs() {
	p.objs = nil
}

func (p *PhysSys) addPhysObj(o *Obj) {
	p.objs = append(p.objs, o)
}
