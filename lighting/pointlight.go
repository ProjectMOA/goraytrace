package lighting

import (
	"github.com/ProjectMOA/goraytrace/image"
	"github.com/ProjectMOA/goraytrace/math3d"
)

// PointLight defines a punctual light in 3D space
type PointLight struct {
	Position  math3d.Vector3
	Intensity image.Color
}
