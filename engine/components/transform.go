package components

import (
	"GraphicsStuff/engine/ecs"

	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	Pos          mgl32.Vec3
	LocalToWorld mgl32.Mat4
}

func (t *Transform) String() string {
	return "Transform"
}

func (t *Transform) Tag() ecs.ComponentTag {
	return TransformComponentTag
}

func (t *Transform) Translate(x, y, z float32) {
	t.Pos = t.Pos.Add(mgl32.Vec3{x, y, z})
	t.calculateMatrix()
}

func (t *Transform) calculateMatrix() {
	translate := mgl32.Translate3D(t.Pos.X(), t.Pos.Y(), t.Pos.Z())
	//scale := mgl32.Scale3D(0.05, 0.05, 0.05)
	t.LocalToWorld = translate //.Mul4(scale)
}
