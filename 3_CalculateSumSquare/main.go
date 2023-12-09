package main

import (
	"fmt"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func calculateSquare(num int) int {
	// Выводим квадрат
	return num * num
}

func main() {
	// Инициализируем слайс с элементами 2, 4, 6, 8, 10
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для суммирования квадратов чисел
	sumChan := make(chan int)

	// Запускаем анонимную горутину
	go func() {
		// Запускаем цикл, который итерируется по элементам слайса
		for _, num := range numbers {
			// Результат функции calculateSquare записываем в канал
			sumChan <- calculateSquare(num)
		}
		close(sumChan)
	}()

	// Инициализируем переменную sum
	var sum int
	// С помощью цикла считываем данные из канала
	for num := range sumChan {
		// Суммируем
		sum += num
	}

	// Выводим результат
	fmt.Printf("sum of squares: %d\n", sum)
}
