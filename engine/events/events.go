package events

import "GraphicsStuff/engine"

const (
	WindowResizedEvent engine.EventType = iota
)

type WindowResizedEventData struct {
	Width  int
	Height int
}
