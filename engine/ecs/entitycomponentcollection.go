package ecs

import (
	"fmt"
)

type componentLookup struct {
	EntityID     EntityID
	ComponentTag ComponentTag
}

type EntityComponentCollection struct {
	entityArchetypes         map[EntityID]Archetype
	archetypeEntities        map[Archetype][]Entity
	archetypeEntitiesIndexes map[EntityID]int

	entityComponents map[componentLookup]Component
}

type EntityIterable [][]Entity

func (i EntityIterable) Each(f func(entity Entity)) {
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

func NewEntityComponentCollection() *EntityComponentCollection {
	return &EntityComponentCollection{
		entityArchetypes:         map[EntityID]Archetype{},
		archetypeEntities:        map[Archetype][]Entity{},
		archetypeEntitiesIndexes: map[EntityID]int{},

		entityComponents: map[componentLookup]Component{},
	}
}

func (c *EntityComponentCollection) Add(entity Entity, component Component) {
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

func (c *EntityComponentCollection) removeFromArchtype(archetype Archetype, entity Entity) {
	archetypeEntities := c.archetypeEntities[archetype]
	lastEntity := archetypeEntities[len(archetypeEntities)-1]
	index := c.archetypeEntitiesIndexes[entity.ID()]
	archetypeEntities[index] = lastEntity
	c.archetypeEntities[archetype] = archetypeEntities[:len(archetypeEntities)-1]
	if len(c.archetypeEntities[archetype]) == 0 {
		delete(c.archetypeEntities, archetype)
	}
}

func (c *EntityComponentCollection) Remove(entity Entity, tag ComponentTag) {
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

func (c *EntityComponentCollection) RemoveEntity(entity Entity) {
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

func (c *EntityComponentCollection) Get(entity Entity, tag ComponentTag) (Component, error) {
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

func (c *EntityComponentCollection) GetEntitiesWithComponents(tags ...ComponentTag) EntityIterable {
	foundEntities := make([][]Entity, 0, len(tags))
	query := NewQuery(tags...)
	for archetype, entities := range c.archetypeEntities {
		if archetype.Satisfies(query) {
			foundEntities = append(foundEntities, entities)
		}
	}

	return foundEntities
}
