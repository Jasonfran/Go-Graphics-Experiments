package engine

import "github.com/go-gl/mathgl/mgl32"

type Vertex struct {
	Pos       mgl32.Vec3
	Normal    mgl32.Vec3
	Color     mgl32.Vec3
	TexCoords mgl32.Vec2
}
