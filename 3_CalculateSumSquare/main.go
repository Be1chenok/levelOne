package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func calculateSquare(num int, wg *sync.WaitGroup) int {
	// В конце функции декрементируем счетчик
	defer wg.Done()
	// Выводим квадрат
	return num * num
}

func main() {
	// Инициализируем слайс с элементами 2, 4, 6, 8, 10
	numbers := []int{2, 4, 6, 8, 10}

	// Инициализируем WaitGroup
	var wg sync.WaitGroup

	// Добавляем в счетчик длину слайса
	wg.Add(len(numbers))

	// Создаем канал для суммирования квадратов чисел
	sumChan := make(chan int)

	// Запускаем анонимную горутину
	go func() {
		// Запускаем цикл, который итерируется по элементам слайса
		for _, num := range numbers {
			/*
				Создаем горутину, которая выполняет анонимную функцию
				Передаем в горутину num чтобы избежать проблем с захватом одной и той же переменной
			*/
			go func(num int) {
				// Результат функции calculateSquare записываем в канал
				sumChan <- calculateSquare(num, &wg)
			}(num)
		}
	}()

	// Закрываем канал, когда все горутины завершены
	go func() {
		defer close(sumChan)
		// Блокируем главную горутину пока не обнулится счетчик
		wg.Wait()
	}()

	// Инициализируем переменную sum равную 0
	sum := 0
	// С помощью цикла считываем данные из канала
	for num := range sumChan {
		// Суммируем
		sum += num
	}

	// Выводим результат
	fmt.Printf("sum of squares: %d\n", sum)
}
