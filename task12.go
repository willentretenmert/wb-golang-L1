package main

func task12(strings []string) map[string]struct{} {
	// Создаем словарь с пустыми значенями и ключом-строкой из входного массива
	set := make(map[string]struct{})
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}
