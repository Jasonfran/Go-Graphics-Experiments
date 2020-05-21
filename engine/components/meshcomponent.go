package components

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/data"
)

type MeshComponent struct {
	Mesh *data.Mesh
}

func (m *MeshComponent) String() string {
	return "MeshComponent"
}

func (m *MeshComponent) Tag() engine.ComponentTag {
	return MeshComponentTag
}
