package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
	"GraphicsStuff/engine/ecs"
)

type CameraSystem struct {
}

func NewCameraSystem() *CameraSystem {
	return &CameraSystem{}
}

func (c *CameraSystem) Init() {
	camera := engine.ECSManager.NewEntity()
	camera.AddComponent(&components.Camera{FOV: 95.0})
}

func (c *CameraSystem) Update(delta float32) {
	cameras := engine.ECSManager.GetEntitiesWithComponents(components.CameraComponentTag)
	renderables := engine.ECSManager.GetEntitiesWithComponents(components.MeshRendererComponentTag)

	cameras.Each(func(entity ecs.Entity) {
		camera, err := components.GetCameraComponent(entity)
		if err != nil {
			return
		}

		renderablesCopy := make([]ecs.Entity, 0, renderables.Len())
		for _, renderableList := range renderables {
			renderablesCopy = append(renderablesCopy, renderableList...)
		}
		camera.Renderables = renderablesCopy
	})
}

func (c *CameraSystem) LateUpdate(delta float32) {

}

func (c *CameraSystem) Shutdown() {

}
