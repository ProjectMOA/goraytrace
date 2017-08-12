package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ProjectMOA/goraytrace/scene"
)

func paniciferr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Setting up a scene
	if len(os.Args) < 2 {
		fmt.Println("Need a scene file as a parameter!")
		os.Exit(1)
	}

	myScene := scene.LoadSceneFile(os.Args[1])
	RenderScene(myScene, "main", true)
}

// RenderScene renders the scene passed as a parameter and saves the image
// with the name
func RenderScene(aScene *scene.Scene, name string, showTime bool) {
	start := time.Now()
	render := aScene.TraceScene(1000, 1000)
	elapsed := time.Since(start)
	if showTime {
		fmt.Printf("Rendered in: %s", elapsed)
	}

	render.Save(name + ".ppm")
}
