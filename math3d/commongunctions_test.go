package math3d

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestClamp(t *testing.T) {
	min := 0.3
	max := 0.6
	for i := 0; i < 100; i++ {
		t := Clamp(rand.Float64(), min, max)
		fmt.Println(t)
		if t < min || t > max {
			panic("Clamp error!")
		}
	}
}
