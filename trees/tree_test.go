package trees

import (
	"fmt"
	"math"
	"testing"
)

func TestFromSlice(t *testing.T) {
	arr := []int{0, 1, 2, math.MinInt64, 3, 4}
	node := FromSlice(arr)
	fmt.Println(node)
}
