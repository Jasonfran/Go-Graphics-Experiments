package ecsmanager

import (
	"GraphicsStuff/engine/ecs"

	"github.com/willf/bitset"
)

type BitsetEntityCollection struct {
	bitsetEntities      map[*bitset.BitSet][]ecs.Entity
	bitsetEntityIndexes map[*bitset.BitSet]map[ecs.EntityId]int
}

func NewBitsetEntityCollection() *BitsetEntityCollection {
	return &BitsetEntityCollection{
		bitsetEntities:      map[*bitset.BitSet][]ecs.Entity{},
		bitsetEntityIndexes: map[*bitset.BitSet]map[ecs.EntityId]int{},
	}
}

func (c *BitsetEntityCollection) Add(entity ecs.Entity, b *bitset.BitSet) {
	entities, ok := c.bitsetEntities[b]
	if !ok {
		entities = make([]ecs.Entity, 0, 20)
		c.bitsetEntities[b] = entities
	}

	entities = append(entities, entity)
	c.bitsetEntities[b] = entities

	indexes, ok := c.bitsetEntityIndexes[b]
	if !ok {
		indexes = map[ecs.EntityId]int{entity.Id(): len(entities) - 1}
		c.bitsetEntityIndexes[b] = indexes
		return
	}

	indexes[entity.Id()] = len(entities) - 1
}

func (c *BitsetEntityCollection) Remove(entity ecs.Entity, b *bitset.BitSet) {
	entities, ok := c.bitsetEntities[b]
	if !ok {
		return
	}

	indexes, ok := c.bitsetEntityIndexes[b]
	if !ok {
		return
	}

	index, ok := indexes[entity.Id()]
	if !ok {
		return
	}

	lastEntity := entities[len(entities)-1]
	entities[index] = lastEntity
	indexes[lastEntity.Id()] = index
	c.bitsetEntities[b] = entities[:len(entities)-1]
	delete(c.bitsetEntityIndexes[b], entity.Id())
}

func (c *BitsetEntityCollection) Get(b *bitset.BitSet) []ecs.Entity {
	entities, ok := c.bitsetEntities[b]
	if !ok {
		return []ecs.Entity{}
	}
	return entities
}
