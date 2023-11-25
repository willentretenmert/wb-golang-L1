package main

func task23(slice []int, i int) []int {
	// Создаем новый слайс с длиной, на один меньше исходного
	newSlice := make([]int, 0, len(slice)-1)

	// Добавляем в новый слайс все элементы, кроме i-го
	for j, val := range slice {
		if j != i {
			newSlice = append(newSlice, val)
		}
	}

	return newSlice
}
