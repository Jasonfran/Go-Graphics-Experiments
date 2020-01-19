package engine

import "github.com/go-gl/glfw/v3.3/glfw"

type keyState [2]bool

type Manager struct {
	mouseX    int
	mouseY    int
	keyStates map[glfw.Key]keyState
}

func NewManager() *Manager {
	return &Manager{
		keyStates: map[glfw.Key]keyState{},
	}
}

func (m *Manager) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	state, ok := m.keyStates[key]
	if !ok {
		state = keyState{false, false}
	}

	if action == glfw.Press {
		state[0] = true
		state[1] = true
	}

	if action == glfw.Release {
		state[1] = false
	}

	m.keyStates[key] = state
}

func (m *Manager) Update() {
	for key, state := range m.keyStates {
		state[0] = false
		m.keyStates[key] = state
	}
}

func (m *Manager) Pressed(key glfw.Key) bool {
	return m.keyStates[key][0]
}

func (m *Manager) Held(key glfw.Key) bool {
	return m.keyStates[key][1]
}
