package main

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(2<<22), big.NewInt(3<<43)

	// Проверка, что значения a и b больше 2^20
	twoTo20 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(20), nil)
	if a.Cmp(twoTo20) <= 0 || b.Cmp(twoTo20) <= 0 {
		fmt.Println("Значения переменных должны быть больше 2^20")
		return
	}

	// Умножение
	result := big.NewInt(0).Mul(a, b)
	fmt.Println("Результат умножения:", result.String())

	// Деление
	result = big.NewInt(0).Div(a, b)
	fmt.Println("Результат деления:", result.String())

	// Сложение
	result = big.NewInt(0).Add(a, b)
	fmt.Println("Результат сложения:", result.String())

	// Вычитание
	result = big.NewInt(0).Sub(a, b)
	fmt.Println("Результат вычитания:", result.String())
}
