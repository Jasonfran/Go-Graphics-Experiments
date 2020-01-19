package ecs

type Camera struct {
	FOV         float32
	Renderables []*Entity
}

func (c *Camera) Name() string {
	return "Camera"
}
