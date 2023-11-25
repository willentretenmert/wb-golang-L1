package main

import (
	"fmt"
	"math"
)

func task10(temperatures []float64) {
	// Словарь для группировки температур
	groupedTemperatures := make(map[int][]float64)

	// Группировка значений
	for _, temp := range temperatures {
		key := int(math.Floor(temp/10)) * 10
		groupedTemperatures[key] = append(groupedTemperatures[key], temp)
	}

	// Вывод результатов
	for key, temps := range groupedTemperatures {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
