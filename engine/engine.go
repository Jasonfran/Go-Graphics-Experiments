package engine

import "GraphicsStuff/engine/ecs"

var (
	ECSManager      = ecs.NewECSManager()
	EventDispatcher = NewDispatcher()
	InputManager    = NewManager()
)
