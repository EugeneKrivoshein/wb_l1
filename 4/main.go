package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Worker считывает число из канала и выводит его
type Worker struct {
	ID   int
	In   <-chan int
	Done chan struct{}
}

// NewWorker создает новый экземпляр Worker
func NewWorker(id int, in <-chan int) *Worker {
	return &Worker{
		ID:   id,
		In:   in,
		Done: make(chan struct{}),
	}
}

// Start начинает работу воркера
func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case n := <-w.In:
			fmt.Printf("Worker %d: got %d\n", w.ID, n)
		case <-w.Done:
			fmt.Printf("Worker %d: exiting\n", w.ID)
			return
		}
	}
}

// Writer записывает случайные числа в канал с определенной периодичностью
type Writer struct {
	Out  chan<- int
	Done chan struct{}
}

// NewWriter создает новый экземпляр Writer
func NewWriter(out chan<- int) *Writer {
	return &Writer{
		Out:  out,
		Done: make(chan struct{}),
	}
}

// Start начинает работу писателя
func (w *Writer) Start(wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done()

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			n := rand.Int()
			w.Out <- n
			fmt.Printf("Writer sent: %d\n", n)
		case <-w.Done:
			fmt.Println("Writer: exiting")
			return
		}
	}
}

// main является главной функцией программы
func main() {
	rand.Seed(time.Now().UnixNano())

	var workersNum int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workersNum)

	input := make(chan int, 2)
	defer close(input)

	_, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Запускаем воркеры
	workers := make([]*Worker, workersNum)
	for i := 0; i < workersNum; i++ {
		worker := NewWorker(i+1, input)
		workers[i] = worker
		wg.Add(1)
		go worker.Start(wg)
	}

	// Запускаем писателя
	writer := NewWriter(input)
	wg.Add(1)
	go writer.Start(wg, 1*time.Second)

	// Перехватываем сигналы завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\nПолучен сигнал прерывания")

	// Останавливаем воркеры и писателя
	cancel()
	close(writer.Done)
	for _, worker := range workers {
		close(worker.Done)
	}

	wg.Wait()
	fmt.Println("Программа завершена")
}
