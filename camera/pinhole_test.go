package camera

import (
	"fmt"
	"testing"

	"github.com/ProjectMOA/goraytrace/math3d"
)

func TestPointIterator(t *testing.T) {
	camera := PinHole{
		FocalPoint:        math3d.Vector3{},
		FoV:               0.3490659,
		Right:             math3d.UnitX,
		Up:                math3d.UnitY,
		Towards:           math3d.UnitZ,
		ViewPlaneDistance: 1.0}

	iterator := camera.GetIterator(10, 10)
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
