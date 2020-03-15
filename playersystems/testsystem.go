package playersystems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
	"GraphicsStuff/engine/ecs"
	"log"
	"time"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type TestPlayerSystem struct {
}

func (t *TestPlayerSystem) Init() {
	log.Println("TestPlayerSystem Init")
	e := engine.ECSManager.NewEntity()
	e.AddComponent(&components.MeshRenderer{Mesh: "Cube"})
	e.AddComponent(&components.Material{Colour: mgl32.Vec3{1, 0, 1}})
	e.AddComponent(&components.PlayerComponent{})
}

func (t *TestPlayerSystem) Update(delta float32) {
	start := time.Now()
	defer func() {
		log.Println("Test system: ", time.Since(start))
	}()

	engine.ECSManager.GetEntitiesWithComponents(components.PlayerComponentTag).Each(func(entity ecs.Entity) {
		entityTransform, _ := components.GetTransformComponent(entity)
		if engine.InputManager.Held(glfw.KeyW) {
			entityTransform.Translate(0, 0, 1*delta)
		}

		if engine.InputManager.Held(glfw.KeyS) {
			entityTransform.Translate(0, 0, -1*delta)
		}

		if engine.InputManager.Held(glfw.KeyD) {
			entityTransform.Translate(-1*delta, 0, 0)
		}

		if engine.InputManager.Held(glfw.KeyA) {
			entityTransform.Translate(1*delta, 0, 0)
		}

		if engine.InputManager.Pressed(glfw.KeyQ) {
			bullet := engine.ECSManager.NewEntity()
			bulletTransform, _ := components.GetTransformComponent(bullet)
			bulletTransform.Pos = entityTransform.Pos
			bullet.AddComponent(&components.MeshRenderer{Mesh: "Cube"})
			bullet.AddComponent(&components.Material{Colour: mgl32.Vec3{1, 0, 0}})
			bullet.AddComponent(&components.PhysicsComponent{Velocity: mgl32.Vec3{0, 0, 1}})
		}

		if engine.InputManager.Held(glfw.KeyE) {
			time.Sleep(50 * time.Millisecond)
		}
	})
}

func (t *TestPlayerSystem) LateUpdate(delta float32) {
	//log.Println("TestPlayerSystem LateUpdate")
}

func (t *TestPlayerSystem) Shutdown() {
	log.Println("TestPlayerSystem Shutdown")
}

func NewTestPlayerSystem() *TestPlayerSystem {
	return &TestPlayerSystem{}
}
