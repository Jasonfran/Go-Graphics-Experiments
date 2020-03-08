package loader

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

type Vec2 struct {
	X float32
	Y float32
}

type Vertex struct {
	Pos       Vec3
	Normal    Vec3
	TexCoords Vec2
}

type Face struct {
	Indices [3]struct {
		Vertex   int
		Normal   int
		TexCoord int
	}
}

type Mesh struct {
	Vertices []Vertex
}

type Model struct {
	Meshes []Mesh
}

type Parser struct {
	Meshes    []Mesh
	vertices  []Vec3
	normals   []Vec3
	texCoords []Vec2
	faces     []Face
}

func (p *Parser) parse(input string) (*Model, error) {
	l := Lex(input)
L:
	for {
		switch item := l.nextItem(); {
		case item.Type == ItemDataDefinition:
			if item.Value == "v" {
				p.vertices = append(p.vertices, p.parseVec3(l))
			} else if item.Value == "vt" {
				p.texCoords = append(p.texCoords, p.parseVec2(l))
			} else if item.Value == "vn" {
				p.normals = append(p.normals, p.parseVec3(l))
			} else if item.Value == "f" {
				p.faces = append(p.faces, p.parseFace(l))
			}
		case item.Type == ItemEOF:
			break L
		}
	}

	vertices := make([]Vertex, 0, len(p.faces))
	for _, face := range p.faces {
		for _, index := range face.Indices {
			vert := Vertex{
				Pos:       p.vertices[index.Vertex],
				Normal:    p.normals[index.Normal],
				TexCoords: p.texCoords[index.TexCoord],
			}
			vertices = append(vertices, vert)
		}
	}

	return &Model{Meshes: []Mesh{{vertices}}}, nil
}

func (p *Parser) parseVec3(l *Lexer) Vec3 {
	i1 := l.nextItem()
	i2 := l.nextItem()
	i3 := l.nextItem()

	if i1.Type != ItemLineData || i2.Type != ItemLineData || i3.Type != ItemLineData {
		log.Fatal("Bad data")
	}

	v1, err := strconv.ParseFloat(i1.Value, 32)
	v2, err := strconv.ParseFloat(i2.Value, 32)
	v3, err := strconv.ParseFloat(i3.Value, 32)
	if err != nil {
		log.Fatal("Bad data")
	}

	return Vec3{float32(v1), float32(v2), float32(v3)}
}

func (p *Parser) parseVec2(l *Lexer) Vec2 {
	i1 := l.nextItem()
	i2 := l.nextItem()

	if i1.Type != ItemLineData || i2.Type != ItemLineData {
		log.Fatal("Bad data")
	}

	v1, err := strconv.ParseFloat(i1.Value, 32)
	v2, err := strconv.ParseFloat(i2.Value, 32)
	if err != nil {
		log.Fatal("Bad data")
	}

	return Vec2{float32(v1), float32(v2)}
}

func (p *Parser) parseFace(l *Lexer) Face {
	data := l.nextItem()
	if data.Type != ItemLineData {
		log.Fatal("Bad data")
	}

	face := Face{
		Indices: [3]struct {
			Vertex   int
			Normal   int
			TexCoord int
		}{},
	}
	fields := strings.Fields(data.Value)
	for i, field := range fields {
		split := strings.Split(field, "/")
		if len(split) > 2 {
			v, err := strconv.ParseInt(split[0], 10, 32)
			vtStr := split[1]
			vn, err := strconv.ParseInt(split[2], 10, 32)
			if err != nil {
				log.Fatal("Bad data")
			}
			face.Indices[i].Vertex = int(v)
			if len(strings.TrimSpace(vtStr)) > 0 {
				vt, err := strconv.ParseInt(vtStr, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				face.Indices[i].TexCoord = int(vt)
			}

			face.Indices[i].Normal = int(vn)
		}
	}
	return face
}

func Parse(file string) (*Model, error) {
	p := &Parser{
		Meshes:    []Mesh{},
		vertices:  []Vec3{},
		normals:   []Vec3{},
		texCoords: []Vec2{},
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return p.parse(string(bytes))
}
