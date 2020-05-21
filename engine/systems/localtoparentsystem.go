package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"

	"github.com/go-gl/mathgl/mgl32"
)

type LocalToParentSystem struct {
}

func NewLocalToParentSystem() *LocalToParentSystem {
	return &LocalToParentSystem{}
}

func (t *LocalToParentSystem) Init(context engine.EngineContext) {

}

func (t *LocalToParentSystem) Update(context engine.EngineContext, delta float32) {
	context.EntityManager.GetEntitiesFromQuery(
		components.Includes(components.TransformComponentTag, components.ChildrenComponentTag).
			Excludes(components.ParentComponentTag)).Each(func(entity engine.IEntity) {
		transformComponent, _ := components.GetTransformComponent(entity)
		childrenComponent, _ := components.GetChildrenComponent(entity)
		for _, child := range childrenComponent.GetChildren() {
			t.TransformChild(transformComponent.LocalToWorld, child)
		}
	})
}

func (t *LocalToParentSystem) TransformChild(parentLocalToWorld mgl32.Mat4, entity engine.IEntity) {
	transformComponent, _ := components.GetTransformComponent(entity)
	position := transformComponent.GetPos()
	scale := transformComponent.GetScale()
	rotation := transformComponent.GetRot()
	tMat4 := mgl32.Translate3D(position.X(), position.Y(), position.Z())
	sMat4 := mgl32.Scale3D(scale.X(), scale.Y(), scale.Z())
	rMat4 := rotation.Mat4()
	localToParentMatrix := tMat4.Mul4(rMat4).Mul4(sMat4)

	transformComponent.LocalToWorld = parentLocalToWorld.Mul4(localToParentMatrix)

	children, err := components.GetChildrenComponent(entity)
	if err != nil {
		return
	}

	for _, child := range children.GetChildren() {
		t.TransformChild(transformComponent.LocalToWorld, child)
	}
}

func (t *LocalToParentSystem) LateUpdate(context engine.EngineContext, delta float32) {

}

func (t *LocalToParentSystem) Shutdown(context engine.EngineContext) {

}
