package components

import (
	"GraphicsStuff/engine"

	"github.com/willf/bitset"
)

type Query struct {
	inclusionSet *bitset.BitSet
	exclusionSet *bitset.BitSet
}

func Includes(tags ...engine.ComponentTag) Query {
	bs := bitset.New(uint(len(tags)))
	for _, tag := range tags {
		bs.Set(uint(tag))
	}

	return Query{
		inclusionSet: bs,
		exclusionSet: bitset.New(0),
	}
}

func (q Query) Excludes(tags ...engine.ComponentTag) Query {
	bs := bitset.New(uint(len(tags)))
	for _, tag := range tags {
		bs.Set(uint(tag))
	}

	q.exclusionSet = bs
	return q
}

func (q Query) InclusionSet() *bitset.BitSet {
	return q.inclusionSet
}

func (q Query) ExclusionSet() *bitset.BitSet {
	return q.exclusionSet
}
