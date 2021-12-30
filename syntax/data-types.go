package syntax

import "fmt"

func DemoInitType() {
	var a int = 123
	b := 456
	var c, d int = 1, 2

	fmt.Println(a + b)
	fmt.Println(c + d)
}
