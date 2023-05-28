package main

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()         // Захватываем мьютекс перед доступом к счетчику
	defer c.mu.Unlock() // Освобождаем мьютекс после завершения операции
	c.value++           // Инкрементируем счетчик
}

func (c *Counter) GetValue() int {
	c.mu.Lock()         // Захватываем мьютекс перед доступом к счетчику
	defer c.mu.Unlock() // Освобождаем мьютекс после завершения операции
	return c.value      // Возвращаем текущее значение счетчика
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	numGoroutines := 100

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			counter.Increment() // Инкрементируем счетчик
			wg.Done()           // Уменьшаем счетчик горутин при завершении работы
		}()
	}

	wg.Wait()                                                      // Ожидаем завершения всех горутин
	fmt.Println("Итоговое значение счетчика:", counter.GetValue()) // Выводим итоговое значение счетчика
}
