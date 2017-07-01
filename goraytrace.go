package main

import (
	"math/rand"

	"github.com/ProjectMOA/goraytrace/image"
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

	anImage := image.New(600, 600)
	// Add some noise to the image
	for i := int32(0); i < 600; i++ {
		for j := int32(0); j < 600; j++ {
			anImage.Pixels[i][j] = image.Color{R: rand.Float64(), G: rand.Float64(), B: rand.Float64()}
		}
	}
	anImage.Save("test.ppm")
}
