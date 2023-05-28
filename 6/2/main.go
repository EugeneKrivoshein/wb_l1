package main

// Реализовать все возможные способы остановки выполнения горутины.

// Использование канала для остановки:
// Вы можете использовать канал для передачи сигнала остановки горутине.

import (
	"fmt"
	"time"
)

func myGoroutine(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Goroutine stopped")
			return
		default:
			// Выполнение работы горутины
			fmt.Println("Doing some work...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stopChan := make(chan struct{})
	go myGoroutine(stopChan)

	time.Sleep(5 * time.Second)
	close(stopChan)

	time.Sleep(2 * time.Second)
}
