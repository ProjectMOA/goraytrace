package math3d

import (
	"fmt"
	"math"
)

const threshold float64 = 0.00001

// Vector3 holds three floats that represent X Y and Z space.
// It holds both 3D vectors and 3D points.
type Vector3 struct {
	X, Y, Z float64
}

// Abs returns the distance from the origin
func (v *Vector3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalized returns the normalized 3D vector
func (v *Vector3) Normalized() *Vector3 {
	return v.Divide(v.Abs())
}

// Divide returns a vector result of dividing all the values in
// the vector by k
func (v *Vector3) Divide(k float64) *Vector3 {
	return &Vector3{v.X / k, v.Y / k, v.Z / k}
}

// Multiply returns a vector result of multiplying all the values
// in the vector by k
func (v *Vector3) Multiply(k float64) *Vector3 {
	return &Vector3{v.X * k, v.Y * k, v.Z * k}
}

// Add returns the result of adding two vectors
func (v *Vector3) Add(v2 *Vector3) *Vector3 {
	return &Vector3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

// Subtract returns the result of subtracting two vectors
func (v *Vector3) Subtract(v2 *Vector3) *Vector3 {
	return &Vector3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

// Dot returns the dot product of the 3D vectors
func (v *Vector3) Dot(v2 *Vector3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Cross returns the cross product of the 3D vectors
func (v *Vector3) Cross(v2 *Vector3) *Vector3 {
	return &Vector3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X}
}

// Subtract returns the vector that goes from pointA to
// pointB.
func Subtract(pointA *Vector3, pointB *Vector3) *Vector3 {
	return &Vector3{X: pointA.X - pointB.X,
		Y: pointA.Y - pointB.Y,
		Z: pointA.Z - pointB.Z}
}

// Distance returns the distance from pointA to pointB.
func Distance(pointA *Vector3, pointB *Vector3) float64 {
	return Subtract(pointA, pointB).Abs()
}

// Equal returns true if both vectors are the same within a
// margin of error
func (v *Vector3) Equal(v2 *Vector3) bool {
	return math.Abs(v.X-v2.X) < threshold &&
		math.Abs(v.Y-v2.Y) < threshold &&
		math.Abs(v.Z-v2.Z) < threshold
}

// Differ returns true if the vectors are not the same within a
// margin of error.
func (v *Vector3) Differ(v2 *Vector3) bool {
	return !v.Equal(v2)
}

// LesserOrEqual returns true if the first vector is smaller or
// equal in the three axes.
func (v *Vector3) LesserOrEqual(v2 *Vector3) bool {
	return v.X-v2.X <= threshold &&
		v.Y-v2.Y <= threshold &&
		v.Z-v2.Z <= threshold
}

// GreaterOrEqual returns true if the first vector is greater or
// equal in the three axes.
func (v *Vector3) GreaterOrEqual(v2 *Vector3) bool {
	return v2.X-v.X <= threshold &&
		v2.Y-v.Y <= threshold &&
		v2.Z-v.Z <= threshold
}

func (v *Vector3) String() string {
	return fmt.Sprintf("[%.3f, %.3f, %.3f]", v.X, v.Y, v.Z)
}

// Print the values in the 3D vector
func (v *Vector3) Print() {
	fmt.Print(v.String())
}
