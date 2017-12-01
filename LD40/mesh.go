package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"strconv"
	"fmt"
)

type Mesh struct {
	rawVertexData []float32
	vertexData []float32
	texData []float32
	indexData []uint16

	tex Texture

	um mgl32.Mat4

	bsc mgl32.Vec3
	bsr float32
}

func parse(s string) []string {
	var returns []string

	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\r' {
			word += string(s[i])
		} else {
			returns = append(returns, word)
			word = ""
		}
	}

	return returns
}

/*
 * fix this thing later (textures missing)
 */
func (m *Mesh) parseFace(s string) (f int, t int) {
	face := ""
	tex := ""

	texstart := false
	for i := 0; i < len(s); i++ {
		if s[i] != '/' && !texstart {
			face += string(s[i])
		} else {
			texstart = true
		}
	}

	facei, _ := strconv.Atoi(face)
	texi, _ := strconv.Atoi(tex)

	return facei - 1, texi - 1 // obj indices start at 1
}

func (m *Mesh) loadMesh(p string) {
	s := readFile(p)

	sa := parse(s)

	m.vertexData = m.vertexData[:0] // leak?

	fmt.Println(sa)

	for i := 0; i < len(sa); i++ {
		if sa[i] == "v" {
			if len(sa) > i + 3 {
				vx, _ := strconv.ParseFloat(sa[i+1], 32)
				vy, _ := strconv.ParseFloat(sa[i+2], 32)
				vz, _ := strconv.ParseFloat(sa[i+3], 32)

				m.rawVertexData = append(m.rawVertexData, (float32)(vx))
				m.rawVertexData = append(m.rawVertexData, (float32)(vy))
				m.rawVertexData = append(m.rawVertexData, (float32)(vz))
				m.rawVertexData = append(m.rawVertexData, 1.0)
			}
		}
		if sa[i] == "vt" {
			if len(sa) > i + 2 {
				tx, _ := strconv.ParseFloat(sa[i+1], 32)
				ty, _ := strconv.ParseFloat(sa[i+2], 32)

				m.texData = append(m.texData, (float32)(tx))
				m.texData = append(m.texData, (float32)(ty))
				m.texData = append(m.texData, 1.0)
				m.texData = append(m.texData, 1.0)
			}
		}
		if sa[i] == "f" {
			if len(sa) > i + 3 {
				f0, _ := m.parseFace(sa[i+1])
				f1, _ := m.parseFace(sa[i+2])
				f2, _ := m.parseFace(sa[i+3])

				m.indexData = append(m.indexData, (uint16)(f0))
				m.indexData = append(m.indexData, (uint16)(f1))
				m.indexData = append(m.indexData, (uint16)(f2))
			}
		}
	}
/*
	for i := 0; i < len(m.indexData); i += 3 {
		fmt.Print(m.indexData[i])
		fmt.Print(" ")
		fmt.Print(m.indexData[i+1])
		fmt.Print(" ")
		fmt.Print(m.indexData[i+2])
		fmt.Print("\n")
	}

	for i := 0; i < len(m.rawVertexData); i += 4 {
		fmt.Print(m.rawVertexData[i])
		fmt.Print(" ")
		fmt.Print(m.rawVertexData[i+1])
		fmt.Print(" ")
		fmt.Print(m.rawVertexData[i+2])
		fmt.Print(" ")
		fmt.Print(m.rawVertexData[i+3])
		fmt.Print("\n")
	}*/
}

func (m *Mesh) transformVerts() {
	m.vertexData = m.rawVertexData
}

func (m *Mesh) getNormals() {

}

func (m *Mesh) update() {
	m.transformVerts()

	m.getNormals()

	m.getBoundingSphere()
}

func (m *Mesh) collideMesh(n *Mesh) {

}

func (m *Mesh) getBoundingSphere() {
	vn := 0
	for i := 0; i < len(m.vertexData); i += 4 {
		vp := mgl32.Vec3{m.vertexData[i], m.vertexData[i+1], m.vertexData[i+2]}
		m.bsc = m.bsc.Add(vp)
		vn += 1
	}
	m.bsc = m.bsc.Mul(1.0/(float32)(vn))

	for i := 0; i < len(m.vertexData); i += 4 {
		vp := mgl32.Vec3{m.vertexData[i], m.vertexData[i+1], m.vertexData[i+2]}
		nr := vp.Sub(m.bsc).Len()
		if nr > 0.0 {
			m.bsr = nr
		}
	}
}

func (m *Mesh) render(r *Renderer) {
	r.render(&m.tex, &m.um, &m.rawVertexData, &m.texData, &m.indexData)
}

func (m *Mesh) draw(r *Renderer) {
	r.draw(&m.tex, &m.rawVertexData, &m.texData, &m.indexData)
}
