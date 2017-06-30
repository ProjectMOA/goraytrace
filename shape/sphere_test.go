package shape

import (
	"testing"

	"github.com/ProjectMOA/goraytrace/math3d"
)

func TestRaySphereIntersection(t *testing.T) {
	mySphere := Sphere{
		Position: math3d.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Radius:   1.0}

	myLightRay := math3d.LightRay{
		Direction: math3d.Vector3{X: 0.0, Y: 0.0, Z: 1.0},
		Source:    math3d.Vector3{X: 0.0, Y: 0.0, Z: -2.0}}

	intersectionDistance := mySphere.Intersect(&myLightRay)
	if intersectionDistance != 1.0 {
		t.Errorf("The lightray should intersect the sphere at D=1.0 but it intersects at %.3f", intersectionDistance)
	}
}
