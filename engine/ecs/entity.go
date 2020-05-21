package ecs

import (
	"GraphicsStuff/engine"
)

type Entity struct {
	Id      engine.EntityID
	manager engine.IEntityManager
}

func (e *Entity) ID() engine.EntityID {
	return e.Id
}

func (e *Entity) AddComponent(component engine.IComponent) {
	e.manager.AddComponent(e, component)
}

func (e *Entity) GetComponent(tag engine.ComponentTag) (engine.IComponent, error) {
	return e.manager.GetComponent(e, tag)
}

func (e *Entity) HasComponent(tag engine.ComponentTag) bool {
	return e.manager.HasComponent(e, tag)
}

func (e *Entity) Destroy() {
	e.manager.DestroyEntity(e)
}
