package main

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Split(input, " ")          // разделяет входную строку на слова
	reversedWords := make([]string, len(words)) // создается новый срез

	for i := 0; i < len(words); i++ {
		reversedWords[i] = words[len(words)-1-i]
	}

	return strings.Join(reversedWords, " ") // объединяет перевернутые слова
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
