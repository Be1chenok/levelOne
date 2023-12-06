package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет,
	что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.
*/

func uniqueChar(str string) bool {
	str = strings.ToLower(str) // Переводим всю строку в нижний регистр

	charMap := make(map[rune]bool) // Создаем карту рун

	// Проверяем каждый символ в строке
	for _, char := range str {
		if charMap[char] { // Если символ находится в карте то возвращаем false
			return false
		}

		charMap[char] = true // Добавляем символ в карту
	}

	return true // Возвращаем true
}

func main() {
	str := "abCdefAaf"
	fmt.Println(uniqueChar(str))
}
