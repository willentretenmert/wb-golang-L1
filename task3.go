package main

import (
	"sync"
)

func task3(numbers []int) int {
	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))
	sum := 0

	wg.Add(len(numbers))

	// Запускаем горутину в цикле
	for _, num := range numbers {
		go func(num int) {
			defer wg.Done()
			ch <- num * num // Возводим в квадрат и отправляем в канал
		}(num)
	}

	// Запускаем горутину для подсчета суммы квадратов
	go func() {
		wg.Wait() // Горутина при старте ждет завершения всех остальных горутин с вычислением квадратов
		close(ch) // Закрываем канал
	}()

	// После разблокировки выполнения программы суммируем результаты
	for square := range ch {
		sum += square
	}

	return sum
}
