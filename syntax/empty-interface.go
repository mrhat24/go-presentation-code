package syntax

import "fmt"

func DemoEmptyInterface() {
	// пустой интерфейс может хранить в себе любой тип данных, аналог any в TS
	// но так делать не лучшая практика
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
