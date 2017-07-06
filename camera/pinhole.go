package camera

import (
	"math"

	"github.com/ProjectMOA/goraytrace/math3d"
)

// PinHole defines a pinhole camera in 3D space.
type PinHole struct {
	Up                math3d.Vector3 `json:"up"`
	Right             math3d.Vector3 `json:"right"`
	Towards           math3d.Vector3 `json:"towards"`
	FocalPoint        math3d.Vector3 `json:"focalpoint"`
	FoV               float64        `json:"fieldofview"`
	ViewPlaneDistance float64        `json:"viewplanedistance"`
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

// PinHoleFromMap returns the pinhole camera defined in the map
func PinHoleFromMap(m map[string]interface{}) PinHole {
	ph := PinHole{}
	ph.FocalPoint = math3d.VectorFromMap(m["focalpoint"].(map[string]interface{}))
	ph.FoV, _ = m["fieldofview"].(float64)
	ph.ViewPlaneDistance, _ = m["viewplanedistance"].(float64)
	ph.Up = math3d.VectorFromMap(m["up"].(map[string]interface{}))
	ph.Right = math3d.VectorFromMap(m["right"].(map[string]interface{}))
	ph.Towards = math3d.VectorFromMap(m["towards"].(map[string]interface{}))
	return ph
}
