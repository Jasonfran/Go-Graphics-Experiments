package components

import (
	"GraphicsStuff/engine"
)

type ChildrenComponent struct {
	children []engine.IEntity
}

func (c *ChildrenComponent) AddChild(entity engine.IEntity) {
	c.children = append(c.children, entity)
}

func (c *ChildrenComponent) GetChildren() []engine.IEntity {
	return c.children
}

func (c *ChildrenComponent) String() string {
	return "ChildrenComponent"
}

func (c *ChildrenComponent) Tag() engine.ComponentTag {
	return ChildrenComponentTag
}

func NewChildrenComponent() *ChildrenComponent {
	return &ChildrenComponent{children: []engine.IEntity{}}
}
