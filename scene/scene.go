package scene

import "github.com/ProjectMOA/goraytrace/shape"

// Scene defines a 3D scene that holds volumetric shapes
type Scene struct {
	spheres []*shape.Sphere
}

// New creates a new empty scene
func New() *Scene {
	return &Scene{spheres: make([]*shape.Sphere, 0, 10)}
}

// Elements returns the number of elements in the scene
func (s *Scene) Elements() int {
	return len(s.spheres)
}

// Add adds a sphere to the array of spheres in this scene.
// Eventually, when go makes more sense, spheres should be
// a list of shapes in some way so that all possible shapes
// are stored in a unique block of memory.
func (s *Scene) Add(sphere *shape.Sphere) {
	s.spheres = append(s.spheres, sphere)
}
