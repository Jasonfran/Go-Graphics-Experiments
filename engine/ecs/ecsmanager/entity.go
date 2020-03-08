package ecsmanager

import (
	"GraphicsStuff/engine/ecs"

	"github.com/willf/bitset"
)

type entity struct {
	ecsManager *ECSManager
	id         ecs.EntityId
	transform  *ecs.Transform
}

func (e *entity) Id() ecs.EntityId {
	return e.id
}

func (e *entity) ComponentBitset() *bitset.BitSet {
	return e.ecsManager.GetEntityBitset(e)
}

func (e *entity) GetComponent(tag ecs.ComponentTag) (ecs.Component, error) {
	return e.ecsManager.GetComponent(e, tag)
}

func (e *entity) AddComponent(tag ecs.ComponentTag, c ecs.Component) {
	e.ecsManager.AddComponent(e, tag, c)
}

func (e *entity) Transform() *ecs.Transform {
	return e.transform
}
