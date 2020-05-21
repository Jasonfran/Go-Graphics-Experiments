package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"

	"github.com/go-gl/mathgl/mgl32"
)

type LocalToWorldSystem struct {
}

func NewLocalToWorldSystem() *LocalToWorldSystem {
	return &LocalToWorldSystem{}
}

func (t *LocalToWorldSystem) Init(context engine.EngineContext) {

}

func (t *LocalToWorldSystem) Update(context engine.EngineContext, delta float32) {
	context.EntityManager.GetEntitiesFromQuery(components.Includes(components.TransformComponentTag)).Each(func(entity engine.IEntity) {
		// If it's parented then we don't care, that is handled by another system
		//if entity.HasComponent(components.ParentComponentTag) {
		//	return
		//}

		transformComponent, _ := components.GetTransformComponent(entity)
		position := transformComponent.GetPos()
		scale := transformComponent.GetScale()
		rotation := transformComponent.GetRot()
		tMat4 := mgl32.Translate3D(position.X(), position.Y(), position.Z())
		sMat4 := mgl32.Scale3D(scale.X(), scale.Y(), scale.Z())
		rMat4 := rotation.Mat4()
		localToWorldMatrix := tMat4.Mul4(rMat4).Mul4(sMat4)

		transformComponent.LocalToWorld = localToWorldMatrix
	})
}

func (t *LocalToWorldSystem) LateUpdate(context engine.EngineContext, delta float32) {

}

func (t *LocalToWorldSystem) Shutdown(context engine.EngineContext) {

}
