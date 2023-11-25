package main

import (
	"fmt"
)

func task9(numbers []int) {
	// Создаем два канала
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Горутина для чтения из массива и записи в первый канал
	go func() {
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan)
	}()

	// Горутина для чтения из первого канала, удвоения числа и записи во второй канал
	go func() {
		for num := range inputChan {
			outputChan <- num * 2
		}
		close(outputChan)
	}()

	// Чтение из второго канала и вывод в stdout
	for result := range outputChan {
		fmt.Println(result)
	}
}
