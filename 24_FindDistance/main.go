package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

// Структура Point
type Point struct {
	// Координаты точки
	x, y float64
}

// Конструктор Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Расчёт дистанции между точками используя формулу Расстояния
func Distance(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point2 := NewPoint(1.0, 2.0)
	point1 := NewPoint(4.0, 6.0)

	fmt.Printf("Distance between points: %.2f\n", Distance(point1, point2))
}
