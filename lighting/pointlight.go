package lighting

import (
	"github.com/ProjectMOA/goraytrace/image"
	"github.com/ProjectMOA/goraytrace/maputil"
	"github.com/ProjectMOA/goraytrace/math3d"
)

// PointLight defines a punctual light in 3D space
type PointLight struct {
	Position  math3d.Vector3 `json:"position"`
	Intensity image.Color    `json:"intensity"`
}

// PointLightsFromMap returns the point light defined in the map
func PointLightsFromMap(m []map[string]interface{}) []PointLight {
	retval := make([]PointLight, 0, len(m))
	for _, v := range m {
		pl := PointLight{}
		pl.Position = math3d.VectorFromMap(v["position"].(map[string]interface{}))
		pl.Intensity = image.ColorFromMap(maputil.ToMapOfFloat64(v["intensity"].(map[string]interface{})))
		retval = append(retval, pl)
	}

	return retval
}
