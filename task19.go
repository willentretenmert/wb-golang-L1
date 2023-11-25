package main

func task19(s string) string {
	// Преобразуем входную строку в список символов из Rune, это позволит корректно обрабатывать Юникод-символы, так как они весят более 1 байта
	symbol := []rune(s)
	for i, j := 0, len(symbol)-1; i < j; i, j = i+1, j-1 {
		symbol[i], symbol[j] = symbol[j], symbol[i]
	}
	return string(symbol)
}
