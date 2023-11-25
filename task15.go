package main

import (
	"fmt"
	"math/rand"
)

/*
Проблема в исходном коде была в том, что строка v в памяти привязана к justString,

	так как justString ссылается на значение v.

Этого можно избежать, взяв нужный срез из v, явно создав новую строку (я сделал через string()),

	и проинициализировать justString уже при помощи этой новой строки.
*/
var justString string

func task15() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
	fmt.Println(justString)
}

func createHugeString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
