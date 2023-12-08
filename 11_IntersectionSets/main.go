package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func intersection(arr1, arr2 []int) []int {

	// Создаем карту set1 для хранения уникальных значений из слайса arr1
	set1 := make(map[int]bool, len(arr1))

	// Записываем в карту set1 значения из слайса arr1
	for _, val := range arr1 {
		set1[val] = true
	}

	// Создаем карту set2 для хранения уникальных значений из слайса arr2
	set2 := make(map[int]bool, len(arr2))

	// Записываем в карту set2 значения из слайса arr2
	for _, val := range arr2 {
		set2[val] = true
	}

	// Создаем переменную result для результата
	var result []int

	// Итерируемся по карте set1 по ключу
	for key := range set1 {
		// Если такой ключ есть в карте set2 то добавляем значение в результат
		if set2[key] {
			result = append(result, key)
		}
	}

	return result // Возвращаем результат
}

func main() {
	set1 := []int{1, 3, 5, 7, 9}
	set2 := []int{2, 4, 5, 6, 8, 9, 10}
	fmt.Println(intersection(set1, set2))
}
