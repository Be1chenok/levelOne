package main

import "fmt"

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

	Этот код может привести к переполнению памяти
*/

// Избавился от глобальной переменной, т. к. их изменение может быть неожиданным

// Возвращает только нужную часть строки
func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]

}

func main() {
	// сохраняем значение в локальной переменной
	justString := someFunc()
	fmt.Println(justString) // Последующая работа с justString
}

func createHugeString(size int) string {
	return ""
}
