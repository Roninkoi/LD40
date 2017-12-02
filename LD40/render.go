package main

import (
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
	"fmt"
	_ "math"
)

const BATCH_SIZE = 16384

type Renderer struct {
	gl *webgl.Context

	ticks float64

	shader Shader

	u_matrix mgl32.Mat4
	c_matrix mgl32.Mat4
	p_matrix mgl32.Mat4

	ul_um  *js.Object
	ul_cm  *js.Object
	ul_pm  *js.Object
	ul_tex *js.Object

	al_pos int
	al_tex int

	vertexBuffer *js.Object
	texBuffer    *js.Object
	indexBuffer  *js.Object

	vertexData [BATCH_SIZE]float32
	texData    [BATCH_SIZE]float32
	indexData  [BATCH_SIZE]uint16

	tex Texture

	camPos mgl32.Vec3
	camRot mgl32.Vec3

	verts int
	inds  int

	draws int

	vertexNum int
	indexNum int
}

func (r *Renderer) init(gl *webgl.Context) {
	r.gl = gl

	r.camPos = mgl32.Vec3{0.0, 0.0, 0.0}
	r.camRot = mgl32.Vec3{0.0, 0.0, 0.0}

	r.verts = 0
	r.inds = 0

	r.shader.loadShader(gl, "shader/vertexShader.vert", "shader/fragmentShader.frag")

	r.gl.UseProgram(r.shader.program)

	r.al_pos = gl.GetAttribLocation(r.shader.program, "a_pos")
	r.al_tex = gl.GetAttribLocation(r.shader.program, "a_tex")

	r.ul_um = gl.GetUniformLocation(r.shader.program, "u")
	r.ul_cm = gl.GetUniformLocation(r.shader.program, "c")
	r.ul_pm = gl.GetUniformLocation(r.shader.program, "p")
	r.ul_tex = gl.GetUniformLocation(r.shader.program, "texture")

	fmt.Println(gl.GetAttribLocation(r.shader.program, "a_pos"))
	fmt.Println(gl.GetAttribLocation(r.shader.program, "a_tex"))

	gl.EnableVertexAttribArray(r.al_pos)
	gl.EnableVertexAttribArray(r.al_tex)

	r.vertexBuffer = gl.CreateBuffer()
	r.texBuffer = gl.CreateBuffer()
	r.indexBuffer = gl.CreateBuffer()

	r.tex.loadTexture(gl, "gfx/test.png")
}

func (r *Renderer) clear() {
	r.gl.ClearColor(0.5, 0.7, 1.0, 1.0)
	r.gl.Clear(r.gl.COLOR_BUFFER_BIT | r.gl.DEPTH_BUFFER_BIT)
}

func mat4ToFloat32(matrix mgl32.Mat4) []float32 {
	return []float32{
		matrix.At(0, 0), matrix.At(0, 1), matrix.At(0, 2), matrix.At(0, 3),
		matrix.At(1, 0), matrix.At(1, 1), matrix.At(1, 2), matrix.At(1, 3),
		matrix.At(2, 0), matrix.At(2, 1), matrix.At(2, 2), matrix.At(2, 3),
		matrix.At(3, 0), matrix.At(3, 1), matrix.At(3, 2), matrix.At(3, 3)}
}

func (r *Renderer) batchAdd(va *[]float32, ta *[]float32, ia *[]uint16) {
	vlen := len(*va)
	ilen := len(*ia)

	if r.verts+vlen >= BATCH_SIZE || r.inds+ilen >= BATCH_SIZE {
		r.flush()
	}

	vertices := (uint16)((float32)(r.verts) / 4.0)

	for i := 0; i < vlen; i++ {
		r.vertexData[r.verts] = (*va)[i]
		r.texData[r.verts] = (*ta)[i]
		r.verts += 1
	}

	for i := 0; i < ilen; i++ {
		r.indexData[r.inds] = (*ia)[i] + vertices
		r.inds += 1
	}
}

func (r *Renderer) draw(t *Texture, va *[]float32, ta *[]float32, ia *[]uint16) {
	if t.p != r.tex.p {
		r.cflush()
	}

	r.tex = *t
	r.u_matrix = mgl32.Ident4()

	r.batchAdd(va, ta, ia)
}

func (r *Renderer) render(t *Texture, um *mgl32.Mat4, va *[]float32, ta *[]float32, ia *[]uint16) {
	r.cflush()

	r.tex = *t
	r.u_matrix = *um

	r.batchAdd(va, ta, ia)
	r.flush()
}

func (r *Renderer) cflush() {
	if r.inds >= 3 {
		r.flush()
	}
}

func (r *Renderer) flush() {
	r.draws += 1

	r.c_matrix = mgl32.Ident4()
	r.p_matrix = mgl32.Ident4()
	r.p_matrix = mgl32.Perspective(90.0, 1280.0/750.0, 0.1, 100.0)

	r.c_matrix = r.c_matrix.Mul4(mgl32.HomogRotate3DX(r.camRot[0]))
	r.c_matrix = r.c_matrix.Mul4(mgl32.HomogRotate3DY(r.camRot[1]))
	r.c_matrix = r.c_matrix.Mul4(mgl32.HomogRotate3DZ(r.camRot[2]))

	r.c_matrix = r.c_matrix.Mul4(mgl32.Translate3D(r.camPos[0], r.camPos[1], r.camPos[2]))

	r.gl.Viewport(0, 0, 1280, 750)

	r.gl.UseProgram(r.shader.program)

	r.gl.UniformMatrix4fv(r.ul_um, false, mat4ToFloat32(r.u_matrix.Transpose()))

	r.gl.UniformMatrix4fv(r.ul_cm, false, mat4ToFloat32(r.c_matrix.Transpose()))

	r.gl.UniformMatrix4fv(r.ul_pm, false, mat4ToFloat32(r.p_matrix.Transpose()))

	r.gl.Disable(r.gl.CULL_FACE)
	r.gl.Disable(r.gl.BLEND)

	r.gl.Enable(r.gl.DEPTH_TEST)
	r.gl.DepthFunc(r.gl.LESS)

	r.gl.ActiveTexture(r.gl.TEXTURE0)
	r.gl.BindTexture(r.gl.TEXTURE_2D, r.tex.tex)
	r.gl.Uniform1i(r.ul_tex, 0)

	r.gl.BindBuffer(r.gl.ARRAY_BUFFER, r.vertexBuffer)
	r.gl.BufferData(r.gl.ARRAY_BUFFER, r.vertexData, r.gl.DYNAMIC_DRAW)
	r.gl.VertexAttribPointer(r.al_pos, 4, r.gl.FLOAT, false, 0, 0)

	r.gl.BindBuffer(r.gl.ARRAY_BUFFER, r.texBuffer)
	r.gl.BufferData(r.gl.ARRAY_BUFFER, r.texData, r.gl.DYNAMIC_DRAW)
	r.gl.VertexAttribPointer(r.al_tex, 4, r.gl.FLOAT, false, 0, 0)

	r.gl.BindBuffer(r.gl.ELEMENT_ARRAY_BUFFER, r.indexBuffer)
	r.gl.BufferData(r.gl.ELEMENT_ARRAY_BUFFER, r.indexData, r.gl.DYNAMIC_DRAW)

	r.gl.DrawElements(r.gl.TRIANGLES, r.inds, r.gl.UNSIGNED_SHORT, 0)

	r.vertexNum += r.verts
	r.indexNum += r.inds

	r.verts = 0
	r.inds = 0
}

func (g *Game) render() {
	rt_delta := timeNow()

	g.renderer.draws = 0
	g.renderer.vertexNum = 0
	g.renderer.indexNum = 0

	g.renderer.ticks = g.ticks
	g.renderer.clear()

	g.world.draw(&g.renderer)

	g.renderer.cflush()

	if (int)(g.ticks)%60 == 0 {
		fmt.Print("draws ")
		fmt.Print(g.renderer.draws)
		fmt.Print(", vertices ")
		fmt.Print(g.renderer.vertexNum)
		fmt.Print(", indices ")
		fmt.Println(g.renderer.indexNum)
	}

	g.render_time += timeNow() - rt_delta
}
