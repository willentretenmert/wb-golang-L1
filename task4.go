package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch <-chan int, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n, ok := <-ch:
			if !ok {
				fmt.Printf("Воркер %d завершает работу\n", id)
				return
			}
			fmt.Printf("Воркер %d получил число: %d\n", id, n)

		case <-stopCh:
			fmt.Printf("Воркер %d получил сигнал к завершению\n", id)
			return
		}
	}
}

func task4() {
	var wg sync.WaitGroup
	ch := make(chan int)                // Канал для передачи чисел воркерам.
	stopCh := make(chan struct{})       // Канал для сигнала остановки работы воркеров.
	signalCh := make(chan os.Signal, 1) // Канал для системных сигналов
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Чтение количества воркеров
	var workerCount int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&workerCount)

	// Запуск воркеров
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go worker(i, ch, stopCh, &wg)
	}

	// Горутина для обработки сигнала завершения.
	go func() {
		<-signalCh
		fmt.Println("\nСигнал к завершению работы получен, завершаем воркеров...")
		close(stopCh) // Отправляем сигнал воркерам прекратить работу
	}()

	// Главная горутина, записывающая данные в канал
	counter := 0
	for {
		select {
		// Проверка сигнала остановки в главной горутине.
		case <-stopCh:
			fmt.Println("Завершение работы главной горутины...")
			close(ch) // Закрываем канал ch, чтобы воркеры могли завершить чтение и выйти
			wg.Wait() // Ожидаем завершения всех воркеров
			fmt.Println("Работа завершена.")
			return
		default:
			ch <- counter // Отправляем значение счетчика в канал ch
			time.Sleep(100 * time.Millisecond)
			counter++
		}
	}
}
