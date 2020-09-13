package misc

import (
	"math"
	"testing"
)

func TestHomer(t *testing.T) {
	exp := []int{1, 2, 3}
	sum := 0
	xVal := 2
	for k, v := range exp {
		sum += int(math.Pow(float64(xVal), float64(k))) * v
	}
	res := Horner(exp, xVal)
	if res != sum {
		t.Errorf("Horner(%v), expected: %v, got: %v", exp, sum, res)
	}
}
