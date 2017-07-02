package camera

import (
	"math"

	"github.com/ProjectMOA/goraytrace/math3d"
)

// PinHole defines a pinhole camera in 3D space.
type PinHole struct {
	Up, Right, Towards, FocalPoint math3d.Vector3
	FoV, ViewPlaneDistance         float64
}

// DefaultPinHole returns a default PinHole camera
func DefaultPinHole() PinHole {
	return PinHole{
		FocalPoint:        *math3d.UnitZ.Multiply(-1.5),
		FoV:               0.3490659,
		Right:             math3d.UnitX,
		Up:                math3d.UnitY,
		Towards:           math3d.UnitZ,
		ViewPlaneDistance: 1.0}
}

// GetIterator returns an iterator for the points that must be traced
// to render an image from a PinHole camera.
func (ph *PinHole) GetIterator(width, height int32) *TracingTargetIterator {
	middlePoint := ph.FocalPoint.Add(ph.Towards.Multiply(ph.ViewPlaneDistance))
	pixelSize := float64((2.0 * math.Tan(ph.FoV/2.0)) / float64(height))
	firstPoint := middlePoint.
		Subtract(ph.Right.Multiply(float64(width-1) / 2.0 * pixelSize).
			Add(ph.Up.Multiply(float64(height-1) / 2.0 * pixelSize)))
	return &TracingTargetIterator{currx: 0, curry: 0, firstPoint: *firstPoint, height: height, width: width, pxsize: pixelSize}
}
