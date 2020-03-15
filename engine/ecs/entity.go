package ecs

type entity struct {
	Id      EntityID
	manager Manager
}

func (e *entity) ID() EntityID {
	return e.Id
}

func (e *entity) AddComponent(component Component) {
	e.manager.AddComponent(e, component)
}

func (e *entity) GetComponent(tag ComponentTag) (Component, error) {
	return e.manager.GetComponent(e, tag)
}

func (e *entity) Destroy() {
	e.manager.DestroyEntity(e)
}
