package main

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.

import (
	"fmt"
	"math/big"
	"sync"
)

type BitNumber struct {
	number *big.Int
	mu     sync.Mutex
}

func (bn *BitNumber) changeBit(index int) {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	bit := new(big.Int)             // новый экземпляр объекта
	bit.Rsh(bn.number, uint(index)) //сдвиг вправо
	bit.And(bit, big.NewInt(1))     // выполняет побитовое И между результатом сдвига и числом 1

	// Если bit равно 0, устанавливаем i-й бит в 1
	// Если bit равно 1, устанавливаем i-й бит в 0
	if bit.Cmp(big.NewInt(0)) == 0 { //сравнениние чисел
		bn.number.SetBit(bn.number, index, 1) // устанавливает определенный бит числа
	} else {
		bn.number.SetBit(bn.number, index, 0)
	}
}

// возвращает копию числа
func (bn *BitNumber) getNumber() *big.Int {
	bn.mu.Lock()
	defer bn.mu.Unlock()

	return new(big.Int).Set(bn.number)
}

func main() {
	var number, index int64
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	for {
		fmt.Print("Введите индекс бита: ")
		_, err := fmt.Scan(&index)
		if err != nil {
			fmt.Println("Ошибка при чтении индекса бита:", err)
			return
		}

		// Проверяем, что индекс находится в допустимом диапазоне [0, 63]
		if index >= 0 && index <= 63 {
			break
		}

		fmt.Println("Недопустимый индекс бита!")
	}
	//создает новый экземпляр структуры
	bn := &BitNumber{
		number: big.NewInt(number),
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		bn.changeBit(int(index))
	}()

	wg.Wait()

	result := bn.getNumber()
	fmt.Println("Результат:", result)
}
