package main

import (
	"fmt"
	"sync"
)

// Counter структура счетчика с защитой для конкурентного доступа.
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment увеличивает счетчик на 1.
func (cnt *Counter) Increment() {
	cnt.mu.Lock()   // Блокировка доступа к счетчику
	cnt.value++     // Инкрементирование счетчика
	cnt.mu.Unlock() // Разблокировка доступа
}

// Value возвращает текущее значение счетчика.
func (cnt *Counter) Value() int {
	cnt.mu.Lock()         // Блокировка для безопасного чтения значения
	defer cnt.mu.Unlock() // Автоматическая разблокировка после возврата значения
	return cnt.value
}

func task18() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Запускаем 1000 горутин для инкрементирования счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
