package ecs

import (
	"testing"
)

const (
	TestTag1 ComponentTag = iota
	TestTag2
	TestTag3
	TestTag4
)

func TestArchetype_Satisfies(t *testing.T) {
	query := NewQuery(TestTag1, TestTag2)
	archetype := NewArchetype(TestTag1, TestTag2, TestTag3, TestTag4)
	if !archetype.Satisfies(query) {
		t.Fail()
	}
}

func TestArchetype_Equality(t *testing.T) {
	archetype1 := NewArchetype(1, 2, 3)
	archetype2 := NewArchetype(1, 2, 3)
	if archetype1 != archetype2 {
		t.Fail()
	}

	archetype3 := NewArchetype(1, 2).AddType(3)
	if archetype1 != archetype3 {
		t.Fail()
	}
}

func BenchmarkArchetype_Satisfies(b *testing.B) {
	query := NewQuery(TestTag1, TestTag2)
	archetype := NewArchetype(TestTag1, TestTag2, TestTag3, TestTag4)
	for i := 0; i < b.N; i++ {
		archetype.Satisfies(query)
	}
}

func BenchmarkArchetype_GetTypes(b *testing.B) {
	archetype := NewArchetype(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for i := 0; i < b.N; i++ {
		archetype.GetTypes()
	}
}
