package components

import (
	"GraphicsStuff/engine/ecs"
)

type PlayerComponent struct {
}

func (p *PlayerComponent) String() string {
	return "PlayerComponent"
}

func (p *PlayerComponent) Tag() ecs.ComponentTag {
	return PlayerComponentTag
}
