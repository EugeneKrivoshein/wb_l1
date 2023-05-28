package main

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func calculateSquare(number int) int {
	return number * number
}

// Функция для запуска горутин и ожидания их завершения
func calculateSumOfSquares(numbers []int) {
	// Создаем канал для передачи результатов
	results := make(chan int)

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

	// Собираем результаты вычислений
	sum := 0
	for square := range results {
		sum += square
	}

	fmt.Println(sum)
}

// Функция для вычисления квадрата числа в асинхронном режиме
func calculateSquareAsync(number int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	square := calculateSquare(number)
	results <- square
}

func main() {
	// Исходный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Обработка чисел с использованием горутин
	calculateSumOfSquares(numbers)
}
