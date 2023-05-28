package main

// Реализовать все возможные способы остановки выполнения горутины.

// Использование флага для остановки:
// Вы можете использовать булевый флаг для управления выполнением горутины.
// Установите флаг в true для остановки горутины.
import (
	"fmt"
	"time"
)

func myGoroutine(stopChan chan bool) {
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
	stopChan := make(chan bool)
	go myGoroutine(stopChan)

	time.Sleep(5 * time.Second)
	stopChan <- true

	time.Sleep(2 * time.Second)
}
