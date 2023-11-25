package main

import (
	"math"
)

// Point структура, представляющая точку с приватными полями x и y.
type Point struct {
	x, y float64
}

// NewPoint конструктор для создания новой точки.
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance метод для расчета расстояния между двумя точками.
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}
