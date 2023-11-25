package main

import (
	"fmt"
	"reflect"
)

func task14(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Тип переменной: %T, значение: %v\n", v, i)
	case string:
		fmt.Printf("Тип переменной: %T, значение: %v\n", v, i)
	case bool:
		fmt.Printf("Тип переменной: %T, значение: %v\n", v, i)
	default:
		t := reflect.TypeOf(i)
		switch t.Kind() {
		case reflect.Chan:
			fmt.Printf("Тип переменной: %v\n", t)
		default:
			fmt.Printf("Неизвестный тип: %T\n", v)
		}
	}
}
