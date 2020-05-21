package components

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/data"
)

type ModelComponent struct {
	Model *data.Model
}

func (m *ModelComponent) String() string {
	return "ModelComponent"
}

func (m *ModelComponent) Tag() engine.ComponentTag {
	return ModelComponentTag
}
