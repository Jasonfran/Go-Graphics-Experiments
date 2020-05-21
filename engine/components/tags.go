package components

import (
	"GraphicsStuff/engine"
)

const (
	TransformComponentTag engine.ComponentTag = iota + 1
	CameraComponentTag
	ModelComponentTag
	MaterialComponentTag
	PlayerComponentTag
	PhysicsComponentTag
	MeshComponentTag
	ParentComponentTag
	ChildrenComponentTag
)
