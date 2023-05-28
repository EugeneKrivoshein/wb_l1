package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования)

import "fmt"

//структура Human
type Human struct {
	name string
	age  int
}

// метод возвращающий имя
func (h *Human) GetName() string {
	return h.name
}

//метод устанавливающий имя человека
func (h *Human) SetName(name string) {
	h.name = name
}

//структура представляющая действие человека
type Action struct {
	*Human //встраивание указателя на структуру Human
}

// метод возвращающий возраст
func (a *Action) GetAge() int {
	return a.age
}

// метод устанавливающий возраст
func (a *Action) SetAge(age int) {
	a.age = age
}

//метод который будет доступен через структуру Action
func (a *Action) SayHello() {
	fmt.Printf("Hello, my name is %s!\n", a.GetName())
}

//метод, который будет доступен через структуру Action
func (a *Action) Introduce() {
	fmt.Printf("I am %d years old.\n", a.GetAge())
}

func main() {
	h := &Human{
		name: "Alex",
		age:  10,
	}
	a := &Action{
		Human: h,
	}

	a.SayHello()  // Вызов метода SayHello через структуру Action
	a.Introduce() // Вызов метода Introduce через структуру Action

	// Изменение имени и возраста через методы встроенной структуры
	a.SetName("Eugene")
	a.SetAge(20)

	fmt.Println("New Data:")

	a.SayHello()
	a.Introduce()
}
