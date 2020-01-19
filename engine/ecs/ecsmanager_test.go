package ecs_test

import (
	"GraphicsStuff/engine/ecs"
	"GraphicsStuff/playersystems"
	"testing"
)

func BenchmarkECSManager_AddComponent(b *testing.B) {
	em := ecs.NewECSManager()
	entity := em.NewEntity()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entity.AddCameraComponent(&ecs.Camera{})
	}
}

func BenchmarkECSManager_GetComponent(b *testing.B) {
	em := ecs.NewECSManager()
	entity := em.NewEntity()
	entity.AddCameraComponent(&ecs.Camera{FOV: 95.0})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := entity.GetCameraComponent()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSystemEntityCollection_AddEntity(b *testing.B) {
	em := ecs.NewECSManager()
	system := playersystems.NewTestPlayerSystem()
	system.SetRequirements(ecs.CameraComponentTag)
	em.AddSystem(ecs.PlayerSystemGroup, system)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entity := em.NewEntity()
		em.AddComponent(entity, ecs.CameraComponentTag, &ecs.Camera{})
	}
}

func BenchmarkECSManager_GetEntitiesWithComponents(b *testing.B) {
	em := ecs.NewECSManager()
	for i := 0; i < 1000; i++ {
		e := em.NewEntity()
		if i%2 == 0 {
			e.AddMeshRendererComponent(&ecs.MeshRenderer{})
			e.AddMaterialComponent(&ecs.Material{})
		}
		e.AddCameraComponent(&ecs.Camera{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cameraentities := em.GetEntitiesWithComponents(ecs.CameraComponentTag)
		meshentities := em.GetEntitiesWithComponents(ecs.MeshRendererComponentTag, ecs.MaterialComponentTag)
		if len(cameraentities) < 1000 {
			b.FailNow()
		}

		if len(meshentities) < 500 {
			b.FailNow()
		}
	}
}
