package ecsmanager

import (
	"GraphicsStuff/engine/ecs"
	"errors"

	"github.com/willf/bitset"
)

type ECSManager struct {
	idCounter        ecs.EntityId
	entities         map[ecs.EntityId]ecs.Entity
	entityComponents map[ecs.EntityId]map[ecs.ComponentTag]ecs.Component
	entityBitsets    map[ecs.EntityId]*bitset.BitSet
	bitsetCache      *BitsetCache
	bitsetEntities   *BitsetEntityCollection
	systems          map[SystemGroupType]*SystemGroup
}

func NewECSManager() *ECSManager {
	return &ECSManager{
		entities:         map[ecs.EntityId]ecs.Entity{},
		entityComponents: map[ecs.EntityId]map[ecs.ComponentTag]ecs.Component{},
		entityBitsets:    map[ecs.EntityId]*bitset.BitSet{},
		bitsetCache:      BitsetCacheInstance,
		bitsetEntities:   NewBitsetEntityCollection(),
		systems: map[SystemGroupType]*SystemGroup{
			PlayerSystemGroup: NewSystemGroup(PlayerSystemGroup),
			EngineSystemGroup: NewSystemGroup(EngineSystemGroup),
		},
	}
}

func (e *ECSManager) NewEntity() ecs.Entity {
	entity := &entity{
		id:         e.idCounter,
		ecsManager: e,
	}

	e.idCounter++
	e.entities[entity.Id()] = entity
	e.entityBitsets[entity.Id()] = e.bitsetCache.New()
	e.entityComponents[entity.Id()] = map[ecs.ComponentTag]ecs.Component{}

	// All entities should have a transform component
	entity.transform = &ecs.Transform{}
	return entity
}

func (e *ECSManager) AddComponent(entity ecs.Entity, tag ecs.ComponentTag, c ecs.Component) {
	e.entityComponents[entity.Id()][tag] = c
	b := e.entityBitsets[entity.Id()]
	e.bitsetEntities.Add(entity, e.bitsetCache.New(tag))
	newSet := e.bitsetCache.Append(b, tag)
	e.bitsetEntities.Add(entity, newSet)
	e.entityBitsets[entity.Id()] = newSet

	e.GetSystemGroup(PlayerSystemGroup).AddEntity(entity)
	e.GetSystemGroup(EngineSystemGroup).AddEntity(entity)
}

func (e *ECSManager) GetComponent(entity ecs.Entity, tag ecs.ComponentTag) (ecs.Component, error) {
	c, ok := e.entityComponents[entity.Id()][tag]
	if !ok {
		return nil, errors.New("entity or component not found")
	}

	return c, nil
}

func (e *ECSManager) AddSystem(group SystemGroupType, s System) {
	g, ok := e.systems[group]
	if !ok {
		g = NewSystemGroup(group)
		e.systems[group] = g
	}

	g.Add(s)
}

func (e *ECSManager) GetEntityBitset(entity ecs.Entity) *bitset.BitSet {
	return e.entityBitsets[entity.Id()]
}

func (e *ECSManager) GetSystemGroup(group SystemGroupType) *SystemGroup {
	return e.systems[group]
}

func (e *ECSManager) GetEntitiesWithComponents(tags ...ecs.ComponentTag) []ecs.Entity {
	b := e.bitsetCache.New(tags...)
	return e.bitsetEntities.Get(b)
}
