package components

import (
	"GraphicsStuff/engine"

	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	pos      mgl32.Vec3
	scale    mgl32.Vec3
	rotation mgl32.Quat

	// Defined as the absolute world position
	LocalToWorld mgl32.Mat4

	// If the entity is parented them this will represent the position relative to the parent
	LocalToParent mgl32.Mat4
}

func (t *Transform) String() string {
	return "Transform"
}

func (t *Transform) Tag() engine.ComponentTag {
	return TransformComponentTag
}

func (t *Transform) Translate(x, y, z float32) {
	t.pos = t.pos.Add(mgl32.Vec3{x, y, z})
}

func (t *Transform) GetPos() mgl32.Vec3 {
	return t.pos
}

func (t *Transform) SetPos(x, y, z float32) {
	t.pos = mgl32.Vec3{x, y, z}
}

func (t *Transform) SetPosVec3(v mgl32.Vec3) {
	t.pos = v
}

func (t *Transform) SetRot(q mgl32.Quat) {
	t.rotation = q
}

func (t *Transform) GetRot() mgl32.Quat {
	return t.rotation
}

func (t *Transform) SetScale(x, y, z float32) {
	t.scale = mgl32.Vec3{x, y, z}
}

func (t *Transform) SetScaleVec3(v mgl32.Vec3) {
	t.scale = v
}

func (t *Transform) GetScale() mgl32.Vec3 {
	return t.scale
}

func (t *Transform) MatrixToTransform(m mgl32.Mat4) {
	t.SetPos(m[12], m[13], m[14])
	sX := mgl32.Vec3{m[0], m[1], m[2]}.Len()
	sY := mgl32.Vec3{m[4], m[5], m[6]}.Len()
	sZ := mgl32.Vec3{m[8], m[9], m[10]}.Len()
	scale := mgl32.Vec3{sX, sY, sZ}
	t.SetScaleVec3(scale)

	invSX := 1 / sX
	invSY := 1 / sY
	invSZ := 1 / sZ

	rotationMatrix := mgl32.Mat4{
		m[0] * invSX, m[1] * invSX, m[2] * invSX, 0,
		m[4] * invSY, m[5] * invSY, m[6] * invSY, 0,
		m[8] * invSZ, m[9] * invSZ, m[10] * invSZ, 0,
		0, 0, 0, 1,
	}

	rotation := mgl32.Mat4ToQuat(rotationMatrix)
	t.SetRot(rotation)
}

func IdentityTransform() *Transform {
	return &Transform{
		pos:           mgl32.Vec3{},
		scale:         mgl32.Vec3{1, 1, 1},
		rotation:      mgl32.QuatIdent(),
		LocalToWorld:  mgl32.Ident4(),
		LocalToParent: mgl32.Ident4(),
	}
}
