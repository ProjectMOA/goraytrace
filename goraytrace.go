package main

import (
	"github.com/ProjectMOA/goraytrace/math3d"
	"github.com/ProjectMOA/goraytrace/scene"
	"github.com/ProjectMOA/goraytrace/shape"
)

func main() {
	// Setting up a scene
	myScene := scene.New()
	aSphere := shape.Sphere{
		Position: math3d.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Radius:   1.0}

	myScene.Add(&aSphere)
}
