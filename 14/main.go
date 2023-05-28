package main

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func getType(value interface{}) string {
	switch v := value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "channel"
	default:
		return getUnknownType(v)
	}
}

// функция getUnknownType() отвечает за получение строкового представления неизвестного типа.
func getUnknownType(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func main() {
	var i interface{}
	i = 42
	fmt.Println(getType(i)) // Выводит: int

	i = "Hello, World!"
	fmt.Println(getType(i)) // Выводит: string

	i = true
	fmt.Println(getType(i)) // Выводит: bool

	ch := make(chan interface{})
	i = ch
	fmt.Println(getType(i)) // Выводит: channel

	i = 3.14
	fmt.Println(getType(i)) // Выводит: float64
}
