package main

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	distance := Distance(p1, p2)
	fmt.Println("Расстояние между точками:", distance)
}

// Логика нахождения расстояния между двумя точками основана на применении
// теоремы Пифагора для прямоугольного треугольника, образованного двумя сторонами,
// которые соответствуют разнице между координатами x и y двух точек.
