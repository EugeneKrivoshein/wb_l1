package main

// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	select { // select для ожидания определенного времени или прерывания.
	case <-time.After(time.Duration(seconds) * time.Second): // создать канал, который будет отправлять сигнал через указанное количество времени
		return
	}
}

func main() {
	fmt.Println("Before sleep")
	sleep(5)
	fmt.Println("After sleep")
}
