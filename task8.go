package main

// n - данное число, i - позиция бита, b - нужное значение бита
func task8(n int64, i uint, b bool) int64 {
	if b {
		// Устанавливаем бит i в 1
		return n | (1 << i)
	}
	// Устанавливаем бит i в 0
	return n & ^(1 << i)
}
