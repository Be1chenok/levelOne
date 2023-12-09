package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func intersection(arr1, arr2 []int) []int {
	// Создаем карту set1 для хранения уникальных значений из слайса arr1
	set1 := make(map[int]struct{}, len(arr1))

	// Записываем в карту set1 значения из слайса arr1
	for _, val := range arr1 {
		set1[val] = struct{}{}
	}

	// Создаем переменную result для результата
	var result []int

	// Итерируемся по слайсу arr1 по ключу
	for _, val := range arr2 {
		// Если такой ключ есть в карте set1 то добавляем значение в результат
		if _, ok := set1[val]; ok {
			result = append(result, val)
		}
	}
	return result //Возвращаем результат
}

func main() {
	set1 := []int{1, 3, 5, 7, 9}
	set2 := []int{2, 4, 5, 6, 8, 9, 10}
	fmt.Println(intersection(set1, set2))
}
