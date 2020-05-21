package components

import (
	"GraphicsStuff/engine"
)

type Camera struct {
	FOV         float32
	Renderables []engine.IEntity
}

func (c *Camera) String() string {
	return "Camera"
}

func (c *Camera) Tag() engine.ComponentTag {
	return CameraComponentTag
}
