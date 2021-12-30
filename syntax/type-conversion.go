package syntax

import (
	"fmt"
	"strconv"
)

func DemoTypeConversion() {
	var a int = 1
	var b float32 = float32(a)
	fmt.Println(b)

	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	if err != nil {
		fmt.Println("String is not number")
		return
	}
	fmt.Println(intVar)
}
