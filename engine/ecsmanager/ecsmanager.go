package ecsmanager

import (
	"GraphicsStuff/engine/cache"
	"GraphicsStuff/engine/component"
	"errors"

	"github.com/willf/bitset"
)

type ECSManager struct {
	idCounter        EntityId
	entities         map[EntityId]*Entity
	entityComponents map[EntityId]map[component.ComponentTag]component.Component
	entityBitsets    map[EntityId]*bitset.BitSet
	bitsetCache      *cache.BitsetCache
	bitsetEntities   *BitsetEntityCollection
	systems          map[SystemGroupType]*SystemGroup
}

func NewECSManager() *ECSManager {
	return &ECSManager{
		entities:         map[EntityId]*Entity{},
		entityComponents: map[EntityId]map[component.ComponentTag]component.Component{},
		entityBitsets:    map[EntityId]*bitset.BitSet{},
		bitsetCache:      cache.Instance,
		bitsetEntities:   NewBitsetEntityCollection(),
		systems: map[SystemGroupType]*SystemGroup{
			PlayerSystemGroup: NewSystemGroup(PlayerSystemGroup),
			EngineSystemGroup: NewSystemGroup(EngineSystemGroup),
		},
	}
}

func (e *ECSManager) NewEntity() *Entity {
	entity := &Entity{
		Id:         e.idCounter,
		ecsManager: e,
	}

	e.idCounter++
	e.entities[entity.Id] = entity
	e.entityBitsets[entity.Id] = e.bitsetCache.New()
	e.entityComponents[entity.Id] = map[component.ComponentTag]component.Component{}

	// All entities should have a transform component
	entity.transform = &component.Transform{}
	return entity
}

func (e *ECSManager) AddComponent(entity *Entity, tag component.ComponentTag, c component.Component) {
	e.entityComponents[entity.Id][tag] = c
	b := e.entityBitsets[entity.Id]
	e.bitsetEntities.Add(entity, e.bitsetCache.New(tag))
	newSet := e.bitsetCache.Append(b, tag)
	e.bitsetEntities.Add(entity, newSet)
	e.entityBitsets[entity.Id] = newSet

	e.GetSystemGroup(PlayerSystemGroup).AddEntity(entity)
	e.GetSystemGroup(EngineSystemGroup).AddEntity(entity)
}

func (e *ECSManager) GetComponent(entity *Entity, tag component.ComponentTag) (component.Component, error) {
	c, ok := e.entityComponents[entity.Id][tag]
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

func (e *ECSManager) GetEntityBitset(entity *Entity) *bitset.BitSet {
	return e.entityBitsets[entity.Id]
}

func (e *ECSManager) GetSystemGroup(group SystemGroupType) *SystemGroup {
	return e.systems[group]
}

func (e *ECSManager) GetEntitiesWithComponents(tags ...component.ComponentTag) []*Entity {
	b := e.bitsetCache.New(tags...)
	return e.bitsetEntities.Get(b)
}

//func (e *ECSManager) removeFromBitsetEntities(entity *Entity, b *bitset.BitSet) {
//	entities, ok := e.bitsetEntities[b]
//	if !ok {
//		return
//	}
//	delete(entities, entity.Id)
//}
//
//func (e *ECSManager) addToBitsetEntities(entity *Entity, b *bitset.BitSet) {
//	entities, ok := e.bitsetEntities[b]
//	if !ok {
//		entities = map[EntityId]*Entity{entity.Id: entity}
//		e.bitsetEntities[b] = entities
//		return
//	}
//
//	entities[entity.Id] = entity
//}
