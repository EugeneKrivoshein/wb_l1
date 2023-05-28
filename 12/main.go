package main

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, item := range sequence {
		set[item] = true
	}

	// Создаем срез для хранения ключей множества
	keys := make([]string, 0, len(set))
	for item := range set {
		keys = append(keys, item)
	}

	// Выводим множество
	fmt.Printf("subset: %v\n", keys)
}
