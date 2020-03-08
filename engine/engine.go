package engine

import (
	"GraphicsStuff/engine/ecs/ecsmanager"
)

var (
	ECSManager      = ecsmanager.NewECSManager()
	EventDispatcher = NewDispatcher()
	InputManager    = NewManager()
)
