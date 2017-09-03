package camera

import "github.com/ProjectMOA/goraytrace/math3d"

// TracingTargetIterator defines an iterator to get all points that must be traced.
// It's initialized from a Camera, a width and a height.
type TracingTargetIterator struct {
	width, height, currx, curry int
	firstPoint                  math3d.Vector3
	pxsize                      float64
}

// HasNext returns true if there are some points left to trace.
func (tti *TracingTargetIterator) HasNext() bool {
	return tti.curry != tti.height
}

// Next returns the next point that must be traced.
func (tti *TracingTargetIterator) Next() (*math3d.Vector3, int, int) {
	retVal := tti.firstPoint.Add(math3d.UnitX.Multiply(tti.pxsize).Multiply(float64(tti.currx)).Add(math3d.UnitY.Multiply(tti.pxsize).Multiply(float64(tti.curry))))
	rx, ry := tti.currx, tti.curry
	tti.currx++
	if tti.currx == tti.width {
		tti.curry++
		tti.currx = 0
	}
	return retVal, rx, ry
}
