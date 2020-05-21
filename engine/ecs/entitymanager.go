package ecs

import (
	"GraphicsStuff/engine"
)

type EntityManager struct {
	idCounter                 uint
	entityIdBin               []engine.EntityID
	entityComponentCollection *EntityComponentCollection
	defaultGenerators         []engine.ComponentGenerator
	systems                   map[engine.SystemGroup]engine.SystemCollection
}

func NewStandardEntityManager() engine.IEntityManager {
	return &EntityManager{
		entityIdBin:               []engine.EntityID{},
		entityComponentCollection: NewEntityComponentCollection(),
		defaultGenerators:         []engine.ComponentGenerator{},
		systems:                   map[engine.SystemGroup]engine.SystemCollection{},
	}
}

func (s *EntityManager) NewEntity() engine.IEntity {
	entity := &Entity{
		Id:      s.generateEntityID(),
		manager: s,
	}

	for _, generator := range s.defaultGenerators {
		entity.AddComponent(generator())
	}

	return entity
}

func (s *EntityManager) AddComponent(entity engine.IEntity, component engine.IComponent) {
	s.entityComponentCollection.Add(entity, component)
}

func (s *EntityManager) RemoveComponent(entity engine.IEntity, tag engine.ComponentTag) {
	s.entityComponentCollection.Remove(entity, tag)
}

func (s *EntityManager) GetComponent(entity engine.IEntity, tag engine.ComponentTag) (engine.IComponent, error) {
	return s.entityComponentCollection.Get(entity, tag)
}

func (s *EntityManager) HasComponent(entity engine.IEntity, tag engine.ComponentTag) bool {
	return s.entityComponentCollection.HasComponent(entity, tag)
}

func (s *EntityManager) GetEntitiesFromQuery(query engine.IQuery) engine.EntityIterable {
	return s.entityComponentCollection.GetEntitiesFromQuery(query)
}

func (s *EntityManager) DestroyEntity(entity engine.IEntity) {
	s.entityIdBin = append(s.entityIdBin, entity.ID())
	s.entityComponentCollection.RemoveEntity(entity)
}

func (s *EntityManager) SetDefaultComponents(generators ...engine.ComponentGenerator) {
	s.defaultGenerators = append(s.defaultGenerators, generators...)
}

func (s *EntityManager) AddSystem(group engine.SystemGroup, system engine.ISystem) {
	systems, ok := s.systems[group]
	if !ok {
		systems = []engine.ISystem{}
	}

	systems = append(systems, system)
	s.systems[group] = systems
}

func (s *EntityManager) AddSystems(group engine.SystemGroup, systems ...engine.ISystem) {
	for _, system := range systems {
		s.AddSystem(group, system)
	}
}

func (s *EntityManager) GetSystemGroup(group engine.SystemGroup) engine.SystemCollection {
	systems, ok := s.systems[group]
	if !ok {
		return []engine.ISystem{}
	}

	return systems
}

func (s *EntityManager) generateEntityID() engine.EntityID {
	var id engine.EntityID
	if len(s.entityIdBin) > 0 {
		id = s.entityIdBin[len(s.entityIdBin)-1]
	} else {
		id = engine.EntityID(s.idCounter)
		s.idCounter++
	}
	return id
}
