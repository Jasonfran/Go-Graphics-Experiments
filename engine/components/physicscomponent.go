package components

import (
	"GraphicsStuff/engine/ecs"

	"github.com/go-gl/mathgl/mgl32"
)

type PhysicsComponent struct {
	Velocity mgl32.Vec3
}

func (p *PhysicsComponent) String() string {
	return "PhysicsComponent"
}

func (p *PhysicsComponent) Tag() ecs.ComponentTag {
	return PhysicsComponentTag
}
