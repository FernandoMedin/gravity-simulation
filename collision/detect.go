package collision

import (
	"gravity-simulation/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Check(mainBody models.Body, body models.Body) bool {
	return rl.CheckCollisionSpheres(mainBody.Position, mainBody.Radius, body.Position, body.Radius)
}
