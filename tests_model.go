// This is a simple model for your tests
package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
)

// Sets the category and name of your test in the global map "TestMap"
// The category name choosen here starts with a "|" so it shows as the
// last category in list. Change "model" to the name of your test.
func init() {
	TestMap["|tests|.model"] = &testsModel{}
}

// This is your test object. You can store state here.
// By convention and to avoid conflict with other demo/tests name it
// using your test category and name.
type testsModel struct {
	grid *graphic.GridHelper // Pointer to a GridHelper created in 'Initialize'
}

// This method will be called once when the test is selected from the list.
// ctx is a pointer to a Context structure built by the main program.
// It allows access to several important object such as the scene (ctx.Scene),
// camera (ctx.Camera) and the window (ctx.Win) among others.
// You can build your scene adding your objects to the ctx.Scene.
func (t *testsModel) Initialize(ctx *Context) {

	// Show axis helper
	ah := graphic.NewAxisHelper(1.0)
	ctx.Scene.Add(ah)

	// Creates a grid helper and saves its pointer in the test state
	t.grid = graphic.NewGridHelper(50, 1, &math32.Color{0.4, 0.4, 0.4})
	ctx.Scene.Add(t.grid)

	// Changes the camera position
	ctx.Camera.GetCamera().SetPosition(0, 4, 10)
}

// This method will be called at every frame
// You can animate your objects here.
func (t *testsModel) Render(ctx *Context) {

	// Rotate the grid, just for show.
	t.grid.AddRotationY(0.005)
}
