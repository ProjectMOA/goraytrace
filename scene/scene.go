package scene

import "github.com/ProjectMOA/goraytrace/shape"

// Scene defines a 3D scene that holds volumetric shapes
type Scene struct {
	shapes []shape.Shape
}

// New creates a new empty scene
func New() *Scene {
	return &Scene{shapes: make([]shape.Shape, 0, 10)}
}

// Elements returns the number of elements in the scene
func (s *Scene) Elements() int {
	return len(s.shapes)
}

// Add adds a shape to the scene.
func (s *Scene) Add(aShape shape.Shape) {
	s.shapes = append(s.shapes, aShape)
}
