package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/ecs"
	"GraphicsStuff/engine/ecs/components"
	"GraphicsStuff/engine/ecs/ecsmanager"
)

type CameraSystem struct {
	*ecsmanager.SystemEntityCollection
}

func NewCameraSystem() *CameraSystem {
	system := &CameraSystem{SystemEntityCollection: ecsmanager.NewSystemEntityCollection()}
	system.SetRequirements(components.CameraComponentTag)
	return system
}

func (c *CameraSystem) Init() {
	camera := engine.ECSManager.NewEntity()
	components.AddCameraComponent(camera, &components.Camera{FOV: 95.0})
}

func (c *CameraSystem) Update(delta float32) {
	renderables := engine.ECSManager.GetEntitiesWithComponents(components.MeshRendererComponentTag)

	for _, entity := range c.Entities() {
		camera, err := components.GetCameraComponent(entity)
		if err != nil {
			continue
		}

		renderablesCopy := make([]ecs.Entity, len(renderables))
		copy(renderablesCopy, renderables)
		camera.Renderables = renderablesCopy
	}
}

func (c *CameraSystem) LateUpdate(delta float32) {

}

func (c *CameraSystem) Shutdown() {

}
