package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func task6() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})                           // Канал для сигнала остановки
	ctx, cancel := context.WithCancel(context.Background()) // Контекст с функцией отмены
	var stopFlag int32                                      // Атомарный флаг для остановки

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh: // Случай остановки через закрытие канала
				fmt.Println("Остановка: канал stopCh закрыт")
				return
			case <-ctx.Done(): // Случай остановки через контекст
				fmt.Println("Остановка: контекст отменен")
				return
			default:
				if atomic.LoadInt32(&stopFlag) > 0 { // Случай остановки через атомарный флаг
					fmt.Println("Остановка: атомарный флаг установлен")
					return
				}
				// Выполняем полезную работу
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Остановка через закрытие канала через 2 секунды
	go func() {
		time.Sleep(2 * time.Second)
		close(stopCh)
	}()

	// Остановка через контекст через 4 секунды
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	// Остановка через атомарный флаг через 6 секунд
	go func() {
		time.Sleep(6 * time.Second)
		atomic.StoreInt32(&stopFlag, 1)
	}()

	wg.Wait() // Ожидание завершения горутины
	fmt.Println("Главная горутина завершила работу")
}
