package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

// Интерфейс который определяет методы Set
type MyMap interface {
	Set(key string, value interface{})
}

type myMap struct {
	// Использую обычный мьютекс для предотвращения одновременной записи
	mu sync.Mutex
	m  map[string]interface{}
}

/*
Конструктор который создает новый экземпляр myMap и возвращает его в виде MyMap(интерфейс)
это позволяет использовать интерфейс MyMap для взаимодействия с картой,
скрывая детали реализации
*/
func NewMap() MyMap {
	return &myMap{
		m: make(map[string]interface{}),
	}
}

/*
Метод Set блокирует мьютекс чтобы защитить доступ
к карте при записи значения, после записи разблокирует мьютекс
*/
func (m *myMap) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

func main() {
	// Создаем карту с мьютексом
	m := NewMap()

	// Инициализируем WaitGroup
	var wg sync.WaitGroup
	// Количество рабочих горутин
	numWorkers := 5

	// Добавляем в счетчик количество горутин
	wg.Add(numWorkers)

	// Запускаем горутину
	go func() {
		// Запускаем цикл который выполнится numWorkers раз
		for i := 0; i < numWorkers; i++ {
			go func(workerID int) {
				//В конце функции декрементирует счетчик
				defer wg.Done()

				//Ключ
				key := fmt.Sprintf("Key%d", workerID)
				//Значение
				value := workerID * 10

				//Вызываем метод Set
				m.Set(key, value)

				fmt.Printf("Worker %d wrote value %d for key %s\n", workerID, value, key)
			}(i)
		}
	}()

	// Блокирует главную горутину до тех пор пока счетчик не станет равным нулю
	wg.Wait()
}
