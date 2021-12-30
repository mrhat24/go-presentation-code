package syntax

// FunctionExample Функция с аргументами
func FunctionExample(a int, b int) int {
	return a + b
}

// FunctionExample2 Функция с одинаковыми типами аргументов
func FunctionExample2(a, b int) int {
	return a + b
}

// FunctionExample3 Функция возвращающая несколько результатов
func FunctionExample3(a, b int) (int, int) {
	return a + b, a - b
}

// FunctionExample4 Функция с именованным результатом
func FunctionExample4(a, b int) (c, d int) {
	c = a + b
	d = a - b
	return
}
