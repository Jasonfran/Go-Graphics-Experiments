package ecsmanager

import (
	"GraphicsStuff/engine/ecs"
	"log"
	"sort"

	"github.com/willf/bitset"
)

var BitsetCacheInstance = NewBitsetCache()

type TagsHash uint64

func HashTags(tags ...ecs.ComponentTag) TagsHash {
	hash := uint64(0)
	for _, tag := range tags {
		hash += uint64(tag)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return TagsHash(hash)
}

type BitsetCache struct {
	empty *bitset.BitSet
	cache map[TagsHash]*bitset.BitSet
	tags  map[*bitset.BitSet][]ecs.ComponentTag
}

func NewBitsetCache() *BitsetCache {
	return &BitsetCache{
		empty: bitset.New(8),
		cache: map[TagsHash]*bitset.BitSet{},
		tags:  map[*bitset.BitSet][]ecs.ComponentTag{},
	}
}

func (c *BitsetCache) New(t ...ecs.ComponentTag) *bitset.BitSet {
	if len(t) <= 0 {
		return c.empty
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	hashKey := HashTags(t...)
	cached, ok := c.cache[hashKey]
	if !ok {
		cached = bitset.New(uint(len(t)))
		for _, tag := range t {
			cached.Set(uint(tag))
		}

		c.cache[hashKey] = cached
		copyTags := make([]ecs.ComponentTag, len(t))
		copy(copyTags, t)
		c.tags[cached] = copyTags
		log.Printf("Created new bitset for %v", t)
	}

	return cached
}

func (c *BitsetCache) Append(b *bitset.BitSet, tag ecs.ComponentTag) *bitset.BitSet {
	tags := c.GetTags(b)
	tags = append(tags, tag)
	return c.New(tags...)
}

func (c *BitsetCache) GetTags(b *bitset.BitSet) []ecs.ComponentTag {
	return c.tags[b]
}
