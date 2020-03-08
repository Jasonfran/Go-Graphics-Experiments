package ecsmanager_test

import (
	"GraphicsStuff/engine/ecs/components"
	"GraphicsStuff/engine/ecs/ecsmanager"
	"GraphicsStuff/playersystems"
	"testing"
)

func BenchmarkECSManager_AddComponent(b *testing.B) {
	em := ecsmanager.NewECSManager()
	entity := em.NewEntity()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		components.AddCameraComponent(entity, &components.Camera{})
	}
}

func BenchmarkECSManager_GetComponent(b *testing.B) {
	em := ecsmanager.NewECSManager()
	entity := em.NewEntity()
	components.AddCameraComponent(entity, &components.Camera{FOV: 95.0})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := components.GetCameraComponent(entity)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSystemEntityCollection_AddEntity(b *testing.B) {
	em := ecsmanager.NewECSManager()
	system := playersystems.NewTestPlayerSystem()
	system.SetRequirements(components.CameraComponentTag)
	em.AddSystem(ecsmanager.PlayerSystemGroup, system)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entity := em.NewEntity()
		components.AddCameraComponent(entity, &components.Camera{FOV: 95.0})
	}
}

func BenchmarkECSManager_GetEntitiesWithComponents(b *testing.B) {
	em := ecsmanager.NewECSManager()
	for i := 0; i < 1000; i++ {
		e := em.NewEntity()
		if i%2 == 0 {
			components.AddMeshRendererComponent(e, &components.MeshRenderer{})
			components.AddMaterialComponent(e, &components.Material{})
		}
		components.AddCameraComponent(e, &components.Camera{})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cameraentities := em.GetEntitiesWithComponents(components.CameraComponentTag)
		meshentities := em.GetEntitiesWithComponents(components.MeshRendererComponentTag, components.MaterialComponentTag)
		if len(cameraentities) < 1000 {
			b.FailNow()
		}

		if len(meshentities) < 500 {
			b.FailNow()
		}
	}
}
