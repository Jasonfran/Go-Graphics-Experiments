package components

import "GraphicsStuff/engine/ecs"

type Camera struct {
	FOV         float32
	Renderables []ecs.Entity
}

func (c *Camera) String() string {
	return "Camera"
}

func (c *Camera) Tag() ecs.ComponentTag {
	return CameraComponentTag
}
