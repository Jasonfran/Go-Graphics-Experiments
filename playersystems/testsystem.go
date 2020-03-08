package playersystems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/ecs/components"
	"GraphicsStuff/engine/ecs/ecsmanager"
	"log"
	"time"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type TestPlayerSystem struct {
	*ecsmanager.SystemEntityCollection
}

func (t *TestPlayerSystem) Init() {
	log.Println("TestPlayerSystem Init")
	e := engine.ECSManager.NewEntity()
	components.AddMeshRendererComponent(e, &components.MeshRenderer{Mesh: "Cube"})
	components.AddMaterialComponent(e, &components.Material{Colour: mgl32.Vec3{1, 0, 1}})
	components.AddPlayerComponent(e, &components.PlayerComponent{})
}

func (t *TestPlayerSystem) Update(delta float32) {
	start := time.Now()
	defer func() {
		log.Println("Test system: ", time.Since(start))
	}()
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

		if engine.InputManager.Pressed(glfw.KeyQ) {
			bullet := engine.ECSManager.NewEntity()
			bullet.Transform().Pos = entity.Transform().Pos
			components.AddMeshRendererComponent(bullet, &components.MeshRenderer{Mesh: "Cube"})
			components.AddMaterialComponent(bullet, &components.Material{Colour: mgl32.Vec3{1, 0, 0}})
			components.AddPhysicsComponent(bullet, &components.PhysicsComponent{Velocity: mgl32.Vec3{0, 0, 1}})
		}

		if engine.InputManager.Held(glfw.KeyE) {
			time.Sleep(50 * time.Millisecond)
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
	system := &TestPlayerSystem{SystemEntityCollection: ecsmanager.NewSystemEntityCollection()}
	system.SetRequirements(components.PlayerComponentTag)
	return system
}
