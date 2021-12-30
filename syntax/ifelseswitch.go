package syntax

import (
	"fmt"
	"math"
)

func DemoIf(a int) {
	if a > 4 {
		fmt.Printf("%d is greater than 4\n", a)
	} else if a == 4 {
		fmt.Printf("%d is equal 4\n", a)
	} else {
		fmt.Printf("%d is not greater than 4\n", a)
	}
}

func DemoIfShortStatement(a float64, b float64) float64 {
	if v := math.Pow(a, 2); v < b {
		return v
	}
	return 0
}

func DemoSwitchCase(a float64) {
	switch pow := math.Pow(a, 2); pow {
	case 4:
		fmt.Println("a = 2")
	case 9:
		fmt.Println("a = 3")
	default:
		fmt.Printf("a = %f\n", pow)
	}
}

func DemoSwitchNoCondition(a float64) {
	switch {
	case a > 2:
		fmt.Println("a > 2")
	case a < 2:
		fmt.Println("a < 3")
	default:
		fmt.Println("a == 2")
	}
}
