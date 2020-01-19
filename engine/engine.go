package engine

import (
	"GraphicsStuff/engine/ecsmanager"
	"GraphicsStuff/engine/events"
	"GraphicsStuff/engine/input"
)

var (
	ECSManager      = ecsmanager.NewECSManager()
	EventDispatcher = events.NewDispatcher()
	InputManager    = input.NewManager()
)
