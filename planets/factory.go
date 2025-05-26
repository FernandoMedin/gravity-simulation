package planets

import (
	"gravity-simulation/models"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GeneratePlanet(central models.Body, distance rl.Vector3, mass float32, radius float32, color rl.Color) models.Body {
	dirToCenter := rl.Vector3Subtract(central.Position, distance)
	dist := rl.Vector3Length(dirToCenter)

	speed := float32(math.Sqrt(float64(models.G * central.Mass / dist)))

	// Tangent direction (e.g., along +Z for orbit in XZ plane)
	tangent := rl.NewVector3(0, 0, 1)
	velocity := rl.Vector3Scale(tangent, speed)

	orbitingBody := models.Body{
		Position:   distance,
		Velocity:   velocity,
		Mass:       mass,
		Radius:     radius,
		Color:      color,
		Collisions: 0,
	}

	return orbitingBody
}

func GenerateMoon(planet models.Body) models.Body {
	moonDistance := float32(5.0) // distance from planet
	moonPos := rl.Vector3Add(planet.Position, rl.NewVector3(moonDistance, 0, 0))

	dirToPlanet := rl.Vector3Subtract(planet.Position, moonPos)
	moonDist := rl.Vector3Length(dirToPlanet)

	moonSpeed := float32(math.Sqrt(float64(models.G * planet.Mass / moonDist)))

	// Tangent direction (e.g., orbit in XZ plane around the planet)
	moonTangent := rl.NewVector3(0, 0, 1)
	moonOrbitalVelocity := rl.Vector3Scale(moonTangent, moonSpeed)

	// Total moon velocity = planet velocity + moon orbital velocity
	totalMoonVelocity := rl.Vector3Add(planet.Velocity, moonOrbitalVelocity)

	moon := models.Body{
		Position: moonPos,
		Velocity: totalMoonVelocity,
		Mass:     1.0,
		Radius:   0.4,
		Color:    rl.Gray,
	}

	return moon
}
