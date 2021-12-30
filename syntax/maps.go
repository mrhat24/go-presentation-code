package syntax

import "fmt"

type Str struct {
	ass string
}

func DemoMaps() {
	var m0 = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m0)
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	m["two"] = 3
	fmt.Println(m)
}
