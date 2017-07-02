package main

import (
	"github.com/ProjectMOA/goraytrace/image"
	"github.com/ProjectMOA/goraytrace/lighting"
	"github.com/ProjectMOA/goraytrace/math3d"
	"github.com/ProjectMOA/goraytrace/scene"
	"github.com/ProjectMOA/goraytrace/shape"
)

func main() {
	// Setting up a scene
	myScene := scene.New()

	myScene.AddShape(&shape.Sphere{
		Position: math3d.Vector3{X: 0.0, Y: 0.0, Z: 1.2},
		Radius:   0.1})

	myScene.AddShape(&shape.Sphere{
		Position: math3d.Vector3{X: -0.2, Y: 0.2, Z: 2.2},
		Radius:   0.1})

	myScene.AddLight(lighting.PointLight{Intensity: image.Color{R: 180}})
	render := myScene.TraceScene(700, 600)
	render.Save("test.ppm")
}
