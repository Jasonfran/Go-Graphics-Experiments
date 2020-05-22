package playersystems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
	"GraphicsStuff/engine/data"
	"log"
	"time"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type TestPlayerSystem struct {
	bulletModel *data.Model
}

func (t *TestPlayerSystem) Init(context engine.EngineContext) {
	log.Println("TestPlayerSystem Init")
	model, err := context.ResourceManager.LoadModel(`C:\Users\jason\Downloads\mosin_nagant_3-line_rifle\scene.gltf`)
	if err != nil {
		log.Fatal(err)
	}
	t.bulletModel = model

	e := context.ResourceManager.ModelToEntity(context, model)
	e.AddComponent(&components.PlayerComponent{})

	//child := context.ResourceManager.ModelToEntity(context, model)
	//child.AddComponent(components.NewParentComponent(e))
	//childTransform, _ := components.GetTransformComponent(child)
	//childTransform.SetPos(1, -2, 0)
	//childTransform.SetScale(0.5, 0.5, 0.5)
}

func (t *TestPlayerSystem) Update(context engine.EngineContext, delta float32) {
	//start := time.Now()
	//defer func() {
	//	log.Println("Test system: ", time.Since(start))
	//}()

	context.EntityManager.GetEntitiesFromQuery(components.Includes(components.PlayerComponentTag)).Each(func(entity engine.IEntity) {
		entityTransform, _ := components.GetTransformComponent(entity)
		if context.InputManager.Held(glfw.KeyW) {
			entityTransform.Translate(0, 0, 1*delta)
		}

		if context.InputManager.Held(glfw.KeyS) {
			entityTransform.Translate(0, 0, -1*delta)
		}

		if context.InputManager.Held(glfw.KeyD) {
			entityTransform.Translate(-1*delta, 0, 0)
		}

		if context.InputManager.Held(glfw.KeyA) {
			entityTransform.Translate(1*delta, 0, 0)
		}

		if context.InputManager.Held(glfw.KeySpace) {
			entityTransform.Translate(0, 1*delta, 0)
		}

		if context.InputManager.Held(glfw.KeyLeftControl) {
			entityTransform.Translate(0, -1*delta, 0)
		}

		if context.InputManager.Held(glfw.KeyUp) {
			entityTransform.SetScaleVec3(entityTransform.GetScale().Add(mgl32.Vec3{1, 1, 1}.Mul(delta)))
		}

		if context.InputManager.Held(glfw.KeyDown) {
			entityTransform.SetScaleVec3(entityTransform.GetScale().Sub(mgl32.Vec3{1, 1, 1}.Mul(delta)))
		}

		if context.InputManager.Pressed(glfw.KeyQ) {
			bullet := context.ResourceManager.ModelToEntity(context, t.bulletModel)
			bulletTransform, _ := components.GetTransformComponent(bullet)
			bulletTransform.SetPosVec3(entityTransform.GetPos())
			bullet.AddComponent(&components.PhysicsComponent{Velocity: mgl32.Vec3{0, 0, 1}})
		}

		if context.InputManager.Held(glfw.KeyE) {
			time.Sleep(50 * time.Millisecond)
		}
	})
}

func (t *TestPlayerSystem) LateUpdate(context engine.EngineContext, delta float32) {
	//log.Println("TestPlayerSystem LateUpdate")
}

func (t *TestPlayerSystem) Shutdown(context engine.EngineContext) {
	log.Println("TestPlayerSystem Shutdown")
}

func NewTestPlayerSystem() *TestPlayerSystem {
	return &TestPlayerSystem{}
}
