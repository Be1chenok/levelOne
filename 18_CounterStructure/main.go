package main

import (
	"fmt"
	"sync"
)

type SimpleCounter interface {
	Increment()
	GetValue() int
}

// Использую интерфейс для сокрытия деталей реализации
func NewSimpleCounter() SimpleCounter {
	return &simpleCounter{}
}

// структура simpleCounter
type simpleCounter struct {
	value int
	// использую мьютекс для корректной записи
	mutex sync.Mutex
}

// Метод для инкремента
func (c *simpleCounter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// Метод для получения данных
func (c *simpleCounter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func main() {
	// Создаю экземпляр simpleCounter
	counter := NewSimpleCounter()

	// Инициализирую WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 9999; i++ {
		// Каждый увеличиваю счетчик на 1
		wg.Add(1)
		go func() {
			// В конце функции уменьшаю счетчик на 1
			defer wg.Done()
			// Вызываю метод инкремент
			counter.Increment()
		}()
	}

	// Блокируем главный поток пока WaitGroup не обнулится
	wg.Wait()
	fmt.Println(counter.GetValue()) // Выводим результат
}
