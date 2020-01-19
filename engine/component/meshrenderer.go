package component

type MeshRenderer struct {
	Mesh string
}

func (m *MeshRenderer) Name() string {
	return "MeshRenderer"
}
