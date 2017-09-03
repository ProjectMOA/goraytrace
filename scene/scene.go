package scene

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"

	"github.com/ProjectMOA/goraytrace/camera"
	"github.com/ProjectMOA/goraytrace/image"
	"github.com/ProjectMOA/goraytrace/lighting"
	"github.com/ProjectMOA/goraytrace/maputil"
	"github.com/ProjectMOA/goraytrace/math3d"
	"github.com/ProjectMOA/goraytrace/shape"
)

// Scene defines a 3D scene that holds volumetric shapes
type Scene struct {
	Camera camera.PinHole        `json:"camera"`
	Shapes []shape.Shape         `json:"shapes"`
	Lights []lighting.PointLight `json:"lights"`
}

// New creates a new empty scene with a default pinhole camera
func New() *Scene {
	return &Scene{Camera: camera.DefaultPinHole(), Shapes: make([]shape.Shape, 0, 10)}
}

// Elements returns the number of elements in the scene
func (s *Scene) Elements() int {
	return len(s.Shapes)
}

// AddShape adds a shape to the scene.
func (s *Scene) AddShape(aShape shape.Shape) {
	s.Shapes = append(s.Shapes, aShape)
}

// AddLight adds a light to the scene.
func (s *Scene) AddLight(aLightsource lighting.PointLight) {
	s.Lights = append(s.Lights, aLightsource)
}

// TraceScene traces the scene as it currently is, returning
// the final image.
func (s *Scene) TraceScene(width, height int) *image.Image {
	targetIt := s.Camera.GetIterator(width, height)
	var x, y int
	var point *math3d.Vector3
	render := image.New(width, height)
	for targetIt.HasNext() {
		point, x, y = targetIt.Next()
		s.traceRay(point, x, y, render)
	}

	return render
}

func (s *Scene) traceRay(p *math3d.Vector3, x int, y int, img *image.Image) {
	// Construct the light ray
	lr := &math3d.LightRay{Direction: *p.Subtract(&s.Camera.FocalPoint).Normalized(), Source: *p}
	// Check intersections with the shapes in the scene
	nearestDistance, nearestShape := s.getNearestIntersection(lr)

	if nearestDistance != math.MaxFloat64 {
		// The lightray intersected a shape
		intersection := lr.Source.Add(lr.Direction.Multiply(nearestDistance))
		// Calculate the radiance at the intersection
		radiance := s.calculateRadianceAt(intersection, lr, nearestShape)
		img.Set(x, y, radiance.ToNRGBA())
	} else {
		// The lightray didn't intersect any shape. Just fill the pixel in black
		img.Set(x, y, image.Black.ToNRGBA())
	}
}

func (s *Scene) calculateRadianceAt(intersection *math3d.Vector3, incidentalRay *math3d.LightRay, sh shape.Shape) image.Color {
	// trace shadow rays towards all light sources
	radiance := image.Color{}
	for _, ls := range s.Lights {
		pointToLightVector := ls.Position.Subtract(intersection)
		shadowRay := math3d.LightRay{Direction: *pointToLightVector.Normalized(), Source: *intersection}
		if !s.inShadow(&shadowRay, pointToLightVector.Abs()) {
			normal := sh.NormalAt(&ls.Position).Normalized()
			// Cosine of the ray of light with the visible normal.
			cosine := shadowRay.Direction.Dot(normal)
			if cosine > 0.0 {
				reflected := incidentalRay.Direction.Multiply(-1).Reflect(normal)
				rCosine := math3d.Clamp(incidentalRay.Direction.Dot(reflected), 0, 1)
				shiny := 0.0
				// Prepared to use phong materials
				phong := image.White.Divide(math.Pi).Add(image.Black.Multiply((shiny + 2) / (2 * math.Pi) * math.Pow(rCosine, shiny)))
				radiance = *radiance.Add(ls.Intensity.CMultiply(phong).Multiply(cosine))
			}
		}
	}
	return radiance
}

func (s *Scene) getNearestIntersection(lr *math3d.LightRay) (float64, shape.Shape) {
	var nearestShape shape.Shape
	nearestDistance := math.MaxFloat64
	for _, s := range s.Shapes {
		intersectionDistance := s.Intersect(lr)
		if intersectionDistance < nearestDistance {
			nearestDistance = intersectionDistance
			nearestShape = s
		}
	}
	return nearestDistance, nearestShape
}

// inShadow returns true if the lightray intersects any shape
// at a distance that is smaller than distance
func (s *Scene) inShadow(lr *math3d.LightRay, distance float64) bool {
	for _, s := range s.Shapes {
		intersectionDistance := s.Intersect(lr)
		if intersectionDistance < distance {
			return true
		}
	}
	return false
}

// SaveSceneFile saves the scene as a file that can be loaded later
func (s *Scene) SaveSceneFile(path string) {
	marshaledScene, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	var mappedScene map[string]interface{}
	err = json.Unmarshal(marshaledScene, &mappedScene)
	if err != nil {
		panic(err)
	}
	mappedScene["shapes"] = shape.AsMap(s.Shapes)
	marshaledScene, err = json.MarshalIndent(mappedScene, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshaledScene))

	err = ioutil.WriteFile(path, marshaledScene, 0644)
	if err != nil {
		panic(err)
	}
}

// LoadSceneFile loads a scene file to a scene object
func LoadSceneFile(path string) *Scene {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var scenemap map[string]interface{}
	err = json.Unmarshal(bytes, &scenemap)
	if err != nil {
		panic(err)
	}
	retscene := &Scene{}
	retscene.Camera = camera.PinHoleFromMap(scenemap["camera"].(map[string]interface{}))
	retscene.Lights = lighting.PointLightsFromMap(maputil.ToSliceOfMap(scenemap["lights"].([]interface{})))
	retscene.Shapes = shape.FromMap(maputil.ToSliceOfMap(scenemap["shapes"].([]interface{})))
	return retscene
}
