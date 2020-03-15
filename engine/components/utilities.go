package components

import "GraphicsStuff/engine/ecs"

func GetTransformComponent(e ecs.Entity) (*Transform, error) {
	c, err := e.GetComponent(TransformComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Transform), nil
}

func GetCameraComponent(e ecs.Entity) (*Camera, error) {
	c, err := e.GetComponent(CameraComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Camera), nil
}

func GetMeshRendererComponent(e ecs.Entity) (*MeshRenderer, error) {
	c, err := e.GetComponent(MeshRendererComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*MeshRenderer), nil
}

func GetMaterialComponent(e ecs.Entity) (*Material, error) {
	c, err := e.GetComponent(MaterialComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*Material), nil
}

func GetPlayerComponent(e ecs.Entity) (*PlayerComponent, error) {
	c, err := e.GetComponent(PlayerComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*PlayerComponent), nil
}

func GetPhysicsComponent(e ecs.Entity) (*PhysicsComponent, error) {
	c, err := e.GetComponent(PhysicsComponentTag)
	if err != nil {
		return nil, err
	}

	return c.(*PhysicsComponent), nil
}
