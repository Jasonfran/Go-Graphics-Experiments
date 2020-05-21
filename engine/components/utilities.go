package components

import (
	"GraphicsStuff/engine"
)

func GetTransformComponent(e engine.IEntity) (*Transform, error) {
	c, err := e.GetComponent(TransformComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Transform), nil
}

func GetCameraComponent(e engine.IEntity) (*Camera, error) {
	c, err := e.GetComponent(CameraComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Camera), nil
}

func GetModelComponent(e engine.IEntity) (*ModelComponent, error) {
	c, err := e.GetComponent(ModelComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*ModelComponent), nil
}

func GetPlayerComponent(e engine.IEntity) (*PlayerComponent, error) {
	c, err := e.GetComponent(PlayerComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*PlayerComponent), nil
}

func GetPhysicsComponent(e engine.IEntity) (*PhysicsComponent, error) {
	c, err := e.GetComponent(PhysicsComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*PhysicsComponent), nil
}

func GetParentComponent(e engine.IEntity) (*ParentComponent, error) {
	c, err := e.GetComponent(ParentComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*ParentComponent), nil
}

func GetChildrenComponent(e engine.IEntity) (*ChildrenComponent, error) {
	c, err := e.GetComponent(ChildrenComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*ChildrenComponent), nil
}

func GetMeshComponent(e engine.IEntity) (*MeshComponent, error) {
	c, err := e.GetComponent(MeshComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*MeshComponent), nil
}
