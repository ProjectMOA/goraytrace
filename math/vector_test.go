package math

import (
	"testing"
)

func TestVector3CreationAndPrint(t *testing.T) {
	v := Vector3{X: 1.0, Y: 2.0, Z: 3.0}
	v.Print()
}

func TestVector3Abs(t *testing.T) {
	v := Vector3{X: 2.0, Y: 2.0, Z: 1.0}
	if v.Abs() != 3 {
		t.Error("Wrong absolute value")
	}
}

func TestEqualityCheck(t *testing.T) {
	v1 := Vector3{X: 2.0, Y: 2.0, Z: 1.0}
	v2 := Vector3{X: 2.0, Y: 2.0, Z: 1.0}
	if !v1.Equal(&v2) {
		t.Error("Both vectors should be equal")
	}
}

func TestVectorDivision(t *testing.T) {
	v := (&Vector3{X: 2.0, Y: 2.0, Z: 1.0}).Divide(2.0)
	if !v.Equal(&Vector3{1.0, 1.0, 1.0}) {
		t.Error("Division went wrong")
	}
}
