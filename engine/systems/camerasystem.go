package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/component"
	"GraphicsStuff/engine/ecsmanager"
)

type CameraSystem struct {
	*ecsmanager.SystemEntityCollection
}

func NewCameraSystem() *CameraSystem {
	return &CameraSystem{SystemEntityCollection: ecsmanager.NewSystemEntityCollection()}
}

func (c *CameraSystem) Init() {
	camera := engine.ECSManager.NewEntity()
	camera.AddCameraComponent(&component.Camera{FOV: 95.0})
}

func (c *CameraSystem) Update(delta float32) {
	renderables := engine.ECSManager.GetEntitiesWithComponents(component.MeshRendererComponentTag)

	for _, _ = range c.Entities() {
		for _, _ = range renderables {
		}
	}
}

func (c *CameraSystem) LateUpdate(delta float32) {

}

func (c *CameraSystem) Shutdown() {

}
