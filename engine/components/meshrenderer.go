package components

import (
	"GraphicsStuff/engine/ecs"
)

type MeshRenderer struct {
	Mesh string
}

func (m *MeshRenderer) String() string {
	return "MeshRenderer"
}

func (m *MeshRenderer) Tag() ecs.ComponentTag {
	return MeshRendererComponentTag
}
