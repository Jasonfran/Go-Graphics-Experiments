package resourcemanager

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
	"GraphicsStuff/engine/data"
	"log"
	"os"

	"github.com/go-gl/mathgl/mgl32"
)

type ResourceManager struct {
	loadedModels map[string]*data.Model
}

func New() *ResourceManager {
	return &ResourceManager{
		loadedModels: map[string]*data.Model{},
	}
}

func (rm *ResourceManager) LoadModel(path string) (*data.Model, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	log.Println(fileInfo.Name())

	return data.LoadGLTF(path)
}

func (rm *ResourceManager) ModelToEntity(context engine.EngineContext, model *data.Model) engine.IEntity {
	modelRoot := context.EntityManager.NewEntity()
	modelChildren := components.NewChildrenComponent()
	for _, node := range model.Nodes {
		modelChildren.AddChild(entityFromNode(context, node, modelRoot))
	}
	modelRoot.AddComponent(modelChildren)

	return modelRoot
}

func entityFromNode(context engine.EngineContext, thisNode *data.Node, parentEntity engine.IEntity) engine.IEntity {
	thisEntity := context.EntityManager.NewEntity()
	thisTransform, _ := components.GetTransformComponent(thisEntity)
	if thisNode.Matrix != mgl32.Ident4() {
		thisTransform.MatrixToTransform(thisNode.Matrix)
	} else {
		thisTransform.SetPosVec3(thisNode.Translation)
		thisTransform.SetRot(thisNode.Rotation)
		thisTransform.SetScaleVec3(thisNode.Scale)
	}

	if thisNode.Mesh != nil {
		thisEntity.AddComponent(&components.MeshComponent{Mesh: thisNode.Mesh})
	}

	if parentEntity != nil {
		thisEntity.AddComponent(components.NewParentComponent(parentEntity))
	}

	children := components.NewChildrenComponent()
	for _, child := range thisNode.Children {
		children.AddChild(entityFromNode(context, child, thisEntity))
	}
	thisEntity.AddComponent(children)

	return thisEntity
}
