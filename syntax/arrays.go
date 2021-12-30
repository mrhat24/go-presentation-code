package syntax

import "fmt"

func DemoArrays() {
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Printf("%s %s\n", array[0], array[1])
	fmt.Println("array", array)

	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	array2Slice := array2[4:8]
	fmt.Println("array2Slice", array2Slice)
	array2[4] = 123
	array2[5] = 123
	fmt.Println("array2Slice", array2Slice)
	array2Slice2 := array2[:8]
	fmt.Println("array2Slice2", array2Slice2)
	array2Slice3 := array2[5:]
	fmt.Println("array2Slice3", array2Slice3)

	array3 := make([]int, 0, 5)
	fmt.Println("array3", array3)

	array4 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println("array4", array4)

	array5 := []int{1, 2, 3}
	array5 = append(array5, 4)
	fmt.Println("array5", array5)
}
