package shader

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	ID uint32
}

func New(vertexPath string, fragmentPath string) (*Shader, error) {
	vertShader, err := ioutil.ReadFile(vertexPath)
	if err != nil {
		return nil, err
	}

	fragShader, err := ioutil.ReadFile(fragmentPath)
	if err != nil {
		return nil, err
	}

	vertexShader, err := compileShader(string(vertShader), gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(string(fragShader), gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
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

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	shader := &Shader{
		ID: program,
	}

	return shader, nil
}

func (s *Shader) Use() {
	gl.UseProgram(s.ID)
}

func (s *Shader) SetBool(name string, value bool) {
	uniformLocation := s.getUniformLocation(name)
	val := 0
	if value {
		val = 1
	}

	gl.Uniform1i(uniformLocation, int32(val))
}

func (s *Shader) getUniformLocation(name string) int32 {
	return gl.GetUniformLocation(s.ID, gl.Str(name+"\x00"))
}

func (s *Shader) SetInt(name string, value int32) {
	uniformLocation := s.getUniformLocation(name)
	gl.Uniform1i(uniformLocation, value)
}

func (s *Shader) SetFloat(name string, value float32) {
	uniformLocation := s.getUniformLocation(name)
	gl.Uniform1f(uniformLocation, value)
}

func (s *Shader) SetVec2(name string, value mgl32.Vec2) {
	uniformLocation := s.getUniformLocation(name)
	gl.Uniform2fv(uniformLocation, 1, &value[0])
}

func (s *Shader) SetVec3(name string, value mgl32.Vec3) {
	uniformLocation := s.getUniformLocation(name)
	gl.Uniform3fv(uniformLocation, 1, &value[0])
}

func (s *Shader) SetVec4(name string, value mgl32.Vec4) {
	uniformLocation := s.getUniformLocation(name)
	gl.Uniform4fv(uniformLocation, 1, &value[0])
}

func (s *Shader) SetMat2(name string, value mgl32.Mat2) {
	uniformLocation := s.getUniformLocation(name)
	gl.UniformMatrix2fv(uniformLocation, 1, false, &value[0])
}

func (s *Shader) SetMat3(name string, value mgl32.Mat3) {
	uniformLocation := s.getUniformLocation(name)
	gl.UniformMatrix3fv(uniformLocation, 1, false, &value[0])
}

func (s *Shader) SetMat4(name string, value mgl32.Mat4) {
	uniformLocation := s.getUniformLocation(name)
	gl.UniformMatrix4fv(uniformLocation, 1, false, &value[0])
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
