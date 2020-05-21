package ecs

import (
	"GraphicsStuff/engine"
	"testing"
)

type TestComponent struct {
	CustomTag engine.ComponentTag
}

func (t *TestComponent) String() string {
	return "TestComponent"
}

func (t *TestComponent) Tag() engine.ComponentTag {
	return t.CustomTag
}

func BenchmarkEntityComponentCollection_Add(b *testing.B) {
	collection := NewEntityComponentCollection()
	entity := &Entity{Id: 1}
	for i := 0; i < b.N; i++ {
		collection.Add(entity, &TestComponent{1})
	}
}

func BenchmarkEntityComponentCollection_Get(b *testing.B) {
	collection := NewEntityComponentCollection()
	entity := &Entity{Id: 1}
	collection.Add(entity, &TestComponent{1})
	for i := 0; i < b.N; i++ {
		collection.Get(entity, engine.ComponentTag(1))
	}
}

func BenchmarkEntityComponentCollection_GetEntitiesWithComponents(b *testing.B) {
	collection := NewEntityComponentCollection()
	for i := 0; i < 1000; i++ {
		entity := &Entity{Id: engine.EntityID(i)}
		if i%2 == 0 {
			collection.Add(entity, &TestComponent{TestTag1})
			collection.Add(entity, &TestComponent{TestTag2})
		}
		collection.Add(entity, &TestComponent{TestTag3})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collection.GetEntitiesFromQuery(TestTag3)
		collection.GetEntitiesFromQuery(TestTag1, TestTag2)
	}
}

func TestEntityComponentCollection_GetEntitiesWithComponents(t *testing.T) {
	collection := NewEntityComponentCollection()
	for i := 0; i < 4; i++ {
		entity := &Entity{Id: engine.EntityID(i)}
		if i%2 == 0 {
			collection.Add(entity, &TestComponent{TestTag1})
			collection.Add(entity, &TestComponent{TestTag2})
		}
		collection.Add(entity, &TestComponent{TestTag3})
	}

	if collection.GetEntitiesFromQuery(TestTag3).Len() != 4 {
		t.Fail()
	}

	if collection.GetEntitiesFromQuery(TestTag1, TestTag2).Len() != 2 {
		t.Fail()
	}
}
