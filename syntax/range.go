package syntax

import (
	"fmt"
	"math"
)

func DemoRange() {
	array := []int{1, 2, 3, 4, 5}
	for i, v := range array {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	array2 := make([]float64, 10)
	for i := range array2 {
		array2[i] = math.Pow(float64(i), float64(i))
	}
	for _, v := range array2 {
		fmt.Printf("value: %f\n", v)
	}
}
