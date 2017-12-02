package main

import (
	"fmt"
	"time"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

type Game struct {
	Running bool

	ticks float64
	fps   int

	fps_ticks int
	fps_time  float64

	render_time float64

	tick_time float64

	time float64

	document *js.Object
	canvas   *js.Object
	gl *webgl.Context

	renderer Renderer

	world World

	input Input
}

func (g *Game) Start() {
	g.ticks = 0
	g.fps = 0
	g.fps_ticks = 0

	g.Running = true

	g.document = js.Global.Get("document")
	g.canvas = g.document.Call("createElement", "canvas")
	g.document.Get("body").Call("appendChild", g.canvas)

	g.canvas.Set("width", 1280)
	g.canvas.Set("height", 750)

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false
	attrs.Depth = true
	attrs.Antialias = false

	g.gl, _ = webgl.NewContext(g.canvas, attrs)

	g.renderer.init(g.gl)

	g.world.loadWorld(g.gl)

	g.input.init()
	g.input.keyHandler()

	g.main(nil)
}

func timeNow() float64 {
	return (float64)(time.Now().UnixNano()) / 1000000.0
}

func (g *Game) main(ftime *js.Object) {
	g.fps_ticks++

	g.time = timeNow()

	if g.time - g.fps_time >= 1000.0 {
		g.fps_time = g.time

		g.fps = g.fps_ticks

		fmt.Printf("fps %d%s%.2f%s", g.fps,
			", rt: ", g.render_time/(float64)(g.fps_ticks), " ms\n")

		g.render_time = 0

		g.fps_ticks = 0
	}

	g.tick()

	g.render()

	for timeNow() - g.time < 16.0 {
		// wait around to sync fps on fast screens
	}

	js.Global.Call("requestAnimationFrame", g.main)
}
