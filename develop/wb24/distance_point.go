package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x, y float64
}

// Конструктор для структуры Point
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Функция для вычисления расстояния между двумя точками
func Distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки с использованием конструктора
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками
	distance := Distance(point1, point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
