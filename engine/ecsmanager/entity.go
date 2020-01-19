package ecsmanager

import (
	"GraphicsStuff/engine/component"

	"github.com/willf/bitset"
)

type EntityId uint

type Entity struct {
	ecsManager *ECSManager
	Id         EntityId
	transform  *component.Transform
}

func (e *Entity) ComponentBitset() *bitset.BitSet {
	return e.ecsManager.GetEntityBitset(e)
}

func (e *Entity) Transform() *component.Transform {
	return e.transform
}

func (e *Entity) AddCameraComponent(t *component.Camera) {
	e.ecsManager.AddComponent(e, component.CameraComponentTag, t)
}

func (e *Entity) GetCameraComponent() (*component.Camera, error) {
	c, err := e.ecsManager.GetComponent(e, component.CameraComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*component.Camera), nil
}

func (e *Entity) AddMeshRendererComponent(t *component.MeshRenderer) {
	e.ecsManager.AddComponent(e, component.MeshRendererComponentTag, t)
}

func (e *Entity) GetMeshRendererComponent() (*component.MeshRenderer, error) {
	c, err := e.ecsManager.GetComponent(e, component.MeshRendererComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*component.MeshRenderer), nil
}

func (e *Entity) AddMaterialComponent(t *component.Material) {
	e.ecsManager.AddComponent(e, component.MaterialComponentTag, t)
}

func (e *Entity) GetMaterialComponent() (*component.Material, error) {
	c, err := e.ecsManager.GetComponent(e, component.MaterialComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*component.Material), nil
}
