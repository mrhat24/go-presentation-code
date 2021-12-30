package tests

import "testing"

func TestTestExampleFunc(t *testing.T) {
	result := TestExampleFunc()
	if result != 2 {
		t.Error("result is not equal 2")
	}
}
