package cache

import (
	"GraphicsStuff/engine/component"
	"log"
	"sort"

	"github.com/willf/bitset"
)

var Instance = NewBitsetCache()

type TagsHash uint64

func HashTags(tags ...component.ComponentTag) TagsHash {
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
	tags  map[*bitset.BitSet][]component.ComponentTag
}

func NewBitsetCache() *BitsetCache {
	return &BitsetCache{
		empty: bitset.New(8),
		cache: map[TagsHash]*bitset.BitSet{},
		tags:  map[*bitset.BitSet][]component.ComponentTag{},
	}
}

func (c *BitsetCache) New(tags ...component.ComponentTag) *bitset.BitSet {
	if len(tags) <= 0 {
		return c.empty
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i] < tags[j]
	})

	hashKey := HashTags(tags...)
	cached, ok := c.cache[hashKey]
	if !ok {
		cached = bitset.New(uint(len(tags)))
		for _, tag := range tags {
			cached.Set(uint(tag))
		}

		c.cache[hashKey] = cached
		copyTags := make([]component.ComponentTag, len(tags))
		copy(copyTags, tags)
		c.tags[cached] = copyTags
		log.Printf("Created new bitset for %v", tags)
	}

	return cached
}

func (c *BitsetCache) Append(b *bitset.BitSet, tag component.ComponentTag) *bitset.BitSet {
	tags := c.GetTags(b)
	tags = append(tags, tag)
	return c.New(tags...)
}

func (c *BitsetCache) GetTags(b *bitset.BitSet) []component.ComponentTag {
	return c.tags[b]
}
