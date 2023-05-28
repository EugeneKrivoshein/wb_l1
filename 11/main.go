package main

// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

func main() {
	set1 := []int{1, 2, 2, 3, 4}
	set2 := []int{2, 3, 3, 4, 5}

	intersection := make([]int, 0) // Создаем пустой срез для хранения пересечения множеств

	set1Map := make(map[int]bool) // Создаем хэш-таблицу для множества set1
	for _, num := range set1 {
		set1Map[num] = true // Заполняем хэш-таблицу значениями из set1
	}

	set2Map := make(map[int]int) // Создаем хэш-таблицу для множества set2(подсчет повторений элементов из второго множества set2)
	for _, num := range set2 {
		set2Map[num]++ // Увеличиваем счетчик для каждого элемента set2 в хэш-таблице set2Map
	}

	for num := range set1Map { // Перебираем элементы из set1Map
		count := set2Map[num] // Получаем количество вхождений элемента в set2
		if count > 0 {        // Если количество вхождений больше 0, значит элемент входит в пересечение
			intersection = append(intersection, num) // Добавляем элемент в пересечение
			set2Map[num]--                           // Уменьшаем счетчик в хэш-таблице set2Map
		}
	}

	fmt.Println(intersection) // Выводим пересечение множеств
}
