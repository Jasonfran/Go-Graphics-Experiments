package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/ecs"
	"GraphicsStuff/primitives"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type RendererSystem struct {
	*ecs.SystemEntityCollection
	vao     uint32
	vbo     uint32
	program uint32
	cube    []engine.Vertex
}

func (r *RendererSystem) Init() {
	log.Println("RendererSystem init")
	glfw.SwapInterval(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	program, err := newProgram("shaders/vert.glsl", "shaders/frag.glsl")
	if err != nil {
		log.Fatal(err)
	}
	r.program = program
	gl.GenVertexArrays(1, &r.vao)
	gl.BindVertexArray(r.vao)

	gl.GenBuffers(1, &r.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	r.cube = primitives.CubeVertex()
	gl.BufferData(gl.ARRAY_BUFFER, int(unsafe.Sizeof(engine.Vertex{}))*len(r.cube), gl.Ptr(r.cube), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(unsafe.Sizeof(engine.Vertex{})), gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(75.0), float32(800)/600, 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	view := mgl32.LookAtV(mgl32.Vec3{0, 1, 0}, mgl32.Vec3{0, 1, 1}, mgl32.Vec3{0, 1, 0})
	viewUniform := gl.GetUniformLocation(program, gl.Str("view\x00"))
	gl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	model := mgl32.Ident4() //mgl32.Translate3D(3, 3, 3)
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	engine.EventDispatcher.Subscribe(engine.TestEvent, func(eventType engine.EventType, i interface{}) {
		log.Println(i)
	})

	time.AfterFunc(5*time.Second, func() {
		engine.EventDispatcher.Trigger(engine.TestEvent, "This works!")
	})
}

func (r *RendererSystem) Update(delta float32) {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, entity := range r.Entities() {
		camera, err := entity.GetCameraComponent()
		if err != nil {
			continue
		}

		for _, renderable := range camera.Renderables {
			transform := renderable.Transform()
			modelUniform := gl.GetUniformLocation(r.program, gl.Str("model\x00"))
			gl.UniformMatrix4fv(modelUniform, 1, false, &transform.LocalToWorld[0])
			gl.BindVertexArray(r.vao)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}
	}
}

func (r *RendererSystem) LateUpdate(delta float32) {
}

func (r *RendererSystem) Shutdown() {
	log.Println("RendererSystem shutdown")
}

func NewRendererSystem() *RendererSystem {
	system := &RendererSystem{SystemEntityCollection: ecs.NewSystemEntityCollection()}
	system.SetRequirements(ecs.CameraComponentTag)
	return system
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {

	vertShader, err := ioutil.ReadFile(vertexShaderSource)
	if err != nil {
		return 0, err
	}

	fragShader, err := ioutil.ReadFile(fragmentShaderSource)
	if err != nil {
		return 0, err
	}

	vertexShader, err := compileShader(string(vertShader), gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(string(fragShader), gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	length := int32(len(source))
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, &length)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
