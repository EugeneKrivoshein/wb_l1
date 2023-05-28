package main

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

import (
	"fmt"
	"strings"
)

func isUniqueString(input string) bool {
	// Преобразование строки к нижнему регистру
	lowerInput := strings.ToLower(input)

	// Создание карты для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Проверка каждого символа в строке
	for _, char := range lowerInput {
		// Если символ уже присутствует в карте, значит он не уникален
		if charMap[char] {
			return false
		}
		// Добавление символа в карту
		charMap[char] = true
	}

	return true
}

func main() {
	// Примеры вызова функции isUniqueString
	fmt.Println(isUniqueString("abcd"))      // true
	fmt.Println(isUniqueString("abCdefAaf")) // false
	fmt.Println(isUniqueString("aabcd"))     // false
}
