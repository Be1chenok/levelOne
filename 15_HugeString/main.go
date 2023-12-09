package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.

var justString string

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}

	createHugeString(1 << 10) создает строку размером в 1024 элемента
	Хотим использовать первые 100 символов
	В примере все остальные элементы нам не нужны, но оставшиеся 924 элемента
	не будут удалены, а сборщик мусора удалит их только тогда,
	когда justString выйдет из области видимости
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Решаем проблему при помощи клонирования
	justString = strings.Clone(v[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	var str string
	for i := 0; i < size; i++ {
		str += "q"
	}
	return str
}
