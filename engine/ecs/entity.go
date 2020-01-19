package ecs

import (
	"github.com/willf/bitset"
)

type EntityId uint

type Entity struct {
	ecsManager *ECSManager
	id         EntityId
	transform  *Transform
}

func (e *Entity) Id() EntityId {
	return e.id
}

func (e *Entity) ComponentBitset() *bitset.BitSet {
	return e.ecsManager.GetEntityBitset(e)
}

func (e *Entity) Transform() *Transform {
	return e.transform
}

func (e *Entity) AddCameraComponent(t *Camera) {
	e.ecsManager.AddComponent(e, CameraComponentTag, t)
}

func (e *Entity) GetCameraComponent() (*Camera, error) {
	c, err := e.ecsManager.GetComponent(e, CameraComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Camera), nil
}

func (e *Entity) AddMeshRendererComponent(t *MeshRenderer) {
	e.ecsManager.AddComponent(e, MeshRendererComponentTag, t)
}

func (e *Entity) GetMeshRendererComponent() (*MeshRenderer, error) {
	c, err := e.ecsManager.GetComponent(e, MeshRendererComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*MeshRenderer), nil
}

func (e *Entity) AddMaterialComponent(t *Material) {
	e.ecsManager.AddComponent(e, MaterialComponentTag, t)
}

func (e *Entity) GetMaterialComponent() (*Material, error) {
	c, err := e.ecsManager.GetComponent(e, MaterialComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Material), nil
}
