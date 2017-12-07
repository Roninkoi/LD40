package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

type Texture struct {
	w uint16
	h uint16

	p string

	tex *js.Object
}

func (t *Texture) loadTexture(gl *webgl.Context, p string) {
	t.tex = gl.CreateTexture()
	t.p = p

	//var img *js.Object
	img := &js.Object{}

	img = js.Global.Get("document").Call("createElement", "img")
	//img = js.Global.Get("document").Call("getElementById", "img")
	img.Set("crossOrigin", "Anonymous")

	imgLoaded := make(chan struct{})

	img.Call("addEventListener", "load", func() {close(imgLoaded)})
	img.Set("src", p)
	<-imgLoaded

	gl.BindTexture(gl.TEXTURE_2D, t.tex)

	gl.PixelStorei(gl.UNPACK_FLIP_Y_WEBGL, 1)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.RGBA, gl.UNSIGNED_BYTE, img)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
}
