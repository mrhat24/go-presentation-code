package syntax

import (
	"fmt"
	"math/rand"
)

func generateNumberGreaterThan1() (int, error) {
	v := rand.Int()
	if v > 1 {
		return 0, fmt.Errorf("error, %d is greater than 1", v)
	}
	return v, nil
}

func DemoErrors() {
	result, err := generateNumberGreaterThan1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got value %d\n", result)
}
