package shape

import (
	"math"

	"github.com/ProjectMOA/goraytrace/math3d"
)

// Sphere defines a spheric shape in 3D space.
type Sphere struct {
	Position math3d.Vector3 `json:"position"`
	Radius   float64        `json:"radius"`
}

// Intersect returns the distance at which the lightray intersects
// the sphere
func (s *Sphere) Intersect(lr *math3d.LightRay) float64 {
	v := lr.Source.Subtract(&s.Position)
	a := lr.Direction.Dot(&lr.Direction)
	b := 2 * lr.Direction.Dot(v)
	c := v.Dot(v) - s.Radius
	bb4ac := b*b - 4*a*c
	if bb4ac < 0 {
		// The lightray misses the sphere
		return math.MaxFloat64
	} else if bb4ac > 0 {
		// The lightray intersects the sphere in two points
		t1 := (-b - math.Sqrt(bb4ac)) / (2 * a)
		t2 := (-b + math.Sqrt(bb4ac)) / (2 * a)
		return math3d.GetNearestInFront(t1, t2)
	}

	// The lightray is tangent to the sphere
	return math3d.DiscardIfTooClose(-b / (2 * a))
}

// NormalAt returns the normal vector of a point of the sphere.
// point must be a point in the surface of the sphere.
func (s *Sphere) NormalAt(point *math3d.Vector3) *math3d.Vector3 {
	return s.Position.Subtract(point).Divide(s.Radius)
}

// AsMap returns a map representation of this shape
func (s *Sphere) AsMap() map[string]interface{} {
	return map[string]interface{}{"type": "sphere", "position": s.Position.AsMap(), "radius": s.Radius}
}

// SphereFromMap returns a sphere with the values in the map
func SphereFromMap(themap map[string]interface{}) *Sphere {
	retval := &Sphere{}
	var ok bool
	retval.Position = math3d.VectorFromMap(themap["position"].(map[string]interface{}))
	retval.Radius, ok = themap["radius"].(float64)
	if !ok {
		panic("The sphere's position was empty or isn't a valid float")
	}
	return retval
}
