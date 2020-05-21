package engine

type EngineContext struct {
	EntityManager   IEntityManager
	EventDispatcher *EventDispatcher // Replace with interface
	InputManager    IInputManager
	ResourceManager IResourceManager // Replace with interface
}
