package main

func task11(set1, set2 []int) []int {
	// Словарь для отслеживания элементов первого множества
	set1Map := make(map[int]bool)
	for _, item := range set1 {
		set1Map[item] = true
	}

	// Поиск пересечений, будем сравнивать значения в словаре со значениями из второго множества
	var result []int
	for _, item := range set2 {
		if set1Map[item] {
			result = append(result, item)
		}
	}
	return result
}
