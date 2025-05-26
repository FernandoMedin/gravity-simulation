package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Octree struct {
	Bounds       Box3D
	Body         *Body
	Divided      bool
	Children     [8]*Octree
	TotalMass    float32
	CenterOfMass rl.Vector3
}

func (qt *Octree) ComputeForce(body *Body) rl.Vector3 {
	if qt.Body == body || qt.TotalMass == 0 {
		return rl.NewVector3(0, 0, 0)
	}

	dir := rl.Vector3Subtract(qt.CenterOfMass, body.Position)
	dist := rl.Vector3Length(dir) + 0.1 // softening factor to avoid singularity
	if dist == 0 {
		return rl.NewVector3(0, 0, 0)
	}
	dir = rl.Vector3Normalize(dir)

	size := qt.Bounds.Width // assume cube
	if !qt.Divided || (size/dist) < Theta {
		forceMag := G * body.Mass * qt.TotalMass / (dist * dist)
		return rl.Vector3Scale(dir, forceMag)
	}

	force := rl.NewVector3(0, 0, 0)
	for _, child := range qt.Children {
		if child != nil {
			childForce := child.ComputeForce(body)
			force = rl.Vector3Add(force, childForce)
		}
	}
	return force
}

func (qt *Octree) Clear() {
	qt.Body = nil
	qt.TotalMass = 0
	qt.CenterOfMass = rl.NewVector3(0, 0, 0)

	if qt.Divided {
		for _, child := range qt.Children {
			if child != nil {
				child.Clear()
			}
		}
	}
}

func (qt *Octree) Reset(bounds Box3D) {
	qt.Bounds = bounds
	qt.Body = nil
	qt.TotalMass = 0
	qt.CenterOfMass = rl.NewVector3(0, 0, 0)
	qt.Divided = false

	for i := range qt.Children {
		qt.Children[i] = nil
	}
}

func (qt *Octree) Subdivide() {
	x, y, z := qt.Bounds.X, qt.Bounds.Y, qt.Bounds.Z
	w, h, d := qt.Bounds.Width/2, qt.Bounds.Height/2, qt.Bounds.Depth/2

	offsets := [][3]float32{
		{0, 0, 0}, {w, 0, 0}, {0, h, 0}, {w, h, 0},
		{0, 0, d}, {w, 0, d}, {0, h, d}, {w, h, d},
	}

	for i, off := range offsets {
		qt.Children[i] = &Octree{
			Bounds: Box3D{
				X: x + off[0], Y: y + off[1], Z: z + off[2],
				Width: w, Height: h, Depth: d,
			},
		}
	}

	qt.Divided = true
}

func (qt *Octree) Insert(body *Body) bool {
	if !qt.Bounds.Contains(body.Position) {
		return false
	}

	inserted := false

	if qt.Body == nil && !qt.Divided {
		qt.Body = body
		inserted = true
	} else {
		if !qt.Divided {
			qt.Subdivide()
			// Move existing body to a child
			if qt.Body != nil {
				oldBody := qt.Body
				qt.Body = nil
				for _, child := range qt.Children {
					if child.Insert(oldBody) {
						break
					}
				}
			}
		}

		for _, child := range qt.Children {
			if child.Insert(body) {
				inserted = true
				break
			}
		}
	}

	if inserted {
		// Update COM and total mass
		totalMassBefore := qt.TotalMass
		qt.TotalMass += body.Mass
		if totalMassBefore == 0 {
			qt.CenterOfMass = body.Position
		} else {
			qt.CenterOfMass = rl.Vector3{
				X: (qt.CenterOfMass.X*totalMassBefore + body.Position.X*body.Mass) / qt.TotalMass,
				Y: (qt.CenterOfMass.Y*totalMassBefore + body.Position.Y*body.Mass) / qt.TotalMass,
				Z: (qt.CenterOfMass.Z*totalMassBefore + body.Position.Z*body.Mass) / qt.TotalMass,
			}
		}
	}

	return inserted
}

func (qt *Octree) Draw(drawBounds bool) {

	if drawBounds {
		b := qt.Bounds
		center := rl.NewVector3(b.X+b.Width/2, b.Y+b.Height/2, b.Z+b.Depth/2)
		size := rl.NewVector3(b.Width, b.Height, b.Depth)
		rl.DrawCubeWires(center, size.X, size.Y, size.Z, rl.Gray)
	}

	if qt.Body != nil {
		rl.DrawSphere(qt.Body.Position, qt.Body.Radius, qt.Body.Color)
	}

	if qt.Divided {
		for _, child := range qt.Children {
			child.Draw(drawBounds)
		}
	}
}

func (tree *Octree) Remove(body *Body) bool {
	if tree == nil || body == nil {
		return false
	}

	if tree.Body == body {
		tree.Body = nil
		tree.TotalMass = 0
		tree.CenterOfMass = rl.NewVector3(0, 0, 0)
		return true
	}

	if tree.Divided {
		for _, child := range tree.Children {
			if child != nil && child.Remove(body) {
				return true
			}
		}
	}

	return false
}

func (qt *Octree) QueryBox(box Box3D, callback func(*Body)) {
	if !qt.Bounds.Intersects(box) {
		return
	}

	if qt.Body != nil && box.Contains(qt.Body.Position) {
		callback(qt.Body)
	}

	if qt.Divided {
		for _, child := range qt.Children {
			if child != nil {
				child.QueryBox(box, callback)
			}
		}
	}
}

func (b Box3D) Contains(p rl.Vector3) bool {
	return p.X >= b.X && p.X < b.X+b.Width &&
		p.Y >= b.Y && p.Y < b.Y+b.Height &&
		p.Z >= b.Z && p.Z < b.Z+b.Depth
}

func (b Box3D) Intersects(other Box3D) bool {
	return b.X < other.X+other.Width && b.X+b.Width > other.X &&
		b.Y < other.Y+other.Height && b.Y+b.Height > other.Y &&
		b.Z < other.Z+other.Depth && b.Z+b.Depth > other.Z
}

func (qt *Octree) CheckCollisionsForBody(mainBody *Body, handler func(*Body)) {
	box := Box3D{
		X:      mainBody.Position.X - mainBody.Radius,
		Y:      mainBody.Position.Y - mainBody.Radius,
		Z:      mainBody.Position.Z - mainBody.Radius,
		Width:  mainBody.Radius * 2,
		Height: mainBody.Radius * 2,
		Depth:  mainBody.Radius * 2,
	}

	qt.QueryBox(box, func(other *Body) {
		if mainBody != other && CheckCollisionDebug(*mainBody, *other) {
			handler(other)
		}
	})
}

func CheckCollisionDebug(mainBody Body, body Body) bool {
	return rl.CheckCollisionSpheres(mainBody.Position, mainBody.Radius, body.Position, body.Radius)
}
