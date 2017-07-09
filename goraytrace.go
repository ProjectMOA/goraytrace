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

	start := time.Now()
	render := myScene.TraceScene(1000, 1000)
	elapsed := time.Since(start)
	fmt.Printf("Rendered in: %s", elapsed)

	render.Save("test.ppm")
}
