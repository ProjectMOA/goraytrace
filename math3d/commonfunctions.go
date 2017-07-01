package math3d

// Clamp limits the value to the min and max specified
func Clamp(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
