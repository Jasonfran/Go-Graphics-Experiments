package engine

import (
	"GraphicsStuff/engine/ecs"
)

var (
	ECSManager      = ecs.NewManager()
	EventDispatcher = NewDispatcher()
	InputManager    = NewManager()
)
