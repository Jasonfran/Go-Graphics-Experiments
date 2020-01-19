package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/ecs"
)

type CameraSystem struct {
	*ecs.SystemEntityCollection
}

func NewCameraSystem() *CameraSystem {
	system := &CameraSystem{SystemEntityCollection: ecs.NewSystemEntityCollection()}
	system.SetRequirements(ecs.CameraComponentTag)
	return system
}

func (c *CameraSystem) Init() {
	camera := engine.ECSManager.NewEntity()
	camera.AddCameraComponent(&ecs.Camera{FOV: 95.0})
}

func (c *CameraSystem) Update(delta float32) {
	renderables := engine.ECSManager.GetEntitiesWithComponents(ecs.MeshRendererComponentTag)

	for _, entity := range c.Entities() {
		camera, err := entity.GetCameraComponent()
		if err != nil {
			continue
		}

		renderablesCopy := make([]*ecs.Entity, len(renderables))
		copy(renderablesCopy, renderables)
		camera.Renderables = renderablesCopy
	}
}

func (c *CameraSystem) LateUpdate(delta float32) {

}

func (c *CameraSystem) Shutdown() {

}
