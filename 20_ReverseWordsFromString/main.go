package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

// Вариант 1
func reverseWordsFromString(str string) string {
	var words []string

	for i := 0; i < len(str); i++ { // Разбиваем строку на отдельные слова через пробел
		if str[i] == ' ' {
			words = append(words, str[:i])
			str = str[i+1:]
			i = 0
		}
		if i == len(str)-1 {
			words = append(words, str)
		}
	}

	/*
		Оператор for инициализируем переменные i и j значением 0 и len(words)-1, где len(words) - длина слайса words.
		Устанавливаем условие выполнения цикла (пока i меньше j).
		После каждой итерации цикла происходит инкремент i и декремент j.
	*/
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
		/*
			Происходит обмен значениями между words[i] и words[j].
			Таким образом мы изменяем порядок значений в слайсе.
			На каждой итерации переменная i указывает на слово в начале слайса,
			а переменная j указывает на слово в конце слайса.
		*/
	}

	return strings.Join(words, " ")
}

// Вариант 2
func reverseWordsFromString1(str string) string {
	words := strings.Split(str, " ") // Разбиваем строку на отдельные слова через пробел
	start, end := 0, len(words)-1    // start указатель на первое слово, а end на последнее

	/*
		Устанавливаем условие выполнения цикла (start должен быть меньше end).
		После каждой итерации цикла происходит инкремент start и декремент end.
	*/

	for ; start < end; start, end = start+1, end-1 {
		words[start], words[end] = words[end], words[start]

		/*
			Происходит обмен значениями между words[start] и words[end].
			Таким образом мы изменяем порядок значений в слайсе.
		*/
	}
	return strings.Join(words, " ")
}

// Вариант 3
func reverseWordsFromString2(str string) string {
	words := strings.Fields(str) // Разбиваем строку на отдельные слова

	/*
		Оператор for инициализируем переменные i и j значением 0 и len(words)-1, где len(words) - длина слайса words.
		Устанавливаем условие выполнения цикла (пока i меньше j).
		После каждой итерации цикла происходит инкремент i и декремент j.
	*/
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
		/*
			Происходит обмен значениями между words[i] и words[j].
			Таким образом мы изменяем порядок значений в слайсе.
			На каждой итерации переменная i указывает на слово в начале слайса,
			а переменная j указывает на слово в конце слайса.
		*/
	}

	return strings.Join(words, " ") // объединяем слова в строку разделяя их пробелом и возвращаем
}

func main() {
	str := "snow dog sun"
	fmt.Printf("Original: %s\n", str)                                    // выводим оригинальную строку
	fmt.Printf("Reverse: %s\n", reverseWordsFromString("snow dog sun"))  // выводим перевернутую строку
	fmt.Printf("Reverse: %s\n", reverseWordsFromString1("snow dog sun")) // выводим перевернутую строку
	fmt.Printf("Reverse: %s\n", reverseWordsFromString2("snow dog sun")) // выводим перевернутую строку
}
