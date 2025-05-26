package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Position   rl.Vector3
	Velocity   rl.Vector3
	Mass       float32
	Radius     float32
	Removed    bool // Flag to mark the body for removal (workaround process for now)
	Color      rl.Color
	Collisions int
}

type Box3D struct {
	X, Y, Z              float32
	Width, Height, Depth float32
}
