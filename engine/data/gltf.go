package data

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/qmuntal/gltf"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type AccessorDataType uint

const (
	AccessorScalar AccessorDataType = iota
	AccessorVec2
	AccessorVec3
	AccessorVec4
	AccessorMat2
	AccessorMat3
	AccessorMat4
)

type Model struct {
	Loaded  bool
	Buffers []Buffer
	Nodes   []*Node
}

func (m *Model) LoadIntoGl() {
	loadModelIntoGl(m)
}

func (m *Model) GatherNodes() []*Node {
	var nodes []*Node
	for _, node := range m.Nodes {
		nodes = append(nodes, node.GatherNodes()...)
	}

	return nodes
}

type Buffer struct {
	Data []byte
}

type Accessor struct {
	Buffer        *Buffer
	GLBuffer      uint32
	initialised   bool
	Offset        uint32
	ComponentType uint32
	Count         uint32
	ByteStride    uint32
	DataType      AccessorDataType
	BufferTarget  uint32
}

func (a *Accessor) IsInitialised() bool {
	return a.initialised
}

func (a *Accessor) GetData() []byte {
	componentSize := getComponentSize(a)
	componentCount := getComponentCount(a)
	byteLength := componentSize * componentCount * a.Count
	return a.Buffer.Data[a.Offset : a.Offset+byteLength]
}

type Node struct {
	Name        string
	Children    []*Node
	Mesh        *Mesh
	Translation mgl32.Vec3
	Rotation    mgl32.Quat
	Scale       mgl32.Vec3
	Matrix      mgl32.Mat4
	HasMatrix   bool
}

func (n *Node) GatherNodes() []*Node {
	var nodes []*Node
	nodes = append(nodes, n)
	for _, child := range n.Children {
		nodes = append(nodes, child.GatherNodes()...)
	}

	return nodes
}

type Mesh struct {
	Name         string
	Primitives   []*Primitive
	LoadedIntoGL bool
}

func (m *Mesh) LoadIntoGL() {
	for _, prim := range m.Primitives {
		loadPrimitiveIntoGl(prim)
	}
	m.LoadedIntoGL = true
}

type Primitive struct {
	VAO        uint32
	Indices    *Accessor
	Attributes map[string]*Accessor
	Mode       uint32
	Material   *Material
}

func LoadGLTF(path string) (*Model, error) {
	doc, err := gltf.Open(path)
	if err != nil {
		return nil, err
	}

	if len(doc.Scenes) == 0 {
		return nil, fmt.Errorf("no scene found")
	}

	if len(doc.Scenes) > 1 {
		return nil, fmt.Errorf("multiple scenes not supported")
	}

	for _, extension := range doc.ExtensionsRequired {
		if strings.EqualFold(extension, "KHR_materials_pbrSpecularGlossiness") {
			log.Println("Currently no support for KHR_materials_pbrSpecularGlossiness materials")
		}
	}

	var nodes []*Node
	var buffers []*Buffer
	scene := doc.Scenes[0]

	for _, buffer := range doc.Buffers {
		buffers = append(buffers, &Buffer{
			Data: buffer.Data,
		})
	}

	for _, node := range scene.Nodes {
		nodes = append(nodes, loadNode(doc, doc.Nodes[node], buffers))
	}

	//for _, node := range nodes {
	//	transformNode(mgl32.Ident4(), node)
	//}

	model := &Model{
		Nodes: nodes,
	}

	return model, nil
}

func loadNode(doc *gltf.Document, node *gltf.Node, buffers []*Buffer) *Node {
	matrix := node.MatrixOrDefault()
	translation := node.TranslationOrDefault()
	rotation := node.RotationOrDefault()
	scale := node.ScaleOrDefault()
	glMatrix := mgl32.Mat4{
		float32(matrix[0]), float32(matrix[1]), float32(matrix[2]), float32(matrix[3]),
		float32(matrix[4]), float32(matrix[5]), float32(matrix[6]), float32(matrix[7]),
		float32(matrix[8]), float32(matrix[9]), float32(matrix[10]), float32(matrix[11]),
		float32(matrix[12]), float32(matrix[13]), float32(matrix[14]), float32(matrix[15])}

	loadedNode := &Node{
		Name:        node.Name,
		Translation: mgl32.Vec3{float32(translation[0]), float32(translation[1]), float32(translation[2])},
		Rotation: mgl32.Quat{
			W: float32(rotation[3]),
			V: mgl32.Vec3{
				float32(rotation[0]),
				float32(rotation[1]),
				float32(rotation[2]),
			},
		},
		Scale:     mgl32.Vec3{float32(scale[0]), float32(scale[1]), float32(scale[2])},
		HasMatrix: matrix != gltf.DefaultMatrix,
		Matrix:    glMatrix,
	}

	//translationMatrix := mgl32.Translate3D(loadedNode.Translation.X(), loadedNode.Translation.Y(), loadedNode.Translation.Z())
	//rotationMatrix := loadedNode.Rotation.Mat4()
	//scaleMatrix := mgl32.Scale3D(loadedNode.Scale.X(), loadedNode.Scale.Y(), loadedNode.Scale.Z())
	////loadedNode.Matrix = loadedNode.Matrix.Mul4(translationMatrix).Mul4(rotationMatrix).Mul4(scaleMatrix)

	if node.Mesh != nil {
		loadedNode.Mesh = loadMesh(doc, doc.Meshes[*node.Mesh], buffers)
	}

	for _, child := range node.Children {
		loadedNode.Children = append(loadedNode.Children, loadNode(doc, doc.Nodes[child], buffers))
	}

	return loadedNode
}

func loadMesh(doc *gltf.Document, mesh *gltf.Mesh, buffers []*Buffer) *Mesh {
	loadedMesh := &Mesh{
		Name: mesh.Name,
	}

	for _, primitive := range mesh.Primitives {
		attributes := map[string]*Accessor{}
		for name, attr := range primitive.Attributes {
			attributes[name] = loadAccessor(doc, doc.Accessors[attr], buffers)
		}

		mode := gl.TRIANGLES
		switch primitive.Mode {
		case gltf.PrimitiveTriangles:
			mode = gl.TRIANGLES
		case gltf.PrimitivePoints:
			mode = gl.POINTS
		case gltf.PrimitiveLines:
			mode = gl.LINES
		case gltf.PrimitiveLineLoop:
			mode = gl.LINE_LOOP
		case gltf.PrimitiveLineStrip:
			mode = gl.LINE_STRIP
		case gltf.PrimitiveTriangleStrip:
			mode = gl.TRIANGLE_STRIP
		case gltf.PrimitiveTriangleFan:
			mode = gl.TRIANGLE_FAN
		}

		loadedPrimitive := &Primitive{
			Attributes: attributes,
			Mode:       uint32(mode),
		}

		if primitive.Indices != nil {
			loadedPrimitive.Indices = loadAccessor(doc, doc.Accessors[*primitive.Indices], buffers)
		}

		loadedMesh.Primitives = append(loadedMesh.Primitives, loadedPrimitive)
	}

	return loadedMesh
}

func loadAccessor(doc *gltf.Document, accessor *gltf.Accessor, buffers []*Buffer) *Accessor {
	bufferView := doc.BufferViews[*accessor.BufferView]
	bufferIndex := bufferView.Buffer
	offset := bufferView.ByteOffset + accessor.ByteOffset
	count := accessor.Count
	componentType := gl.FLOAT

	switch accessor.ComponentType {
	case gltf.ComponentFloat:
		componentType = gl.FLOAT
	case gltf.ComponentByte:
		componentType = gl.BYTE
	case gltf.ComponentUbyte:
		componentType = gl.UNSIGNED_BYTE
	case gltf.ComponentShort:
		componentType = gl.SHORT
	case gltf.ComponentUshort:
		componentType = gl.UNSIGNED_SHORT
	case gltf.ComponentUint:
		componentType = gl.UNSIGNED_INT
	}

	byteStride := bufferView.ByteStride
	dataType := AccessorDataType(accessor.Type)
	buffer := buffers[bufferIndex]

	return &Accessor{
		Buffer:        buffer,
		Offset:        offset,
		ComponentType: uint32(componentType),
		ByteStride:    byteStride,
		Count:         count,
		DataType:      dataType,
		BufferTarget:  uint32(bufferView.Target),
	}
}

func loadModelIntoGl(model *Model) {
	for _, node := range model.Nodes {
		loadNodeIntoGl(node)
	}

	model.Loaded = true
}

func loadNodeIntoGl(node *Node) {
	mesh := node.Mesh
	if mesh != nil {
		for _, primitive := range mesh.Primitives {
			loadPrimitiveIntoGl(primitive)
		}
	}

	for _, child := range node.Children {
		loadNodeIntoGl(child)
	}
}

func loadPrimitiveIntoGl(primitive *Primitive) {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	position := primitive.Attributes["POSITION"]
	bindAccessor(position)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, position.ComponentType, false, int32(position.ByteStride), gl.PtrOffset(0))

	bindAccessor(primitive.Indices)
	gl.BindVertexArray(0)
	primitive.VAO = VAO
}

func bindAccessor(accessor *Accessor) {
	if !accessor.IsInitialised() {
		createAccessorGLBuffer(accessor.BufferTarget, accessor)
	}

	gl.BindBuffer(accessor.BufferTarget, accessor.GLBuffer)
}

func createAccessorGLBuffer(target uint32, accessor *Accessor) *Accessor {
	var bufferID uint32
	gl.GenBuffers(1, &bufferID)
	gl.BindBuffer(target, bufferID)
	gl.BufferData(target, len(accessor.GetData()), gl.Ptr(accessor.GetData()), gl.STATIC_DRAW)
	gl.BindBuffer(target, 0)
	accessor.initialised = true
	accessor.GLBuffer = bufferID
	return accessor
}

var componentCount = map[AccessorDataType]uint32{
	AccessorScalar: 1,
	AccessorVec2:   2,
	AccessorVec3:   3,
	AccessorVec4:   4,
	AccessorMat2:   4,
	AccessorMat3:   9,
	AccessorMat4:   16,
}

func getComponentCount(accessor *Accessor) uint32 {
	return componentCount[accessor.DataType]
}

func getComponentSize(accessor *Accessor) uint32 {
	switch accessor.ComponentType {
	case gl.BYTE:
		fallthrough
	case gl.UNSIGNED_BYTE:
		return 1
	case gl.SHORT:
		fallthrough
	case gl.UNSIGNED_SHORT:
		return 2
	case gl.UNSIGNED_INT:
		fallthrough
	case gl.FLOAT:
		return 4
	}
	return 0
}
