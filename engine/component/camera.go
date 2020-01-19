package component

type Camera struct {
	FOV float32
}

func (c *Camera) Name() string {
	return "Camera"
}
