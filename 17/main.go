package main

// Реализовать бинарный поиск встроенными методами языка.

import "fmt"

func binarySearch(arr []int, query int) int {
	minIndex := 0
	maxIndex := len(arr) - 1

	for minIndex <= maxIndex {
		midIndex := (maxIndex + minIndex) / 2
		midItem := arr[midIndex]

		switch {
		case query == midItem:
			return midIndex
		case midItem < query:
			minIndex = midIndex + 1
		case midItem > query:
			maxIndex = midIndex - 1
		}
	}

	return -1
}

func main() {
	// Пример использования бинарного поиска
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	query := 6

	index := binarySearch(arr, query)
	if index != -1 {
		fmt.Printf("Элемент %d найден в индексе %d\n", query, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", query)
	}
}

// Алгоритм бинарного поиска работает на основе принципа "разделяй и властвуй", позволяя эффективно искать элементы в
// отсортированном массиве. За счет деления массива на половины на каждой итерации, бинарный поиск обычно имеет временную
// сложность O(log n).
