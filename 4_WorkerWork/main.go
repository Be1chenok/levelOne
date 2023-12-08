package main

import (
	"context"
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

func main() {

	// Инициализируем переменную n
	var n int
	// Получаем ввод с консоли
	fmt.Scan(&n)

	// Создаем контекст c отменой для завершения горутин
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	// Анонимная горутина для записи в канал
	go func() {
		for i := 1; ; i++ {
			select {
			case ch <- i: // Записываем в канал
			case <-ctx.Done(): // Выйдем из цикла когда отменится контекст
				return
			}
		}
	}()

	var wg sync.WaitGroup // Создаем WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)                  // Добавляем в счетчик 1
		go worker(i, ch, &wg, ctx) // Запускаем горутину с функцией worker
	}

	stop := make(chan os.Signal, 1) // Буферизированный канал для сигнала операционной системы
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// SIGINT - сигнал прерывания, при нажатии Ctrl+C
	// SIGTERM - сигнал завершения для плавного выхода из программы
	<-stop // Ожидает сигнал SIGINT или SIGTERM, когда сигнал будет получен будет продолжено выполнение следующих инструкций

	cancel() // Отменяем контекст тем самым выходим из циклов в горутинах

	close(ch) // Закрываем канал
	wg.Wait() // Блокируем главную горутину пока не обнулится счетчик

	fmt.Println("finish")
}

func worker(id int, ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done() // В конце функции декрементируем счетчик
	for {
		select {
		case value := <-ch: // Считываем значение из канала
			fmt.Printf("Worker %d read value: %d\n", id, value) // Выводим значение
		case <-ctx.Done(): // Выйдет из цикла когда отменится контекст
			return
		}
	}
}
