package commands

import (
	"gravity-simulation/models"
	"gravity-simulation/planets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SpawPlanet(central models.Body) models.Body {
	var color rl.Color

	x := float32(rl.GetRandomValue(200, 2000)) / 10.0
	y := float32(rl.GetRandomValue(-1000, 1000)) / 10.0
	z := float32(rl.GetRandomValue(-500, 500)) / 10.0

	// if rl.GetRandomValue(0, 1) == 0 {
	// 	x = -x
	// }

	pos := rl.NewVector3(x, y, z)
	mass := float32(rl.GetRandomValue(30, 150))
	radius := float32(3)

	if mass > 140 {
		color = rl.Brown
		radius = float32(6)
	} else if mass > 90 {
		color = rl.Orange
		radius = float32(5)
	} else if mass > 60 {
		color = rl.Purple
		radius = float32(4)
	} else {
		color = rl.Green
	}

	return planets.GeneratePlanet(
		central,
		pos,
		mass,
		radius,
		color,
	)
}
