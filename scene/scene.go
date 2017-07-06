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
func (s *Scene) TraceScene(width, height int32) *image.Image {
	targetIt := s.Camera.GetIterator(width, height)
	var x, y int32
	var point *math3d.Vector3
	render := image.New(width, height)
	for targetIt.HasNext() {
		point, x, y = targetIt.Next()
		s.traceRay(point, x, y, render)
	}

	return render
}

func (s *Scene) traceRay(p *math3d.Vector3, x int32, y int32, img *image.Image) {
	lr := &math3d.LightRay{Direction: *p.Subtract(&s.Camera.FocalPoint).Normalized(), Source: *p}
	for _, s := range s.Shapes {
		intersected := s.Intersect(lr) != math.MaxFloat64
		if intersected {
			img.Pixels[y][x] = image.Color{R: 200}
		}
	}
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
