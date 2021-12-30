package syntax

import (
	"fmt"
	"strings"
)

type UpperCaseAbleStruct struct {
	name string
}

type UpperCaseAbleType string

type UpperCaseAble interface {
	getUpperName() string
}

func (e UpperCaseAbleStruct) getUpperName() string {
	return strings.ToUpper(e.name)
}

func (e *UpperCaseAbleStruct) setUpperName() {
	e.name = strings.ToUpper(e.name)
}

func (e UpperCaseAbleType) getUpperName() string {
	return strings.ToUpper(string(e))
}

func PrintUpperCase(e UpperCaseAble) {
	fmt.Println(e.getUpperName())
}

func DemoMethods() {
	e := &UpperCaseAbleStruct{
		name: "some name",
	}
	var s UpperCaseAbleType = "some text"
	PrintUpperCase(*e)
	PrintUpperCase(s)
	fmt.Println(e.name)
	e.setUpperName()
	fmt.Println(e.name)
}
