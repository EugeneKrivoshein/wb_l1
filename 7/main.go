package main

// Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	wg := sync.WaitGroup{}           // WaitGroup для ожидания
	resultChan := make(chan int, 20) // Буферизированный канал для хранения результатов

	wg.Add(2) // 2 горутины 2 итерации ожидания

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		for _, v := range sl {
			resultChan <- v // Отправляем результат в канал
		}
		wg.Done() // -1 к итерации ожидания
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		sl := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
		for _, v := range sl {
			resultChan <- v // Отправляем результат в канал
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()         // Ждем пока горутины закончат работу
		close(resultChan) // Закрываем канал после завершения горутин
	}()

	// Читаем и выводим результаты из канала
	for num := range resultChan {
		fmt.Printf("%v ", num)
	}

	fmt.Println("\nВсе данные обработаны.")
}
