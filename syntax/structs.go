package syntax

import "fmt"

type DemoNestedStruct struct {
	Z int
}
type DemoStruct struct {
	z DemoNestedStruct
	X int
	Y int
}

func DemoStructs() {
	demo := DemoStruct{
		DemoNestedStruct{
			Z: 3,
		},
		1,
		2,
	}
	demo2 := DemoStruct{
		z: DemoNestedStruct{
			Z: 3,
		},
		X: 1,
		Y: 2,
	}
	fmt.Println(demo.X)
	fmt.Println(demo.Y)
	fmt.Println(demo.z.Z)
	demoPointer := &demo2

	fmt.Println(demo2.X)
	fmt.Println(demo2.Y)
	fmt.Println(demo2.z.Z)
	demoPointer.z.Z = 4
	fmt.Println(demo2.z.Z)
}
