package shape

import (
	"github.com/ProjectMOA/goraytrace/math3d"
)

// Shape defines the methods shared by all 3D shapes
type Shape interface {
	Intersect(lr *math3d.LightRay) float64
	AsMap() map[string]interface{}
}

// AsMap turns the input slice of shapes to a slice of maps that can be
// serialized.
func AsMap(shapes []Shape) []map[string]interface{} {
	retval := make([]map[string]interface{}, 0, len(shapes))
	for _, s := range shapes {
		retval = append(retval, s.AsMap())
	}
	return retval
}

// FromMap returns a slice of shapes made from the slice of map
func FromMap(themap []map[string]interface{}) []Shape {
	shapes := make([]Shape, 0, len(themap))
	for _, m := range themap {
		switch m["type"] {
		case "sphere":
			shapes = append(shapes, SphereFromMap(m))
		default:
			panic("That shape is not implemented yet or the type field is empty")
		}
	}
	return shapes
}
