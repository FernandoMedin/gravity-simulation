package main

import (
	"fmt"
	"gravity-simulation/collision"
	"gravity-simulation/commands"
	"gravity-simulation/models"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	buffer := models.Buffer{}
	rl.InitWindow(models.Width, models.Height, "Gravity Simulation - Octree 'space time'")
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	camera := rl.Camera{
		Position:   rl.NewVector3(10, 30, 400),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       50.0,
		Projection: rl.CameraPerspective,
	}

	root := &models.Octree{
		Bounds: models.DefaultBounds(),
	}

	central := models.Body{
		Position:   rl.NewVector3(0, 0, 0),
		Mass:       30000.0,
		Radius:     30.0,
		Color:      rl.Yellow,
		Collisions: 0,
	}
	root.Insert(&central)

	bodies := []models.Body{}

	for !rl.WindowShouldClose() {
		// Register/Execute input in buffer
		commands.GetInput(&buffer, &bodies, &central)

		dt := float32(0.1)

		for i := range bodies {
			force := root.ComputeForce(&bodies[i])

			// acceleration = F / m
			acc := rl.Vector3Scale(force, models.G/bodies[i].Mass)

			// integrate velocity and position
			bodies[i].Velocity = rl.Vector3Add(bodies[i].Velocity, rl.Vector3Scale(acc, dt))
			bodies[i].Position = rl.Vector3Add(bodies[i].Position, rl.Vector3Scale(bodies[i].Velocity, dt))

			if collision.Check(central, bodies[i]) {
				bodies[i].Removed = true
			}

			// root.CheckCollisionsForBody(&bodies[i], func(collided *models.Body) {
			// 	bodies[i].Removed = true
			// })
		}

		// Rebuild slice and tree without removed bodies
		root.Reset(models.DefaultBounds())
		root.Insert(&central)

		var newBodies []models.Body
		for i := range bodies {
			if !bodies[i].Removed {
				newBodies = append(newBodies, bodies[i])
				root.Insert(&newBodies[len(newBodies)-1])
			}
		}
		bodies = newBodies

		// rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(camera)
		root.Draw()
		rl.EndMode3D()

		rl.DrawText(buffer.MenuOptions, 10, 10, 20, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("Bodies Count: %d", len(bodies)), 10, 40, 20, rl.DarkGray)

		rl.DrawText(fmt.Sprintf("Command Buffer: %s", strings.Join(buffer.Text[:], " - ")), 10, models.Height-30, 20, rl.DarkGray)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
