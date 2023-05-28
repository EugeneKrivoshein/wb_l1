package main

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.
// var justString string
// func someFunc() {
// 	  v := createHugeString(1 << 10)
//    justString = v[:100]
// }
// func main() {
//    someFunc()
// }

import (
	"fmt"
	"math/rand"
	"time"
)

func createHugeString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var sb []rune
	for i := 0; i < n; i++ {
		// Генерируем случайный символ Unicode от 0 до 65535
		// и добавляем его к строке
		r := rune(rand.Intn(65536))
		sb = append(sb, r)
	}

	return string(sb)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	var justString []rune
	if len(v) >= 100 {
		justString = []rune(v[:100])
	} else {
		justString = []rune(v)
	}
	// Освобождаем память, выделенную для строки v
	v = ""
	return string(justString)
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}

// негативные последствия:
// 1) Потенциальное переполнение памяти:
// 2) Ошибки при обработке строк: Если строка v содержит менее 100 символов,
//    при попытке обращения к v[:100] возникнет паника (runtime panic), так
//    как индекс будет выходить за пределы длины строки.
// 3) глобальная переменная (Проблемы с параллельностью и потокобезопасностью - race conditions,
//    могут быть изменены из любой части программы, создает скрытую зависимость между различными частями программы)
