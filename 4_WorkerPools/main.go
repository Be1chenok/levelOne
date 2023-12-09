package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

// Структура worker
type Worker struct {
	ID int
}

// Конструктор worker
func NewWorker(id int) *Worker {
	return &Worker{
		ID: id,
	}
}

// Метод worker'а DoWork
func (w *Worker) DoWork(ch <-chan string, wg *sync.WaitGroup) {
	// В конце функции декрементируем счетчик
	defer wg.Done()

	// Чтение данных из канала и вывод в stdout
	for data := range ch {
		fmt.Printf("Worker %d: %s\n", w.ID, data)
	}
}

func main() {
	// Инициализируем переменную numWorkers
	var numWorkers int

	// Присваиваем кол-во worker'ов с консоли
	fmt.Scan(&numWorkers)

	// Создаем канал для данных
	ch := make(chan string)

	// Записываем данные в канал
	go func() {
		for i := 1; i <= 20; i++ {
			ch <- fmt.Sprintf("Data %d", i)
		}
		close(ch)
	}()

	var wg sync.WaitGroup

	// Создаем n worker'ов и запускаем метод DoWork у каждого worker'a
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Добавляем в счетчик 1
		worker := NewWorker(i)
		go worker.DoWork(ch, &wg)
	}

	// Создаем буферизованный канал c размером буфера 1 для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)
	// SIGINT сигнал OC при нажатии Ctrl+C
	<-stop

	wg.Wait()
}
