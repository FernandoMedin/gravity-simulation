package models

const G = 1.0     // Gravitational constant
const Theta = 0.9 // Opening angle
const Width = 1400
const Height = 900

func DefaultBounds() Box3D {
	return Box3D{X: -300, Y: -300, Z: -300, Width: 900, Height: 900, Depth: 900}
}
