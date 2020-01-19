package ecsmanager_test

import (
	"GraphicsStuff/engine/component"
	"GraphicsStuff/engine/ecsmanager"
	"GraphicsStuff/playersystems"
	"testing"
)

func BenchmarkECSManager_AddComponent(b *testing.B) {
	em := ecsmanager.NewECSManager()
	entity := em.NewEntity()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entity.AddCameraComponent(&component.Camera{})
	}
}

func BenchmarkECSManager_GetComponent(b *testing.B) {
	em := ecsmanager.NewECSManager()
	entity := em.NewEntity()
	entity.AddCameraComponent(&component.Camera{FOV: 95.0})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := entity.GetCameraComponent()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSystemEntityCollection_AddEntity(b *testing.B) {
	em := ecsmanager.NewECSManager()
	system := playersystems.NewTestPlayerSystem()
	system.SetRequirements(component.CameraComponentTag)
	em.AddSystem(ecsmanager.PlayerSystemGroup, system)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entity := em.NewEntity()
		em.AddComponent(entity, component.CameraComponentTag, &component.Camera{})
	}
}

func BenchmarkECSManager_GetEntitiesWithComponents(b *testing.B) {
	em := ecsmanager.NewECSManager()
	for i := 0; i < 1000; i++ {
		e := em.NewEntity()
		if i%2 == 0 {
			e.AddMeshRendererComponent(&component.MeshRenderer{})
			e.AddMaterialComponent(&component.Material{})
		}
		e.AddCameraComponent(&component.Camera{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cameraentities := em.GetEntitiesWithComponents(component.CameraComponentTag)
		meshentities := em.GetEntitiesWithComponents(component.MeshRendererComponentTag, component.MaterialComponentTag)
		if len(cameraentities) < 1000 {
			b.FailNow()
		}

		if len(meshentities) < 500 {
			b.FailNow()
		}
	}
}
