package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func task5() {
	ch := make(chan string)
	done := make(chan bool)
	var duration int
	fmt.Print("Введите время работы программы: ")
	fmt.Scan(&duration)

	// Горутина для отправки данных в канал
	go func() {
		for i := 0; ; i++ {
			ch <- RandStringBytes(7)           // Решил использовать случайные строки как отправляемые данные
			time.Sleep(400 * time.Millisecond) // Отправляем данные каждую секунду
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println("Получено значение:", v)
			case <-done:
				return
			}
		}
	}()

	// Запускаем таймер для завершения программы
	fmt.Printf("Программа будет работать %d секунд...\n", duration)
	time.Sleep(time.Duration(duration) * time.Second)
	done <- true
	fmt.Println("Время вышло, программа завершена.")
}
