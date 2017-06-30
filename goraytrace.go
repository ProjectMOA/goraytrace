package main

import "github.com/ProjectMOA/goraytrace/scene"
import "fmt"

func main() {
	// Setting up a scene
	myScene := scene.MakeScene()
	fmt.Println(myScene.Elements())
}
