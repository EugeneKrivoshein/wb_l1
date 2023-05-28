package main

// Поменять местами два числа без создания временной переменной.

import "fmt"

func swapNumbers(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	x := 10
	y := 5
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)

	x, y = swapNumbers(x, y)

	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}
