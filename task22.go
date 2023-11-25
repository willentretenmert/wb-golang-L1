package main

import (
	"fmt"
)

func task22() {
	var a, b int64 = 1 << 21, 1 << 22 // Значения больше 2^20

	// Сложение
	res1 := a + b
	fmt.Printf("%d + %d = %d\n", a, b, res1)

	// Вычитание
	res2 := a - b
	fmt.Printf("%d - %d = %d\n", a, b, res2)

	// Умножение
	res3 := a * b
	fmt.Printf("%d * %d = %d\n", a, b, res3)

	// Деление
	res4 := a / b
	fmt.Printf("%d / %d = %d\n", a, b, res4)

}
