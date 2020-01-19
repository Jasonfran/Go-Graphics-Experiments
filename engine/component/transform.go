package component

import "github.com/go-gl/mathgl/mgl32"

type Transform struct {
	Pos          mgl32.Vec3
	LocalToWorld mgl32.Mat4
}

func (t *Transform) Name() string {
	return "Transform"
}

func (t *Transform) Translate(x, y, z float32) {
	t.Pos = t.Pos.Add(mgl32.Vec3{x, y, z})
	t.calculateMatrix()
}

func (t *Transform) calculateMatrix() {
	translate := mgl32.Translate3D(t.Pos.X(), t.Pos.Y(), t.Pos.Z())
	t.LocalToWorld = translate
}
