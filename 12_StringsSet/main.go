package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

func set(strs []string) []string {
	// Создаем карту strsSet для хранения уникальных значений из слайса strs
	strsSet := make(map[string]struct{}, len(strs))

	// Записываем в карту strsSet значения из слайса strs
	for _, str := range strs {
		strsSet[str] = struct{}{}
	}

	// Создаем переменную result для результата
	var result []string

	// Проходимся по ключам карты и добавляем значение в результат
	for key := range strsSet {
		result = append(result, key)
	}

	return result // Возвращаем результат
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set(words))
}
