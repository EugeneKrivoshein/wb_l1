package main

// Удалить i-ый элемент из слайса.

import (
	"fmt"
)

func removeElement(slice []int, index int) []int { // slice[:index] - это срез элементов slice от начала до индекса index.
	return append(slice[:index], slice[index+1:]...) // slice[index+1:] - это срез элементов slice начиная с индекса index+1 и до конца среза.
} // append чтобы объединить два среза.

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Текущий срез:", slice)

	fmt.Print("Введите индекс элемента для удаления: ")
	var index int
	_, err := fmt.Scanln(&index)
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	if index < 0 || index >= len(slice) {
		fmt.Println("Неверный индекс!")
		return
	}

	slice = removeElement(slice, index)

	fmt.Println("Срез после удаления элемента:", slice)
}
