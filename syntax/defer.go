package syntax

import "fmt"

func DemoDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
