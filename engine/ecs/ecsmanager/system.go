package ecsmanager

import (
	"GraphicsStuff/engine/ecs"

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
	AddEntity(e ecs.Entity)
	RemoveEntity(id ecs.EntityId)
	GetRequirements() *bitset.BitSet
	SetRequirements(tags ...ecs.ComponentTag)
}

type SystemEntityCollection struct {
	bitsetCache   *BitsetCache
	entitiesList  []ecs.Entity
	entitiesIndex map[ecs.EntityId]int
	requirements  *bitset.BitSet
}

func NewSystemEntityCollection() *SystemEntityCollection {
	return &SystemEntityCollection{
		bitsetCache:   BitsetCacheInstance,
		entitiesList:  []ecs.Entity{},
		entitiesIndex: map[ecs.EntityId]int{},
		requirements:  BitsetCacheInstance.New(),
	}
}

func (c *SystemEntityCollection) AddEntity(e ecs.Entity) {
	c.entitiesList = append(c.entitiesList, e)
	c.entitiesIndex[e.Id()] = len(c.entitiesList) - 1
}

func (c *SystemEntityCollection) RemoveEntity(id ecs.EntityId) {
	index := c.entitiesIndex[id]
	lastEntity := c.entitiesList[len(c.entitiesList)-1]
	c.entitiesList[index] = lastEntity
	c.entitiesIndex[lastEntity.Id()] = index
	c.entitiesList = c.entitiesList[:len(c.entitiesList)-1]
	delete(c.entitiesIndex, id)
}

func (c *SystemEntityCollection) GetRequirements() *bitset.BitSet {
	return c.requirements
}

func (c *SystemEntityCollection) SetRequirements(tags ...ecs.ComponentTag) {
	c.requirements = c.bitsetCache.New(tags...)
}

func (c *SystemEntityCollection) Entities() []ecs.Entity {
	return c.entitiesList
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

func (g *SystemGroup) AddEntity(e ecs.Entity) {
	for _, system := range g.systems {
		req := system.GetRequirements()
		if e.ComponentBitset().IsSuperSet(req) {
			system.AddEntity(e)
		}
	}
}

func (g *SystemGroup) RemoveEntity(id ecs.EntityId) {
	panic("implement me")
}

func (g *SystemGroup) GetRequirements() *bitset.BitSet {
	panic("implement me")
}

func (g *SystemGroup) SetRequirements(tags ...ecs.ComponentTag) {
	panic("implement me")
}
