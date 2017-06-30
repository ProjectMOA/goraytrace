package math3d

import "math"

// DiscardIfTooClose returns t if t is greater than a threshold.
// Otherwise, it returns math.MaxFloat64
func DiscardIfTooClose(t float64) float64 {
	if t > threshold {
		return t
	}
	return math.MaxFloat64
}

// GetNearestInFront returns the minimum of t1 and t2 so long as it is greater
// than a threshold. If neither is, it returns math.MaxFloat64
func GetNearestInFront(t1 float64, t2 float64) float64 {
	/*
	 * Intersection Point 1 is in front of the camera and
	 * (before Point 2 or it is the only in front of the camera).
	 */
	if t1 > threshold && (t1 < t2 || t2 <= threshold) {
		return t1
	} else if t2 > threshold && (t2 < t1 || t1 <= threshold) {
		//The intersection Point 2 is in front of the camera and
		//(before Point 1 or it is the only in front of the camera).
		return t2
	}
	// Both intersection points are behind the camera.
	return math.MaxFloat64
}
