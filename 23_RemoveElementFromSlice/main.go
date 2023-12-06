package main

import "fmt"

/*
	Удалить i-ый элемент из слайса.
*/

func removeElementFromSlice(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
	/*
		Создаем новый слайс(slice[i:]) который содержит все элементы до i-го элемента
		Так же создаем слайс(slice[i+1:]) который содержит все элементы после i-го элемента
		Затем объединяем эти два слайса с помощью функции append
		Возвращаем результат
	*/
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(removeElementFromSlice(arr, 1))
}
