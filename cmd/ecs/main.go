package main

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
	"GraphicsStuff/engine/ecs"
	"GraphicsStuff/engine/events"
	"GraphicsStuff/engine/input"
	"GraphicsStuff/engine/resourcemanager"
	"GraphicsStuff/engine/systems"
	"GraphicsStuff/playersystems"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

type Game struct {
	Window  *glfw.Window
	Width   int
	Height  int
	Context engine.EngineContext
}

func (g *Game) Init(context engine.EngineContext, width int, height int) {
	g.Context = context

	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(width, height, "Testing", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()
	window.SetSizeCallback(g.windowSizeCallback)
	window.SetKeyCallback(g.Context.InputManager.KeyCallback)
	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	g.Width = width
	g.Height = height
	g.Window = window

}

func (g *Game) Terminate() {
	glfw.Terminate()
}

func (g *Game) windowSizeCallback(w *glfw.Window, width int, height int) {
	g.Width = width
	g.Height = height

	g.Context.EventDispatcher.Trigger(events.WindowResizedEvent, events.WindowResizedEventData{
		Width:  width,
		Height: height,
	})
}

func (g *Game) Run() {
	g.Context.EntityManager.SetDefaultComponents(func() engine.IComponent {
		return components.IdentityTransform()
	})

	g.Context.EntityManager.AddSystems(engine.PlayerSystemGroup,
		playersystems.NewTestPlayerSystem(),
		playersystems.NewPhysicsSystem())

	g.Context.EntityManager.AddSystems(engine.EngineSystemGroup,
		systems.NewLocalToWorldSystem(),
		systems.NewLocalToParentSystem(),
		systems.NewCameraSystem(),
		systems.NewRendererSystem())

	psystems := g.Context.EntityManager.GetSystemGroup(engine.PlayerSystemGroup)
	esystems := g.Context.EntityManager.GetSystemGroup(engine.EngineSystemGroup)

	psystems.Init(g.Context)
	esystems.Init(g.Context)

	previousFrameTime := glfw.GetTime()
	for !g.Window.ShouldClose() {
		now := glfw.GetTime()
		delta := float32(now - previousFrameTime)
		previousFrameTime = now

		loopStart := time.Now()
		g.Context.InputManager.Update()
		glfw.PollEvents()
		psystems.Update(g.Context, delta)
		esystems.Update(g.Context, delta)

		psystems.LateUpdate(g.Context, delta)
		esystems.LateUpdate(g.Context, delta)

		mainLoopTime := time.Since(loopStart)
		// Do OpenGL stuff.
		g.Window.SwapBuffers()

		g.Window.SetTitle(fmt.Sprintf("Render: %.2f - Main loop: %v", delta*1000, mainLoopTime))
	}

	psystems.Shutdown(g.Context)
	esystems.Shutdown(g.Context)
}

func CreateEngineContext() engine.EngineContext {
	return engine.EngineContext{
		EntityManager:   ecs.NewStandardEntityManager(),
		EventDispatcher: engine.NewDispatcher(),
		InputManager:    input.NewStandardInputManager(),
		ResourceManager: resourcemanager.New(),
	}
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	game := &Game{}
	defer game.Terminate()
	game.Init(CreateEngineContext(), 800, 600)
	game.Run()
}
