package components

import (
	"GraphicsStuff/engine"
)

type PlayerComponent struct {
}

func (p *PlayerComponent) String() string {
	return "PlayerComponent"
}

func (p *PlayerComponent) Tag() engine.ComponentTag {
	return PlayerComponentTag
}
