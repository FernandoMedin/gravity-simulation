package utils

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Vector3MinValue(v rl.Vector3) float32 {
	return float32(math.Min(float64(v.X), math.Min(float64(v.Y), float64(v.Z))))
}

func Vector3MaxValue(v rl.Vector3) float32 {
	return float32(math.Max(float64(v.X), math.Max(float64(v.Y), float64(v.Z))))
}
