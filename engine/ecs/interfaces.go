package ecs

import "fmt"

type ComponentTag uint

type ComponentGenerator func() Component

type Component interface {
	fmt.Stringer
	Tag() ComponentTag
}

type EntityID uint

type Entity interface {
	ID() EntityID
	AddComponent(component Component)
	GetComponent(tag ComponentTag) (Component, error)
	Destroy()
}

type Manager interface {
	NewEntity() Entity
	AddComponent(entity Entity, component Component)
	RemoveComponent(entity Entity, tag ComponentTag)
	GetComponent(entity Entity, tag ComponentTag) (Component, error)
	GetEntitiesWithComponents(tags ...ComponentTag) EntityIterable
	DestroyEntity(entity Entity)
	SetDefaultComponents(generators ...ComponentGenerator)
	AddSystem(group SystemGroup, system System)
	GetSystemGroup(group SystemGroup) SystemCollection
}

type SystemGroup uint

const (
	PlayerSystemGroup SystemGroup = iota
	EngineSystemGroup
)

type SystemCollection []System

func (c SystemCollection) Init() {
	for _, system := range c {
		system.Init()
	}
}

func (c SystemCollection) Update(delta float32) {
	for _, system := range c {
		system.Update(delta)
	}
}

func (c SystemCollection) LateUpdate(delta float32) {
	for _, system := range c {
		system.LateUpdate(delta)
	}
}

func (c SystemCollection) Shutdown() {
	for _, system := range c {
		system.Shutdown()
	}
}

type System interface {
	Init()
	Update(delta float32)
	LateUpdate(delta float32)
	Shutdown()
}
