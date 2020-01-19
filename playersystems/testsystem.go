package playersystems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/component"
	"GraphicsStuff/engine/ecsmanager"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/go-gl/mathgl/mgl32"
)

type TestPlayerSystem struct {
	*ecsmanager.SystemEntityCollection
}

func (t *TestPlayerSystem) Init() {
	log.Println("TestPlayerSystem Init")
	e := engine.ECSManager.NewEntity()
	e.AddMeshRendererComponent(&component.MeshRenderer{Mesh: "Cube"})
	e.AddMaterialComponent(&component.Material{Colour: mgl32.Vec3{1, 1, 0}})

	//for i := 0; i < 100; i++ {
	//	entity := engine.ECSManager.NewEntity()
	//	entity.AddMeshRendererComponent(&component.MeshRenderer{Mesh: "Cube"})
	//}
}

func (t *TestPlayerSystem) Update(delta float32) {
	for _, entity := range t.Entities() {
		transform := entity.Transform()
		if engine.InputManager.Held(glfw.KeyW) {
			transform.Translate(0, 0, 1*delta)
		}

		if engine.InputManager.Held(glfw.KeyS) {
			transform.Translate(0, 0, -1*delta)
		}

		if engine.InputManager.Held(glfw.KeyD) {
			transform.Translate(-1*delta, 0, 0)
		}

		if engine.InputManager.Held(glfw.KeyA) {
			transform.Translate(1*delta, 0, 0)
		}
	}
}

func (t *TestPlayerSystem) LateUpdate(delta float32) {
	//log.Println("TestPlayerSystem LateUpdate")
}

func (t *TestPlayerSystem) Shutdown() {
	log.Println("TestPlayerSystem Shutdown")
}

func NewTestPlayerSystem() *TestPlayerSystem {
	return &TestPlayerSystem{SystemEntityCollection: ecsmanager.NewSystemEntityCollection()}
}
