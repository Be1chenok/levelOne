package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a := 1 // Инициализируем переменную a с значение 1
	b := 2 // Инициализируем переменную b с значением 2

	fmt.Printf("a: %d, b: %d\n", a, b) // Выводим в консоль a и b

	a, b = b, a // Обмен значений между a и b

	fmt.Printf("a: %d, b: %d\n", a, b) // Выводим в консоль a и b
}
