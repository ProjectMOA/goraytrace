package scene

import (
	"math"

	"github.com/ProjectMOA/goraytrace/camera"
	"github.com/ProjectMOA/goraytrace/image"
	"github.com/ProjectMOA/goraytrace/lighting"
	"github.com/ProjectMOA/goraytrace/math3d"
	"github.com/ProjectMOA/goraytrace/shape"
)

// Scene defines a 3D scene that holds volumetric shapes
type Scene struct {
	Camera camera.PinHole
	shapes []shape.Shape
	lights []lighting.PointLight
}

// New creates a new empty scene with a default pinhole camera
func New() *Scene {
	return &Scene{Camera: camera.DefaultPinHole(), shapes: make([]shape.Shape, 0, 10)}
}

// Elements returns the number of elements in the scene
func (s *Scene) Elements() int {
	return len(s.shapes)
}

// AddShape adds a shape to the scene.
func (s *Scene) AddShape(aShape shape.Shape) {
	s.shapes = append(s.shapes, aShape)
}

// AddLight adds a light to the scene.
func (s *Scene) AddLight(aLightSource lighting.PointLight) {
	s.lights = append(s.lights, aLightSource)
}

// TraceScene traces the scene as it currently is, returning
// the final image.
func (s *Scene) TraceScene(width, height int32) *image.Image {
	targetIt := s.Camera.GetIterator(width, height)
	var x, y int32
	var point *math3d.Vector3
	render := image.New(width, height)
	for targetIt.HasNext() {
		point, x, y = targetIt.Next()
		lr := &math3d.LightRay{Direction: *point.Subtract(&s.Camera.FocalPoint).Normalized(), Source: *point}
		for _, s := range s.shapes {
			intersected := s.Intersect(lr) != math.MaxFloat64
			if intersected {
				render.Pixels[y][x] = image.Color{R: 200}
			}
		}
	}

	return render
}
