package main

import (
	"fmt"
	"sync"
)

func task2(numbers []int) {
	// Создаем WorkGroup
	var wg sync.WaitGroup

	// Добавляем в счетчик WorkGroup количество наших будущих горутин, в нашем случае длину массива numbers
	wg.Add(len(numbers))

	// Проходимся циклом по массиву чисел
	for _, num := range numbers {
		// Вызываем анонимную функцию с передачей в нее элемента num
		go func(num int) {
			defer wg.Done() // Завершаем горутину в WorkGroup
			fmt.Printf("Квадрат числа %d равен %d\n", num, num*num)
		}(num)
	}

	wg.Wait() // Блокируем выполнение программы, пока все горутины не завершатся
}
