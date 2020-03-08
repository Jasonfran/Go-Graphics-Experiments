package components

import "github.com/go-gl/mathgl/mgl32"

type PhysicsComponent struct {
	Velocity mgl32.Vec3
}

func (p *PhysicsComponent) Name() string {
	return "PhysicsComponent"
}
