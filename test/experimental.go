package test

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Position rl.Vector2
	Mass     float32
}

type Quadtree struct {
	Boundary       rl.Rectangle
	Body           *Body
	Divided        bool
	NW, NE, SW, SE *Quadtree
}

func NewQuadtree(boundary rl.Rectangle) *Quadtree {
	return &Quadtree{
		Boundary: boundary,
	}
}

func (qt *Quadtree) Draw() {
	// Draw the boundary
	rl.DrawRectangleLines(int32(qt.Boundary.X), int32(qt.Boundary.Y),
		int32(qt.Boundary.Width), int32(qt.Boundary.Height), rl.DarkGray)

	// Draw the body if it exists
	if qt.Body != nil {
		if qt.Body.Mass > 8 {
			rl.DrawCircleV(qt.Body.Position, 4, rl.Red)
		} else if qt.Body.Mass > 4 {
			rl.DrawCircleV(qt.Body.Position, 3, rl.Blue)
		} else {
			rl.DrawCircleV(qt.Body.Position, 2, rl.Green)
		}
	}

	// Recursively draw children if subdivided
	if qt.Divided {
		qt.NW.Draw()
		qt.NE.Draw()
		qt.SW.Draw()
		qt.SE.Draw()
	}
}

func (qt *Quadtree) Subdivide() {
	x := qt.Boundary.X
	y := qt.Boundary.Y
	w := qt.Boundary.Width / 2
	h := qt.Boundary.Height / 2

	qt.NW = &Quadtree{Boundary: rl.NewRectangle(x, y, w, h)}
	qt.NE = &Quadtree{Boundary: rl.NewRectangle(x+w, y, w, h)}
	qt.SW = &Quadtree{Boundary: rl.NewRectangle(x, y+h, w, h)}
	qt.SE = &Quadtree{Boundary: rl.NewRectangle(x+w, y+h, w, h)}

	qt.Divided = true
}

func (qt *Quadtree) Insert(body *Body) bool {
	// Check if body is inside this node's boundary
	if !rl.CheckCollisionPointRec(body.Position, qt.Boundary) {
		return false
	}

	// If no body exists, place this one here
	if qt.Body == nil && !qt.Divided {
		qt.Body = body
		return true
	}

	// If already divided, insert into a child
	if !qt.Divided {
		qt.Subdivide()
	}

	// Try inserting into children
	return qt.NW.Insert(body) || qt.NE.Insert(body) || qt.SW.Insert(body) || qt.SE.Insert(body)
}

func GenerateRandomBodies() []Body {
	var bodies []Body

	for i := 0; i < 200; i++ {
		bodies = append(bodies, Body{
			Position: rl.Vector2{X: float32(rl.GetRandomValue(0, 800)), Y: float32(rl.GetRandomValue(0, 800))},
			Mass:     float32(rl.GetRandomValue(1, 10)),
		})
	}

	return bodies
}

func experimental() {
	rl.InitWindow(800, 800, "Barnes-Hut Quadtree")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var bodies []Body = GenerateRandomBodies()
	var root *Quadtree

	i := 0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if i < len(bodies) {
			if root == nil {
				root = NewQuadtree(rl.NewRectangle(0, 0, 800, 800))
			}
			root.Insert(&bodies[i])
			i++
		}

		if root != nil {
			root.Draw()
		}

		rl.EndDrawing()
	}
}
