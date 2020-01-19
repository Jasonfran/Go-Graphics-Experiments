package main

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/component"
	"GraphicsStuff/engine/ecsmanager"
	"GraphicsStuff/engine/systems"
	"GraphicsStuff/playersystems"
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

type Game struct {
	Window *glfw.Window
	Width  int
	Height int
}

func (g *Game) Init(width int, height int) {
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
	window.SetFramebufferSizeCallback(g.framebufferSizeCallback)
	window.SetKeyCallback(engine.InputManager.KeyCallback)
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

func (g *Game) framebufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	g.Width = width
	g.Height = height
}

func (g *Game) Run() {
	testplayersystem := playersystems.NewTestPlayerSystem()
	testplayersystem.SetRequirements()
	engine.ECSManager.AddSystem(ecsmanager.PlayerSystemGroup, testplayersystem)

	renderer := systems.NewRendererSystem()
	renderer.SetRequirements(component.MeshRendererComponentTag)

	cameraSystem := systems.NewCameraSystem()
	cameraSystem.SetRequirements(component.CameraComponentTag)

	engine.ECSManager.AddSystem(ecsmanager.EngineSystemGroup, cameraSystem)
	engine.ECSManager.AddSystem(ecsmanager.EngineSystemGroup, renderer)

	psystems := engine.ECSManager.GetSystemGroup(ecsmanager.PlayerSystemGroup)
	esystems := engine.ECSManager.GetSystemGroup(ecsmanager.EngineSystemGroup)

	psystems.Init()
	esystems.Init()

	previousFrameTime := glfw.GetTime()
	for !g.Window.ShouldClose() {
		now := glfw.GetTime()
		delta := float32(now - previousFrameTime)
		previousFrameTime = now

		engine.InputManager.Update()
		glfw.PollEvents()
		psystems.Update(delta)
		esystems.Update(delta)

		psystems.LateUpdate(delta)
		esystems.LateUpdate(delta)

		// Do OpenGL stuff.
		g.Window.SwapBuffers()

		g.Window.SetTitle(fmt.Sprintf("Render: %.2f", delta*1000))
	}

	psystems.Shutdown()
	esystems.Shutdown()
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	game := &Game{}
	defer game.Terminate()
	game.Init(800, 600)
	game.Run()
}
