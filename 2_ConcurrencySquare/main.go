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
	var wg sync.WaitGroup

	// Добавляем в счетчик длину слайса
	wg.Add(len(numbers))

	// Запускаем горутину
	go func() {
		// Запускаем цикл, который итерируется по элементам слайса
		for _, num := range numbers {
			// Запускаем горутину для функции calculateSquare и передаем в нее элемент из слайса и ссылку на WaitGroup
			go calculateSquare(num, &wg)
		}
	}()

	// Блокируем главную горутину, пока не обнулится счетчик
	wg.Wait()
}
