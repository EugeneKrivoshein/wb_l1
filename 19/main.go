package main

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(str string) string {
	runes := []rune(str)                  // Преобразуем строку в срез рун
	length := utf8.RuneCountInString(str) // Получаем длину строки в символах Unicode

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 { // Переворачиваем срез рун
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) // Возвращаем перевернутую строку
}

func main() {
	str := "главрыба"
	reversed := reverseString(str)
	fmt.Println(str)
	fmt.Println(reversed)
}
