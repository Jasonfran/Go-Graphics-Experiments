package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
)

type CameraSystem struct {
}

func NewCameraSystem() *CameraSystem {
	return &CameraSystem{}
}

func (c *CameraSystem) Init(context engine.EngineContext) {
	camera := context.EntityManager.NewEntity()
	camera.AddComponent(&components.Camera{FOV: 95.0})
}

func (c *CameraSystem) Update(context engine.EngineContext, delta float32) {
	cameras := context.EntityManager.GetEntitiesFromQuery(components.Includes(components.CameraComponentTag))
	renderables := context.EntityManager.GetEntitiesFromQuery(components.Includes(components.MeshComponentTag))

	cameras.Each(func(entity engine.IEntity) {
		camera, err := components.GetCameraComponent(entity)
		if err != nil {
			return
		}

		renderablesCopy := make([]engine.IEntity, 0, renderables.Len())
		for _, renderableList := range renderables {
			renderablesCopy = append(renderablesCopy, renderableList...)
		}
		camera.Renderables = renderablesCopy
	})
}

func (c *CameraSystem) LateUpdate(context engine.EngineContext, delta float32) {

}

func (c *CameraSystem) Shutdown(context engine.EngineContext) {

}
