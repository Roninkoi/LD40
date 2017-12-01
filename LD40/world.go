package main

type World struct {
	testmesh Mesh
}

func (w *World) draw(r *Renderer) {
	//for i := 0; i < 10000; i++ {
		w.testmesh.draw(r)
	//}
}

func (w *World) loadWorld() {
	w.testmesh.loadMesh("ico.obj")
}
