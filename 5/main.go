package main

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= count; i++ {
		ch <- i
		fmt.Println("Отправлено:", i)
		time.Sleep(time.Second) // Задержка 1 секунда
	}
	close(ch)
}

func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		fmt.Println("Получено:", value)
	}
}

func main() {
	N := 5 // Время работы программы в секундах
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go sender(ch, 10, &wg) // Отправляем 10 значений в канал

	go receiver(ch, &wg) // Читаем значения из канала

	wg.Wait() // Ожидание завершения всех горутин

	timer := time.NewTimer(time.Duration(N) * time.Second) // Таймер на N секунд
	<-timer.C                                              // Ожидание истечения таймера

	fmt.Println("Программа завершена")
}
