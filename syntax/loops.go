package syntax

import (
	"fmt"
	"time"
)

func DemoLoops() {

	// Стандартный цикл
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while
	sum := 0
	for sum < 10 {
		sum += 1
		fmt.Println(sum)
	}

	// все части могут быть опциональны
	sum2 := 0
	for ; sum2 < 10; sum2++ {
		fmt.Println(sum2)
	}

	for {
		fmt.Println("I am forever loop!")
		time.Sleep(time.Second)
	}

}
