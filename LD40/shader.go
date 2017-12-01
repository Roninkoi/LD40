package main

import (
	"github.com/gopherjs/webgl"
	"github.com/gopherjs/gopherjs/js"
	"net/http"
	"io/ioutil"
	"fmt"
)

type Shader struct {
	vs_source string
	fs_source string

	vs *js.Object
	fs *js.Object

	program *js.Object
}

func (s *Shader) compile(gl *webgl.Context, shaderType int, source string) *js.Object {
	newShader := gl.CreateShader(shaderType)
	gl.ShaderSource(newShader, source)
	gl.CompileShader(newShader)

	fmt.Println(gl.GetShaderInfoLog(newShader))

	return newShader
}

func (s *Shader) link(gl *webgl.Context) {
	s.program = gl.CreateProgram()
	gl.AttachShader(s.program, s.vs)
	gl.AttachShader(s.program, s.fs)

	gl.LinkProgram(s.program)
}

func readFile(p string) string {
	var returns string

	var client http.Client
	resp, err := client.Get(p)
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		returns = string(bodyBytes)
	}

	return returns
}

func (s *Shader) loadShader(gl *webgl.Context, vsPath string, fsPath string) {
	s.vs_source = readFile(vsPath)
	s.fs_source = readFile(fsPath)

	s.vs = s.compile(gl, gl.VERTEX_SHADER, s.vs_source)
	s.fs = s.compile(gl, gl.FRAGMENT_SHADER, s.fs_source)

	s.link(gl)
}
