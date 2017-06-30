package shape

import "github.com/ProjectMOA/goraytrace/math"

// Sphere defines a spheric shape in 3D space.
type Sphere struct {
	Position math.Vector3
	Radius   float64
}
