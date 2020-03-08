package playersystems

import (
	"GraphicsStuff/engine/ecs/components"
	"GraphicsStuff/engine/ecs/ecsmanager"
	"log"
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

type PhysicsSystem struct {
	*ecsmanager.SystemEntityCollection
}

func NewPhysicsSystem() *PhysicsSystem {
	system := &PhysicsSystem{SystemEntityCollection: ecsmanager.NewSystemEntityCollection()}
	system.SetRequirements(components.PhysicsComponentTag)
	return system
}

func (s *PhysicsSystem) Init() {

}

func (s *PhysicsSystem) Update(delta float32) {

}

func (s *PhysicsSystem) LateUpdate(delta float32) {
	start := time.Now()
	defer func() {
		log.Println("Physics system: ", time.Since(start))
	}()
	for _, entity := range s.Entities() {
		transform := entity.Transform()
		physics, err := components.GetPhysicsComponent(entity)
		if err != nil {
			continue
		}

		transform.Translate(physics.Velocity.X(), physics.Velocity.Y(), physics.Velocity.Z())
		physics.Velocity = physics.Velocity.Sub(mgl32.Vec3{1 * physics.Velocity.X() * delta, 1 * physics.Velocity.Y() * delta, 1 * physics.Velocity.Z() * delta})
	}
}

func (s *PhysicsSystem) Shutdown() {

}
