package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func calculateSquare(num int, wg *sync.WaitGroup) {
	// В конце функции декрементируем счетчик
	defer wg.Done()
	// Выводим квадрат
	fmt.Println(num * num)
}

func main() {
	// Инициализируем слайс с элементами 2, 4, 6, 8, 10
	numbers := []int{2, 4, 6, 8, 10}

	// Инициализируем WaitGroup
	wg := new(sync.WaitGroup)

	// Запускаем цикл, который итерируется по элементам слайса
	for _, num := range numbers {
		wg.Add(1) // Добавляем в счетчик 1
		// Запускаем горутину для функции calculateSquare и передаем в нее элемент из слайса и ссылку на WaitGroup
		go calculateSquare(num, wg)
	}

	// Блокируем главную горутину, пока не обнулится счетчик
	wg.Wait()
}
