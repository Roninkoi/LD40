package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"strconv"
)

type GUI struct {
	guiTex Texture

	textSprite Sprite

	healthbar Sprite
	ragebar   Sprite

	healthmeter Sprite
	ragemeter   Sprite

	coinicon   Sprite
	gemicon    Sprite
	beetleicon Sprite
	scimitar   Sprite
	ankh       Sprite

	attacking     bool
	scimitarswing float64

	seconds int

	ticks float64

	coinnum   int
	gemnum    int
	beetlenum int

	score int

	health float32
	rage   float32

	hurting bool

	lootPerc int
}

func (g *GUI) renderText(r *Renderer, s string, x float32, y float32, scale float32) { // sprite 158, 0
	xOffs := (float32)(0.0)
	for i := 0; i < len(s); i++ {
		xOffs += 1.0 * (((6.0 / 11.0) * scale) / 0.55)
		ai := []rune(string(s[i]))[0]
		x0 := (ai - 32) % 16
		y0 := (int)(math.Floor((float64)(ai-32) / 16.0))

		g.textSprite.animLoad([]int{0}, 1.0, []mgl32.Vec4{
			{(float32)(x0*6 + 158.0), (float32)(y0 * 11), 6.0, 11.0}})

		g.textSprite.mesh.um = mgl32.Translate3D(xOffs+x, y, -0.12)
		g.textSprite.mesh.um = g.textSprite.mesh.um.Mul4(mgl32.Scale3D(((6.0/11.0)*scale)/0.55, scale, 1.0))
		g.textSprite.mesh.um = g.textSprite.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
		g.textSprite.mesh.update()

		g.textSprite.animDraw(r)
	}
}

func (g *GUI) renderFancyText(r *Renderer, s string, x float32, y float32, scale float32) { // sprite 158, 0
	xOffs := (float32)(0.0)
	for i := 0; i < len(s); i++ {
		xOffs += 1.0 * (((6.0 / 11.0) * scale) / 0.55)
		ai := []rune(string(s[i]))[0]
		x0 := (ai - 32) % 16
		y0 := (int)(math.Floor((float64)(ai-32) / 16.0))

		g.textSprite.animLoad([]int{0}, 1.0, []mgl32.Vec4{
			{(float32)(x0*6 + 158.0), (float32)(y0 * 11), 6.0, 11.0}})

		g.textSprite.mesh.um = mgl32.Translate3D(xOffs+x, y, -0.12)
		g.textSprite.mesh.um = g.textSprite.mesh.um.Mul4(mgl32.Scale3D(((6.0/11.0)*scale)/0.55, scale, 1.0))
		g.textSprite.mesh.um = g.textSprite.mesh.um.Mul4(mgl32.HomogRotate3DY(
			(float32)(math.Pi)*0.5 + (float32)(math.Sin(g.ticks/10.0+(float64)(xOffs*20)))*0.005))
		g.textSprite.mesh.update()

		g.textSprite.animDraw(r)
	}
}

func (g *GUI) tickGUI(ticks float32) {
	if g.attacking {
		g.scimitarswing += 0.2

		if g.scimitarswing >= math.Pi*2.0 {
			g.attacking = false
			g.scimitarswing = 0.0
		}
	}
	g.scimitar.mesh.um = mgl32.Ident4()
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.Translate3D(0.2+0.24, -0.27, -0.18))
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.Scale3D(((36.0/32.0)*0.15)/0.55, 0.15, 1.0))
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.HomogRotate3DZ((float32)(0.5 * math.Sin(g.scimitarswing+math.Pi))))
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.Translate3D(-0.5, 1.0, 0.0))
	//g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.Translate3D(0.2 + 0.09, -0.12, -0.18))
	//g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.Scale3D(((36.0/32.0)*0.15)/0.55, 0.15, 1.0))
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.scimitar.mesh.update()

	hsx := ((122.0 / 11.0) * 0.020 * (-g.health)) / 0.55
	g.healthbar.mesh.um = mgl32.Translate3D(-0.025-(hsx*0.5)-0.2, -0.17, -0.15)
	g.healthbar.mesh.um = g.healthbar.mesh.um.Mul4(mgl32.Scale3D(hsx, 0.020, 1.0))
	g.healthbar.mesh.um = g.healthbar.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.healthbar.mesh.update()

	if g.hurting {
		g.healthmeter.animLoad([]int{0}, 1.0, []mgl32.Vec4{
			{132.0, 113.0, 122.0, 11.0}})
	} else {
		g.healthmeter.animLoad([]int{0}, 1.0, []mgl32.Vec4{
			{102.0, 76.0, 122.0, 11.0}})
	}

	rsy := 0.1 * (-g.rage)
	g.ragebar.mesh.um = mgl32.Translate3D(-0.37, 0.05-(rsy*0.9)-0.09, -0.18)
	g.ragebar.mesh.um = g.ragebar.mesh.um.Mul4(mgl32.Scale3D(((6.0/35.0)*0.1)/0.55, rsy, 1.0))
	g.ragebar.mesh.um = g.ragebar.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.ragebar.mesh.update()
}

func (g *GUI) drawStartScreen(r *Renderer) {
	g.renderFancyText(r, "- Press SPACE to start -", -0.25, -0.1, 0.02)
}

func (g *GUI) drawWinScreen(r *Renderer) {
	g.renderText(r, "- You win! -", -0.13, 0.12, 0.02)
	g.renderText(r, "Score: " + strconv.Itoa(g.score), -0.15, 0.06, 0.02)

	g.renderText(r, "Loot: (" + strconv.Itoa(g.coinnum) + ", " + strconv.Itoa(g.gemnum) + ", " + strconv.Itoa(g.beetlenum) + ") " + strconv.Itoa(g.lootPerc)+"%", -0.22, -0.0, 0.02)

	//g.renderText(r, "100%", -0.05, -0.065, 0.02)

	g.renderText(r, "Time left: " + strconv.Itoa(g.seconds), -0.15, -0.055, 0.02)

	g.renderText(r, "LEFT to previous level, RIGHT to next level", -0.22, -0.11, 0.01)
	g.renderText(r, "SPACE to restart", -0.09, -0.13, 0.01)
}

func (g *GUI) drawLoseScreen(r *Renderer) {
	g.renderText(r, "You're dead!", -0.12, 0.11, 0.02)

	g.renderText(r, "LEFT to previous level", -0.16, -0.06, 0.015)
	g.renderText(r, "SPACE to restart", -0.12, -0.10, 0.015)
}

func (g *GUI) drawGUI(r *Renderer) {
	r.tex = g.guiTex

	g.healthbar.animDraw(r)
	g.healthmeter.animDraw(r)

	g.ragebar.animDraw(r)
	g.ragemeter.animDraw(r)

	g.coinicon.animDraw(r)
	g.gemicon.animDraw(r)
	g.beetleicon.animDraw(r)
	g.ankh.animDraw(r)

	g.scimitar.animDraw(r)

	secondstext := strconv.Itoa(g.seconds)
	g.renderText(r, secondstext, -0.28, -0.08, 0.02)

	cointext := strconv.Itoa(g.coinnum)
	g.renderText(r, cointext, -0.2-0.03+0.035*2.0, 0.135, 0.02)

	gemtext := strconv.Itoa(g.gemnum)
	g.renderText(r, gemtext, 0.0-0.03+0.035, 0.135, 0.02)

	beetletext := strconv.Itoa(g.beetlenum)
	g.renderText(r, beetletext, 0.2-0.03, 0.135, 0.02)
}

func (g *GUI) loadGUI(gl *webgl.Context) {
	g.hurting = false

	g.guiTex.loadTexture(gl, "gfx/sprites.png")

	g.textSprite.loadSprite(nil, "nil")

	g.healthbar.loadSprite(nil, "nil")
	g.healthbar.mesh.um = mgl32.Translate3D(-0.025, -0.17, -0.15)
	g.healthbar.mesh.um = g.healthbar.mesh.um.Mul4(mgl32.Scale3D(((122.0/11.0)*0.020)/0.55, 0.020, 1.0))
	g.healthbar.mesh.um = g.healthbar.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.healthbar.mesh.update()
	g.healthbar.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{102.0, 91.0, 122.0, 11.0}})

	g.healthmeter.loadSprite(nil, "nil")
	g.healthmeter.mesh.um = mgl32.Translate3D(-0.025, -0.17, -0.15)
	g.healthmeter.mesh.um = g.healthmeter.mesh.um.Mul4(mgl32.Scale3D(((122.0/11.0)*0.020)/0.55, 0.020, 1.0))
	g.healthmeter.mesh.um = g.healthmeter.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.healthmeter.mesh.update()
	g.healthmeter.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{102.0, 76.0, 122.0, 11.0}})

	g.ragebar.loadSprite(nil, "nil")
	g.ragebar.mesh.um = mgl32.Translate3D(-0.37, 0.05, -0.18)
	g.ragebar.mesh.um = g.ragebar.mesh.um.Mul4(mgl32.Scale3D(((6.0/35.0)*0.1)/0.55, 0.1, 1.0))
	g.ragebar.mesh.um = g.ragebar.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.ragebar.mesh.update()
	g.ragebar.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{141.0 - 35.0, 2.0, 6.0, 35.0}})

	g.ragemeter.loadSprite(nil, "nil")
	g.ragemeter.mesh.um = mgl32.Translate3D(-0.37, 0.05, -0.18)
	g.ragemeter.mesh.um = g.ragemeter.mesh.um.Mul4(mgl32.Scale3D(((6.0/35.0)*0.1)/0.55, 0.1, 1.0))
	g.ragemeter.mesh.um = g.ragemeter.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.ragemeter.mesh.update()
	g.ragemeter.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{141.0, 2.0, 6.0, 35.0}})

	g.coinicon.loadSprite(nil, "nil")
	g.coinicon.mesh.um = mgl32.Translate3D(-0.2-0.03, 0.17, -0.15)
	g.coinicon.mesh.um = g.coinicon.mesh.um.Mul4(mgl32.Scale3D(((16.0/16.0)*0.025)/0.55, 0.025, 1.0))
	g.coinicon.mesh.um = g.coinicon.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.coinicon.mesh.update()
	g.coinicon.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{84.0, 112.0, 16.0, 16.0}})

	g.gemicon.loadSprite(nil, "nil")
	g.gemicon.mesh.um = mgl32.Translate3D(0.0-0.03, 0.17, -0.15)
	g.gemicon.mesh.um = g.gemicon.mesh.um.Mul4(mgl32.Scale3D(((16.0/16.0)*0.025)/0.55, 0.025, 1.0))
	g.gemicon.mesh.um = g.gemicon.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.gemicon.mesh.update()
	g.gemicon.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{104.0, 112.0, 16.0, 16.0}})

	g.beetleicon.loadSprite(nil, "nil")
	g.beetleicon.mesh.um = mgl32.Translate3D(0.2-0.03, 0.17, -0.15)
	g.beetleicon.mesh.um = g.beetleicon.mesh.um.Mul4(mgl32.Scale3D(((16.0/16.0)*0.025)/0.55, 0.025, 1.0))
	g.beetleicon.mesh.um = g.beetleicon.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.beetleicon.mesh.update()
	g.beetleicon.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{64.0, 112.0, 16.0, 16.0}})

	g.scimitar.loadSprite(nil, "nil")
	g.scimitar.mesh.um = mgl32.Translate3D(0.2+0.09, -0.12, -0.18)
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.Scale3D(((36.0/32.0)*0.15)/0.55, 0.15, 1.0))
	g.scimitar.mesh.um = g.scimitar.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.scimitar.mesh.update()
	g.scimitar.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{108.0, 40.0, 36.0, 32.0}})

	g.ankh.loadSprite(nil, "nil")
	g.ankh.mesh.um = mgl32.Translate3D(-0.33, 0.05, -0.18)
	g.ankh.mesh.um = g.ankh.mesh.um.Mul4(mgl32.Scale3D(((20.0/20.0)*0.04)/0.55, 0.04, 1.0))
	g.ankh.mesh.um = g.ankh.mesh.um.Mul4(mgl32.HomogRotate3DY((float32)(math.Pi) * 0.5))
	g.ankh.mesh.update()
	g.ankh.animLoad([]int{0}, 1.0, []mgl32.Vec4{
		{116.0, 4.0, 20.0, 20.0}})
}
