package loader

import "github.com/qmuntal/gltf"

func LoadGLTF(path string) (*gltf.Document, error) {
	doc, err := gltf.Open(path)
	if err != nil {
		return doc, err
	}

	return doc, nil
}
