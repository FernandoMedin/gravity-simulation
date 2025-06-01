package commands

import (
	"gravity-simulation/models"
	"gravity-simulation/planets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SpawDust(central models.Body, n int) []models.Body {
	bodies := []models.Body{}

	for i := 0; i < n; i++ {
		x := float32(rl.GetRandomValue(200, 3000)) / 10.0
		y := float32(rl.GetRandomValue(-500, 500)) / 10.0
		z := float32(rl.GetRandomValue(-1000, 1000)) / 10.0

		// if rl.GetRandomValue(0, 1) == 0 {
		// 	x = -x
		// }

		radius := float32(0.2)
		pos := rl.NewVector3(x, y, z)
		mass := float32(rl.GetRandomValue(1, 20)) / 10.0

		// Generate Dust? Generate Body?
		bodies = append(bodies, planets.GeneratePlanet(
			central,
			pos,
			mass,
			radius,
			rl.White,
		))
	}

	return bodies
}
