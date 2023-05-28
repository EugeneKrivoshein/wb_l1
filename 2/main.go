package main

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

// Функция для вычисления квадрата числа
func calculateSquare(number int) int {
	return number * number
}

// Функция для запуска горутин и ожидания их завершения
func processNumbers(numbers []int) {
	// Создаем канал для передачи результатов
	results := make(chan int, len(numbers))

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquareAsync(num, &wg, results)
	}

	// Закрываем канал с результатами после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Читаем результаты из канала и выводим их в stdout
	for result := range results {
		fmt.Println(result)
	}
}

// Функция для вычисления квадрата числа в асинхронном режиме
func calculateSquareAsync(number int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	result := calculateSquare(number)
	results <- result
}

func main() {
	// Исходный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Обработка чисел с использованием горутин
	processNumbers(numbers)
}
