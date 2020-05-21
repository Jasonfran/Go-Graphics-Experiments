package systems

import (
	"GraphicsStuff/engine"
	"GraphicsStuff/engine/components"
	"GraphicsStuff/engine/data"
	"GraphicsStuff/engine/events"
	"GraphicsStuff/engine/shader"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type RendererSystem struct {
	//vao         uint32
	//vbo         uint32
	basicShader *shader.Shader
	cube        []engine.Vertex
}

func (r *RendererSystem) Init(context engine.EngineContext) {
	log.Println("RendererSystem init")
	glfw.SwapInterval(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	basicShader, err := shader.New("shaders/vert.glsl", "shaders/frag.glsl")
	if err != nil {
		log.Fatal(err)
	}
	r.basicShader = basicShader
	//gl.GenVertexArrays(1, &r.vao)
	//gl.BindVertexArray(r.vao)
	//
	//gl.GenBuffers(1, &r.vbo)
	//gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	//
	//r.cube = primitives.CubeVertex()
	//gl.BufferData(gl.ARRAY_BUFFER, int(unsafe.Sizeof(engine.Vertex{}))*len(r.cube), gl.Ptr(r.cube), gl.STATIC_DRAW)
	//gl.EnableVertexAttribArray(0)
	//gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(unsafe.Sizeof(engine.Vertex{})), gl.PtrOffset(0))
	//
	//gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	//gl.BindVertexArray(0)
	basicShader.Use()

	projection := mgl32.Perspective(mgl32.DegToRad(75.0), float32(800)/600, 0.1, 1000.0)
	basicShader.SetMat4("projection", projection)

	view := mgl32.LookAtV(mgl32.Vec3{0, 1, 0}, mgl32.Vec3{0, 1, 1}, mgl32.Vec3{0, 1, 0})
	basicShader.SetMat4("view", view)

	model := mgl32.Ident4() //mgl32.Translate3D(3, 3, 3)
	basicShader.SetMat4("model", model)

	context.EventDispatcher.Subscribe(events.WindowResizedEvent, func(eventType engine.EventType, d interface{}) {
		eventData, ok := d.(events.WindowResizedEventData)
		if !ok {
			return
		}

		gl.Viewport(0, 0, int32(eventData.Width), int32(eventData.Width))
		projection := mgl32.Perspective(mgl32.DegToRad(90.0), float32(eventData.Width)/float32(eventData.Height), 0.1, 1000.0)
		basicShader.SetMat4("projection", projection)
	})
}

func (r *RendererSystem) Update(context engine.EngineContext, delta float32) {
	//defer func(t time.Time) {
	//	log.Println(time.Since(t))
	//}(time.Now())

	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	drawCalls := 0
	context.EntityManager.GetEntitiesFromQuery(components.Includes(components.CameraComponentTag)).Each(func(entity engine.IEntity) {
		camera, err := components.GetCameraComponent(entity)
		if err != nil {
			return
		}

		for _, renderable := range camera.Renderables {
			transform, _ := components.GetTransformComponent(renderable)
			r.basicShader.SetMat4("model", transform.LocalToWorld)
			//log.Println(transform.LocalToWorld)
			meshComponent, err := components.GetMeshComponent(renderable)
			if err == nil {
				if !meshComponent.Mesh.LoadedIntoGL {
					meshComponent.Mesh.LoadIntoGL()
				}
				for _, primitive := range meshComponent.Mesh.Primitives {
					color := mgl32.Vec3{1, 1, 1}
					material := &data.Material{Colour: color}
					if primitive.Material != nil {
						material = primitive.Material
					}
					r.basicShader.SetVec3("color", material.Colour)
					gl.BindVertexArray(primitive.VAO)
					gl.DrawElements(primitive.Mode, int32(primitive.Indices.Count), primitive.Indices.ComponentType, nil)
					drawCalls += 1
					gl.BindVertexArray(0)
				}
			}
		}
	})
	//log.Println("Draw calls:", drawCalls)
}

func (r *RendererSystem) LateUpdate(context engine.EngineContext, delta float32) {
}

func (r *RendererSystem) Shutdown(context engine.EngineContext) {
	log.Println("RendererSystem shutdown")
}

func NewRendererSystem() *RendererSystem {
	return &RendererSystem{}
}
