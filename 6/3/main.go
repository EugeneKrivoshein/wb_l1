package main

// Реализовать все возможные способы остановки выполнения горутины.

// Использование контекста для остановки:
// Вы можете использовать контекст для остановки горутины.
// Вызовите метод cancel() контекста, чтобы остановить выполнение горутины.
import (
	"context"
	"fmt"
	"time"
)

func myGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
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
	ctx, cancel := context.WithCancel(context.Background())
	go myGoroutine(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
}
