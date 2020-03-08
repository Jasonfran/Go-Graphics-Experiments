package components

import "GraphicsStuff/engine/ecs"

type Camera struct {
	FOV         float32
	Renderables []ecs.Entity
}

func (c *Camera) Name() string {
	return "Camera"
}
