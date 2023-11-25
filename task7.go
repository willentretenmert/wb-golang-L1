package main

import (
	"fmt"
	"sync"
)

func task7() {
	var wg sync.WaitGroup
	m := &sync.Map{}

	// Функция для записи данных в мапу
	writeToMap := func(key, value int) {
		m.Store(key, value) // Безопасная запись в sync.Map
	}

	// Запускаем несколько горутин для записи данных в мапу
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			writeToMap(i, i*i) // Запись квадрата числа в мапу
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Вывод содержимого мапы
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true // Продолжаем итерацию
	})
}
