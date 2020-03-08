package components

import "GraphicsStuff/engine/ecs"

func AddCameraComponent(e ecs.Entity, t *Camera) {
	e.AddComponent(CameraComponentTag, t)
}

func GetCameraComponent(e ecs.Entity) (*Camera, error) {
	c, err := e.GetComponent(CameraComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Camera), nil
}

func AddMeshRendererComponent(e ecs.Entity, t *MeshRenderer) {
	e.AddComponent(MeshRendererComponentTag, t)
}

func GetMeshRendererComponent(e ecs.Entity) (*MeshRenderer, error) {
	c, err := e.GetComponent(MeshRendererComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*MeshRenderer), nil
}

func AddMaterialComponent(e ecs.Entity, t *Material) {
	e.AddComponent(MaterialComponentTag, t)
}

func GetMaterialComponent(e ecs.Entity) (*Material, error) {
	c, err := e.GetComponent(MaterialComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Material), nil
}

func AddPlayerComponent(e ecs.Entity, t *PlayerComponent) {
	e.AddComponent(PlayerComponentTag, t)
}

func GetPlayerComponent(e ecs.Entity) (*PlayerComponent, error) {
	c, err := e.GetComponent(PlayerComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*PlayerComponent), nil
}

func AddPhysicsComponent(e ecs.Entity, t *PhysicsComponent) {
	e.AddComponent(PhysicsComponentTag, t)
}

func GetPhysicsComponent(e ecs.Entity) (*PhysicsComponent, error) {
	c, err := e.GetComponent(PhysicsComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*PhysicsComponent), nil
}
