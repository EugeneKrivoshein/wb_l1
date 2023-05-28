package main

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

import (
	"fmt"
	"sync"
)

func main() {
	// Инициализация конвейера
	inCh := make(chan int)
	outCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go stage1(inCh, &wg)

	wg.Add(1)
	go stage2(inCh, outCh, &wg)

	go output(outCh)

	wg.Wait()
}

// Запись чисел в первый канал
func stage1(inCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	input := []int{1, 2, 3, 4, 5}
	for _, num := range input {
		inCh <- num
	}
	close(inCh)
}

// Умножение чисел и запись во второй канал
func stage2(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range inCh {
		result := num * 2
		outCh <- result
	}
	close(outCh)
}

// Вывод результатов в stdout
func output(outCh <-chan int) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for result := range outCh {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
