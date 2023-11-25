package main

import "fmt"

func task13(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}
