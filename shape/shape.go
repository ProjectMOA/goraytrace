package shape

import (
	"github.com/ProjectMOA/goraytrace/math3d"
)

// Shape defines the methods shared by all 3D shapes
type Shape interface {
	Intersect(lr *math3d.LightRay) float64
}
