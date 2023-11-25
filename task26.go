package main

import (
	"strings"
)

func task26(s string) bool {
	s = strings.ToLower(s) // Преобразование строки в нижний регистр
	charSet := make(map[rune]bool)

	for _, char := range s {
		if _, exists := charSet[char]; exists {
			// Если символ уже встречался, возвращаем false
			return false
		}
		charSet[char] = true
	}

	return true
}
