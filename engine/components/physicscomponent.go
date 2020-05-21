package components

import (
	"GraphicsStuff/engine"

	"github.com/go-gl/mathgl/mgl32"
)

type PhysicsComponent struct {
	Velocity mgl32.Vec3
}

func (p *PhysicsComponent) String() string {
	return "PhysicsComponent"
}

func (p *PhysicsComponent) Tag() engine.ComponentTag {
	return PhysicsComponentTag
}
