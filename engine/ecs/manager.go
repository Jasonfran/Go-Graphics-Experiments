package ecs

type standardManager struct {
	idCounter                 uint
	entityIdBin               []EntityID
	entityComponentCollection *EntityComponentCollection
	defaultGenerators         []ComponentGenerator
	systems                   map[SystemGroup]SystemCollection
}

func NewManager() Manager {
	return &standardManager{
		entityIdBin:               []EntityID{},
		entityComponentCollection: NewEntityComponentCollection(),
		defaultGenerators:         []ComponentGenerator{},
		systems:                   map[SystemGroup]SystemCollection{},
	}
}

func (s *standardManager) NewEntity() Entity {
	entity := &entity{
		Id:      s.generateEntityID(),
		manager: s,
	}

	for _, generator := range s.defaultGenerators {
		entity.AddComponent(generator())
	}

	return entity
}

func (s *standardManager) AddComponent(entity Entity, component Component) {
	s.entityComponentCollection.Add(entity, component)
}

func (s *standardManager) RemoveComponent(entity Entity, tag ComponentTag) {
	s.entityComponentCollection.Remove(entity, tag)
}

func (s *standardManager) GetComponent(entity Entity, tag ComponentTag) (Component, error) {
	return s.entityComponentCollection.Get(entity, tag)
}

func (s *standardManager) GetEntitiesWithComponents(tags ...ComponentTag) EntityIterable {
	return s.entityComponentCollection.GetEntitiesWithComponents(tags...)
}

func (s *standardManager) DestroyEntity(entity Entity) {
	s.entityIdBin = append(s.entityIdBin, entity.ID())
	s.entityComponentCollection.RemoveEntity(entity)
}

func (s *standardManager) SetDefaultComponents(generators ...ComponentGenerator) {
	s.defaultGenerators = append(s.defaultGenerators, generators...)
}

func (s *standardManager) AddSystem(group SystemGroup, system System) {
	systems, ok := s.systems[group]
	if !ok {
		systems = []System{}
	}

	systems = append(systems, system)
	s.systems[group] = systems
}

func (s *standardManager) GetSystemGroup(group SystemGroup) SystemCollection {
	systems, ok := s.systems[group]
	if !ok {
		return []System{}
	}

	return systems
}

func (s *standardManager) generateEntityID() EntityID {
	var id EntityID
	if len(s.entityIdBin) > 0 {
		id = s.entityIdBin[len(s.entityIdBin)-1]
	} else {
		id = EntityID(s.idCounter)
		s.idCounter++
	}
	return id
}
