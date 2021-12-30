package syntax

import (
	"fmt"
	"strings"
)

func funcUsedFuncAsArgument(text string, fn func(str string) string) string {
	return fn(text)
}
func funcReturnsFunc(fn func(str string) string) func(str string) string {
	return fn
}

func funcWithClosure() func() int {
	callCounter := 0
	return func() int {
		callCounter++
		return callCounter
	}
}

func DemoFunctionsExtended() {
	transformedText := funcUsedFuncAsArgument("lowercase text", func(str string) string {
		return strings.ToUpper(str)
	})
	fmt.Println(transformedText)
	transformedText2 := funcReturnsFunc(func(str string) string {
		return strings.ToLower(str)
	})(transformedText)
	fmt.Println(transformedText2)

	callCounter := funcWithClosure()
	fmt.Println(callCounter())
	fmt.Println(callCounter())
	fmt.Println(callCounter())
}
