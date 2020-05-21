package ecs

import (
	"GraphicsStuff/engine"
	"fmt"
)

type componentLookup struct {
	EntityID     engine.EntityID
	ComponentTag engine.ComponentTag
}

type EntityComponentCollection struct {
	entityArchetypes         map[engine.EntityID]Archetype
	archetypeEntities        map[Archetype][]engine.IEntity
	archetypeEntitiesIndexes map[engine.EntityID]int

	entityComponents map[componentLookup]engine.IComponent
}

func NewEntityComponentCollection() *EntityComponentCollection {
	return &EntityComponentCollection{
		entityArchetypes:         map[engine.EntityID]Archetype{},
		archetypeEntities:        map[Archetype][]engine.IEntity{},
		archetypeEntitiesIndexes: map[engine.EntityID]int{},

		entityComponents: map[componentLookup]engine.IComponent{},
	}
}

func (c *EntityComponentCollection) Add(entity engine.IEntity, component engine.IComponent) {
	// add component
	cLookup := componentLookup{
		EntityID:     entity.ID(),
		ComponentTag: component.Tag(),
	}

	c.entityComponents[cLookup] = component

	// update archetype
	archetype, ok := c.entityArchetypes[entity.ID()]
	if ok {
		// remove from existing archetype entities
		c.removeFromArchtype(archetype, entity)

		// change archetype
		archetype = archetype.AddType(component.Tag())
		c.entityArchetypes[entity.ID()] = archetype
	} else {
		// create archetype
		archetype = NewArchetype(component.Tag())
		c.entityArchetypes[entity.ID()] = archetype
	}

	// add to new archetype entities
	archetypeEntities := c.archetypeEntities[archetype]
	c.archetypeEntities[archetype] = append(archetypeEntities, entity)
	index := len(c.archetypeEntities[archetype]) - 1
	c.archetypeEntitiesIndexes[entity.ID()] = index
}

func (c *EntityComponentCollection) removeFromArchtype(archetype Archetype, entity engine.IEntity) {
	archetypeEntities := c.archetypeEntities[archetype]
	lastEntity := archetypeEntities[len(archetypeEntities)-1]
	index := c.archetypeEntitiesIndexes[entity.ID()]
	archetypeEntities[index] = lastEntity
	c.archetypeEntities[archetype] = archetypeEntities[:len(archetypeEntities)-1]
	if len(c.archetypeEntities[archetype]) == 0 {
		delete(c.archetypeEntities, archetype)
	}
}

func (c *EntityComponentCollection) Remove(entity engine.IEntity, tag engine.ComponentTag) {
	cLookup := componentLookup{
		EntityID:     entity.ID(),
		ComponentTag: tag,
	}

	delete(c.entityComponents, cLookup)

	archetype, ok := c.entityArchetypes[entity.ID()]
	if ok {
		// remove from existing archetype entities
		c.removeFromArchtype(archetype, entity)

		// change archetype
		archetype = archetype.RemoveType(tag)
		c.entityArchetypes[entity.ID()] = archetype

		// add to new archetype entities
		archetypeEntities := c.archetypeEntities[archetype]
		c.archetypeEntities[archetype] = append(archetypeEntities, entity)
		index := len(c.archetypeEntities[archetype]) - 1
		c.archetypeEntitiesIndexes[entity.ID()] = index
	}
}

func (c *EntityComponentCollection) RemoveEntity(entity engine.IEntity) {
	archetype, ok := c.entityArchetypes[entity.ID()]
	if !ok {
		return
	}

	for _, tag := range archetype.GetTypes() {
		c.Remove(entity, tag)
	}

	// remove from existing archetype entities
	archetypeEntities := c.archetypeEntities[archetype]
	lastEntity := archetypeEntities[len(archetypeEntities)-1]
	index := c.archetypeEntitiesIndexes[entity.ID()]
	archetypeEntities[index] = lastEntity
	archetypeEntities = archetypeEntities[:len(archetypeEntities)-1]

	delete(c.archetypeEntitiesIndexes, entity.ID())
	delete(c.entityArchetypes, entity.ID())
}

func (c *EntityComponentCollection) Get(entity engine.IEntity, tag engine.ComponentTag) (engine.IComponent, error) {
	cLookup := componentLookup{
		EntityID:     entity.ID(),
		ComponentTag: tag,
	}

	comp, ok := c.entityComponents[cLookup]
	if !ok {
		return nil, fmt.Errorf("could not find component (%v) for entity (%v)", entity.ID(), tag)
	}

	return comp, nil
}

func (c *EntityComponentCollection) HasComponent(entity engine.IEntity, tag engine.ComponentTag) bool {
	cLookup := componentLookup{
		EntityID:     entity.ID(),
		ComponentTag: tag,
	}

	_, ok := c.entityComponents[cLookup]
	return ok
}

func (c *EntityComponentCollection) GetEntitiesFromQuery(query engine.IQuery) engine.EntityIterable {
	foundEntities := make([][]engine.IEntity, 0, query.InclusionSet().Count())
	for archetype, entities := range c.archetypeEntities {
		if archetype.Satisfies(query) {
			foundEntities = append(foundEntities, entities)
		}
	}

	return foundEntities
}
