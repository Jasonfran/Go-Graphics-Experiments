package components

import (
	"GraphicsStuff/engine/ecs"

	"github.com/go-gl/mathgl/mgl32"
)

type Material struct {
	Colour mgl32.Vec3
}

func (m *Material) String() string {
	return "Material"
}

func (m *Material) Tag() ecs.ComponentTag {
	return MaterialComponentTag
}
