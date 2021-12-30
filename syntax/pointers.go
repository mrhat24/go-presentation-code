package syntax

import "fmt"

func DemoPointers() {
	a := 2
	b := &a
	c := a
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", *b)
	fmt.Printf("c = %d\n", c)
	*b = 3
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", *b)
	fmt.Printf("c = %d\n", c)
	a = 4
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", *b)
	fmt.Printf("c = %d\n", c)
}
