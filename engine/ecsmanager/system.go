package ecsmanager

import (
	"GraphicsStuff/engine/cache"
	"GraphicsStuff/engine/component"

	"github.com/willf/bitset"
)

type SystemGroupType uint

const (
	EngineSystemGroup SystemGroupType = iota
	PlayerSystemGroup
)

type System interface {
	Init()
	Update(delta float32)
	LateUpdate(delta float32)
	Shutdown()
	AddEntity(e *Entity)
	RemoveEntity(id EntityId)
	GetRequirements() *bitset.BitSet
	SetRequirements(tags ...component.ComponentTag)
}

type SystemEntityCollection struct {
	bitsetCache  *cache.BitsetCache
	entities     map[EntityId]*Entity
	requirements *bitset.BitSet
}

func NewSystemEntityCollection() *SystemEntityCollection {
	return &SystemEntityCollection{
		bitsetCache:  cache.Instance,
		entities:     map[EntityId]*Entity{},
		requirements: cache.Instance.New(),
	}
}

func (c *SystemEntityCollection) AddEntity(e *Entity) {
	c.entities[e.Id] = e
}

func (c *SystemEntityCollection) RemoveEntity(id EntityId) {
	delete(c.entities, id)
}

func (c *SystemEntityCollection) GetRequirements() *bitset.BitSet {
	return c.requirements
}

func (c *SystemEntityCollection) SetRequirements(tags ...component.ComponentTag) {
	c.requirements = c.bitsetCache.New(tags...)
}

func (c *SystemEntityCollection) Entities() []*Entity {
	entities := make([]*Entity, 0, len(c.entities))
	for _, entity := range c.entities {
		entities = append(entities, entity)
	}

	return entities
}

type SystemGroup struct {
	groupType SystemGroupType
	systems   []System
}

func NewSystemGroup(group SystemGroupType) *SystemGroup {
	return &SystemGroup{
		groupType: group,
		systems:   []System{},
	}
}

func (g *SystemGroup) Add(s System) {
	g.systems = append(g.systems, s)
}

func (g *SystemGroup) Init() {
	for _, system := range g.systems {
		system.Init()
	}
}

func (g *SystemGroup) Update(delta float32) {
	for _, system := range g.systems {
		system.Update(delta)
	}
}

func (g *SystemGroup) LateUpdate(delta float32) {
	for _, system := range g.systems {
		system.LateUpdate(delta)
	}
}

func (g *SystemGroup) Shutdown() {
	for _, system := range g.systems {
		system.Shutdown()
	}
}

func (g *SystemGroup) AddEntity(e *Entity) {
	for _, system := range g.systems {
		req := system.GetRequirements()
		if e.ComponentBitset().IsSuperSet(req) {
			system.AddEntity(e)
		}
	}
}

func (g *SystemGroup) RemoveEntity(id EntityId) {
	panic("implement me")
}

func (g *SystemGroup) GetRequirements() *bitset.BitSet {
	panic("implement me")
}

func (g *SystemGroup) SetRequirements(tags ...component.ComponentTag) {
	panic("implement me")
}
