package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("TASK 1:")
	task1()
	fmt.Println("================")

	fmt.Println("TASK 2:")
	task2([]int{2, 4, 6, 8, 10})
	fmt.Println("================")

	fmt.Println("TASK 3:")
	fmt.Println(task3([]int{2, 4, 6, 8, 10})) // =220
	fmt.Println("================")

	fmt.Println("TASK 4:")
	task4() // кол-во воркеров вводим с клавиатуры
	fmt.Println("================")

	fmt.Println("TASK 5:")
	task5() // время работы программы вводим с клавиатуры
	fmt.Println("================")

	fmt.Println("TASK 6:")
	task6()
	fmt.Println("================")

	fmt.Println("TASK 7:")
	task7()
	fmt.Println("================")

	fmt.Println("TASK 8:")
	fmt.Println(task8(10, 1, false)) // делает из 1010 => 1000
	fmt.Println("================")

	fmt.Println("TASK 9:")
	task9([]int{1, 2, 3, 4, 5})
	fmt.Println("================")

	fmt.Println("TASK 10:")
	task10([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5})
	fmt.Println("================")

	fmt.Println("TASK 11:")
	fmt.Println(task11([]int{4, 2, 5, 7, -10}, []int{2, 11, 4, 6})) // множества перескаются в значениях 2 и 4
	fmt.Println("================")

	fmt.Println("TASK 12:")
	fmt.Println(task12([]string{"cat", "cat", "dog", "cat", "tree"}))
	fmt.Println("================")

	fmt.Println("TASK 13:")
	task13(111, 222) // 222 111
	fmt.Println("================")

	fmt.Println("TASK 14:")
	var x interface{}
	x = 42
	task14(x)
	x = "hello world"
	task14(x)
	x = true
	task14(x)
	x = make(chan float64)
	task14(x)
	x = 3.14
	task14(x)
	fmt.Println("================")

	fmt.Println("TASK 15:")
	task15()
	fmt.Println("================")

	fmt.Println("TASK 16:")
	arr16 := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Неотсортированный массив:", arr16)
	quickSort(arr16, 0, len(arr16)-1)
	fmt.Println("Отсортированный массив:", arr16)
	fmt.Println("================")

	fmt.Println("TASK 17:")
	arr17 := []int{-10, -5, -1, 0, 3, 7, 9, 15}
	target := -1
	fmt.Printf("Ищем элемент со значением %v в массиве %v\n", target, arr17)
	result17 := task17(arr17, target)
	if result17 != -1 {
		fmt.Printf("Искомый элемент найден на позиции %v\n", task17(arr17, target))
	} else {
		fmt.Println("Искомый элемент не найден")
	}
	fmt.Println("================")

	fmt.Println("TASK 18:")
	task18()
	fmt.Println("================")

	fmt.Println("TASK 19:")
	str := "главрыба"
	reversed := task19(str)
	fmt.Printf("Оригинальная строка: %s\nПеревернутая строка: %s\n", str, reversed)
	fmt.Println("================")

	fmt.Println("TASK 20:")
	str20 := "snow dog sun"
	reversed20 := task20(str20)
	fmt.Printf("Оригинальная строка: %s\nПеревернутая строка: %s\n", str20, reversed20)
	fmt.Println("================")

	fmt.Println("TASK 21:")
	oldPlayer := &OldMusicPlayer{}
	adapter := &MusicPlayerAdapter{player: oldPlayer} // Создаем адаптер для старого плеера
	adapter.PlaySong("Nirvana - Come As You Are")     // Используем адаптер для воспроизведения песни
	fmt.Println("================")

	fmt.Println("TASK 22:")
	task22()
	fmt.Println("================")

	fmt.Println("TASK 23:")
	slice23 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Оригинальный слайс: %v\n", slice23)
	i := 2
	removedItemSlice := task23(slice23, i)
	fmt.Printf("Удален %v-й элемент\nНовый слайс: %v\n", i, removedItemSlice)
	fmt.Println("================")

	fmt.Println("TASK 24:")
	p1 := NewPoint(1, 2) //Создаем две точки
	p2 := NewPoint(4, 6)
	distance := p1.Distance(p2) // Расчет расстояния между p1 и p2
	fmt.Printf("Расстояние между точками: %v\n", distance)
	fmt.Println("================")

	fmt.Println("TASK 25:")
	sleepingTime := 10
	fmt.Printf("Спим %v секунд...\n", sleepingTime)
	task25(time.Duration(sleepingTime) * time.Second)
	fmt.Println("Проснулись!")
	fmt.Println("================")

	fmt.Println("TASK 26:")
	fmt.Println(task26("abcd"))      // true
	fmt.Println(task26("abCdefAaf")) // false
	fmt.Println(task26("	aabcd"))    // false
	fmt.Println("================")

}
