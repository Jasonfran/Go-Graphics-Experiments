package ecs

import (
	"GraphicsStuff/engine"

	"github.com/willf/bitset"
)

type archetypeHash uint64

var archetypeCache = map[archetypeHash]Archetype{
	hashTags(): {bs: bitset.New(0)},
}

// Archetype contains a description of components and entity contains.
// This shouldn't be created manually if you wish to keep map key equality
type Archetype struct {
	bs *bitset.BitSet
}

func NewArchetype(tags ...engine.ComponentTag) Archetype {
	hash := hashTags(tags...)
	archetype, ok := archetypeCache[hash]
	if ok {
		return archetype
	}

	bs := bitset.New(uint(len(tags)))
	for _, tag := range tags {
		bs.Set(uint(tag))
	}

	archetype = Archetype{
		bs: bs,
	}

	archetypeCache[hash] = archetype
	return archetype
}

func (a Archetype) Satisfies(query engine.IQuery) bool {
	return a.bs.IsSuperSet(query.InclusionSet()) && a.bs.IntersectionCardinality(query.ExclusionSet()) == 0
}

func (a Archetype) AddType(tag engine.ComponentTag) Archetype {
	return NewArchetype(append(a.GetTypes(), tag)...)
}

func (a Archetype) RemoveType(tag engine.ComponentTag) Archetype {
	tags := a.GetTypes()
	for i, componentTag := range tags {
		if componentTag == tag {
			tags[i] = tags[len(tags)-1]
			tags = tags[:len(tags)-1]
			return NewArchetype(tags...)
		}
	}
	return a
}

func (a Archetype) GetTypes() []engine.ComponentTag {
	tags := make([]engine.ComponentTag, 0, a.bs.Len())
	for i := uint(0); i < a.bs.Len(); i++ {
		if a.bs.Test(i) {
			tags = append(tags, engine.ComponentTag(i))
		}
	}

	return tags
}

func hashTags(tags ...engine.ComponentTag) archetypeHash {
	hash := uint64(0)
	for _, tag := range tags {
		hash += uint64(tag)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return archetypeHash(hash)
}
