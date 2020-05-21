package components

import "GraphicsStuff/engine"

type ParentComponent struct {
	Value engine.IEntity
}

func (p *ParentComponent) String() string {
	return "ParentComponent"
}

func (p *ParentComponent) Tag() engine.ComponentTag {
	return ParentComponentTag
}

func NewParentComponent(parent engine.IEntity) *ParentComponent {
	return &ParentComponent{
		Value: parent,
	}
}
