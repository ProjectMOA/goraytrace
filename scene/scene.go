package scene

import "github.com/ProjectMOA/goraytrace/shape"

// Scene defines a 3D scene that holds volumetric shapes
type Scene struct {
	spheres []*shape.Sphere
}

// MakeScene creates a new empty scene
func MakeScene() *Scene {
	return &Scene{spheres: make([]*shape.Sphere, 0, 10)}
}

// Elements returns the number of elements in the scene
func (s *Scene) Elements() int {
	return len(s.spheres)
}
