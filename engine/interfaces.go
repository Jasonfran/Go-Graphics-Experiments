package engine

import (
	"GraphicsStuff/engine/data"
	"fmt"

	"github.com/willf/bitset"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type ComponentTag uint

type ComponentGenerator func() IComponent

type IComponent interface {
	fmt.Stringer
	Tag() ComponentTag
}

type EntityID uint

type IEntity interface {
	ID() EntityID
	AddComponent(component IComponent)
	GetComponent(tag ComponentTag) (IComponent, error)
	HasComponent(tag ComponentTag) bool
	Destroy()
}

type EntityIterable [][]IEntity

func (i EntityIterable) Each(f func(entity IEntity)) {
	for _, entityList := range i {
		for _, entity := range entityList {
			f(entity)
		}
	}
}

func (i EntityIterable) Len() int {
	length := 0
	for _, entityList := range i {
		length += len(entityList)
	}

	return length
}

type IEntityManager interface {
	NewEntity() IEntity
	AddComponent(entity IEntity, component IComponent)
	RemoveComponent(entity IEntity, tag ComponentTag)
	GetComponent(entity IEntity, tag ComponentTag) (IComponent, error)
	HasComponent(entity IEntity, tag ComponentTag) bool
	GetEntitiesFromQuery(query IQuery) EntityIterable
	DestroyEntity(entity IEntity)
	SetDefaultComponents(generators ...ComponentGenerator)
	AddSystem(group SystemGroup, system ISystem)
	AddSystems(group SystemGroup, systems ...ISystem)
	GetSystemGroup(group SystemGroup) SystemCollection
}

type SystemGroup uint

const (
	PlayerSystemGroup SystemGroup = iota
	EngineSystemGroup
)

type SystemCollection []ISystem

func (c SystemCollection) Init(context EngineContext) {
	for _, system := range c {
		system.Init(context)
	}
}

func (c SystemCollection) Update(context EngineContext, delta float32) {
	for _, system := range c {
		system.Update(context, delta)
	}
}

func (c SystemCollection) LateUpdate(context EngineContext, delta float32) {
	for _, system := range c {
		system.LateUpdate(context, delta)
	}
}

func (c SystemCollection) Shutdown(context EngineContext) {
	for _, system := range c {
		system.Shutdown(context)
	}
}

type ISystem interface {
	Init(context EngineContext)
	Update(context EngineContext, delta float32)
	LateUpdate(context EngineContext, delta float32)
	Shutdown(context EngineContext)
}
type IInputManager interface {
	KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
	Update()
	Pressed(key glfw.Key) bool
	Held(key glfw.Key) bool
}
type IResourceManager interface {
	LoadModel(path string) (*data.Model, error)
	ModelToEntity(context EngineContext, model *data.Model) IEntity
}

type IQuery interface {
	InclusionSet() *bitset.BitSet
	ExclusionSet() *bitset.BitSet
}
