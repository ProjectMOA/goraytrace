package maputil

import "fmt"

// ToSliceOfMap returns a slice of maps with the information in the original
// slice
func ToSliceOfMap(slice []interface{}) []map[string]interface{} {
	retval := make([]map[string]interface{}, 0, len(slice))
	for _, v := range slice {
		retval = append(retval, v.(map[string]interface{}))
	}
	return retval
}

// ToMapOfFloat64 returns a map[string]float64 from the original map.
// This will panic if any value in m is not a float.
func ToMapOfFloat64(m map[string]interface{}) map[string]float64 {
	var ok bool
	retval := make(map[string]float64)
	for k, v := range m {
		retval[k], ok = v.(float64)
		if !ok {
			panic(fmt.Sprint(v) + " is not a float")
		}
	}
	return retval
}
