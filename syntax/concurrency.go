package syntax

import (
	"fmt"
	"sync"
	"time"
)

func numberGenerator(c *chan int) {
	for i := 0; i < 5; i++ {
		*c <- i
		time.Sleep(time.Second)
	}
}

func ConcurrencyDemo() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	c := make(chan int)

	go func(c *chan int) {
		numberGenerator(c)
		close(*c)
		wg.Done()
	}(&c)

	go func(c *chan int) {
		for v := range *c {
			fmt.Printf("got value %d\n", v)
		}
	}(&c)

	wg.Wait()
}
