package ecs

import (
	"github.com/willf/bitset"
)

type EntityId uint
type ComponentTag uint64

type Component interface {
	Name() string
}

type Entity interface {
	Id() EntityId
	ComponentBitset() *bitset.BitSet
	GetComponent(tag ComponentTag) (Component, error)
	AddComponent(tag ComponentTag, c Component)
	Transform() *Transform
}
