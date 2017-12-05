package main

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	LEFT  = 37
	UP    = 38
	RIGHT = 39
	DOWN  = 40

	SPACE = 32

	A = 65
	D = 68
	F = 70
	Q = 81
	R = 82
	S = 83
	W = 87
)

type Input struct {
	keys [256]bool
	keyobjs [256]*js.Object

	space_pressed bool
	right_pressed bool
	left_pressed bool
}

func (i *Input) init() {
	for j := 0; j < 256; j++ {
		i.keyobjs[j] = new(js.Object)
	}
}

func (i *Input) resetKeys() {
	for j := 0; j < len(i.keys); j++ {
		i.keys[j] = false
	}
}

func (i *Input) setKey(code int) {
	i.keys[code] = true
}

func (i *Input) getKeys() {
	for j := 0; j < 256; j++ {
		i.keys[j] = i.keyobjs[j].Get("kc").Bool()
	}
}

func inputHandlerDown(event *js.Object, input *Input) {
	keycode := event.Get("keyCode").Int()

	if keycode < 256 {
		input.keyobjs[keycode].Set("kc", false)
	}
}
func inputHandlerUp(event *js.Object, input *Input) {
	keycode := event.Get("keyCode").Int()

	if keycode < 256 {
		input.keyobjs[keycode].Set("kc", false)
	}
}

func (i *Input) keyHandler() {
	js.Global.Call("addEventListener", "keydown", func(event *js.Object) {
		keycode := event.Get("keyCode").Int()

		if keycode < 256 {
			i.keyobjs[keycode].Set("kc", true)
		}
	}, true)

	js.Global.Call("addEventListener", "keyup", func(event *js.Object) {
		keycode := event.Get("keyCode").Int()

		if keycode < 256 {
			i.keyobjs[keycode].Set("kc", false)
		}
	}, true)
}
