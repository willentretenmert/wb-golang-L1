package main

// splitter разделяет часть массива на две секции
func splitter(arr []int, low, high int) int {
	half := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < half {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// А сама функция сортировки рекурсивно сортирует части массива
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := splitter(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}
