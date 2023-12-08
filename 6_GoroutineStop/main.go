package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func goroutineStop() {
	// Способ 1 использование контекста

	ctx, cancel := context.WithCancel(context.Background())

	// Выполняется до получения сигнала ctx.Done
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("example1 finish")
				return
			default:
				fmt.Println("example1 in progress")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel() // Отменяем контекст

	time.Sleep(1 * time.Second) // Для наглядной демонстрации

	// Способ 2 использование стоп флага

	stopFlag := false

	// Пока флаг равен false горутина выполняется
	go func() {
		for !stopFlag {
			fmt.Println("example2 in progress")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("example2 finish")
	}()

	time.Sleep(5 * time.Second)
	stopFlag = true // Меняем флаг на true и горутина выходит из цикла и завершается

	time.Sleep(1 * time.Second) // Для наглядной демонстрации

	// Способ 3 использование канала

	stopChan := make(chan struct{})

	// Горутина выполняется в цикле до получения сигнала от stopChan
	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("example3 finish")
				return
			default:
				fmt.Println("example3 in progress")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(stopChan) // Горутина выходит из цикла и завершается

	time.Sleep(1 * time.Second) // Для наглядной демонстрации

	// Использую WaitGroup для демонстрации
	var wg sync.WaitGroup
	wg.Add(1)

	// Способ 4 использование таймера

	timer := time.NewTimer(5 * time.Second)

	// Горутина завершится по истечению времени
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("example4 finish by time")
				wg.Done()
				return
			default:
				fmt.Println("example4 in progress")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	wg.Wait()
}

func main() {
	goroutineStop()
}
