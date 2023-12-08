package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func mySleep(t time.Duration) {
	/*
		Блокирует главную горутину
		пока не будет получено значение
		из этого канала, значение будет получено
		из канала по истечению времени
	*/
	<-time.After(t)
}

func main() {
	start := time.Now() // Для демонстрации
	fmt.Println("started")

	fmt.Println("sleep 3 seconds")
	mySleep(3 * time.Second) // Вызываем функцию mySleep передав ей 3 секунды

	fmt.Printf("Finish by time: %s", time.Since(start)) // Для демонстрации
}
