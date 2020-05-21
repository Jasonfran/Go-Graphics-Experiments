package playersystems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"

	"github.com/go-gl/mathgl/mgl32"
)

type PhysicsSystem struct {
}

func NewPhysicsSystem() *PhysicsSystem {
	return &PhysicsSystem{}
}

func (s *PhysicsSystem) Init(context engine.EngineContext) {

}

func (s *PhysicsSystem) Update(context engine.EngineContext, delta float32) {

}

func (s *PhysicsSystem) LateUpdate(context engine.EngineContext, delta float32) {
	//start := time.Now()
	//defer func() {
	//	log.Println("Physics system: ", time.Since(start))
	//}()

	context.EntityManager.GetEntitiesFromQuery(components.Includes(components.PhysicsComponentTag)).Each(func(entity engine.IEntity) {
		transform, _ := components.GetTransformComponent(entity)
		physics, err := components.GetPhysicsComponent(entity)
		if err != nil {
			return
		}

		transform.Translate(physics.Velocity.X(), physics.Velocity.Y(), physics.Velocity.Z())
		physics.Velocity = physics.Velocity.Sub(mgl32.Vec3{1 * physics.Velocity.X() * delta, 1 * physics.Velocity.Y() * delta, 1 * physics.Velocity.Z() * delta})
	})
}

func (s *PhysicsSystem) Shutdown(context engine.EngineContext) {

}
