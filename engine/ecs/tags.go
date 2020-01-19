package ecs

type ComponentTag uint64

const (
	//TransformComponentTag ComponentTag = iota
	CameraComponentTag ComponentTag = iota
	MeshRendererComponentTag
	MaterialComponentTag
)
