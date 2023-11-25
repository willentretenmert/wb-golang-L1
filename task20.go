package main

import (
	"strings"
)

func task20(s string) string {
	words := strings.Fields(s) // Разделение строки на слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ") // Соединение слов в строку
}
