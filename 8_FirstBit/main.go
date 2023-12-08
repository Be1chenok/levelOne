package main

import "fmt"

func main() {
	var num int64
	fmt.Println("enter number")
	fmt.Scan(&num) // записываем в num с консоли

	var bitIndex int // Тип int64 имеет 64 бита от 0 до 63 индекса
	fmt.Println("enter bit index from 0 to 63")
	fmt.Scan(&bitIndex)                // записываем в bitIndex с консоли
	if bitIndex < 0 || bitIndex > 63 { // Проверяем верность введенных данных
		fmt.Printf("wrong number\n: %d", bitIndex)
		return
	}

	var bitValue int
	fmt.Println("enter bit value from 0 to 1") // Бит можно установить только в значения 0 и 1
	fmt.Scan(&bitValue)                        // записываем в bitValue с консоли
	if bitValue < 0 || bitValue > 1 {          // Проверяем верность введенных данных
		fmt.Printf("wrong number\n: %d", bitValue)
		return
	}

	fmt.Println(setBit(num, bitIndex, bitValue)) // Выводим результат функции setBit
}

func setBit(num int64, bitIndex, bitValue int) int64 {
	/*
		Создаем маску бита, с одним битов, сдвигаем этот бит влево
		на значение равное bitIndex
	*/
	mask := int64(1) << bitIndex
	/*
		Инвертирование маски, все биты в маске,
		кроме выбранного устанавливаются в единицу
	*/
	mask = ^mask
	/*
		Все биты установленные в единицу такими же и остаются,
		а выбранный бит обнуляется
	*/
	num &= mask
	/*
		Создаем новое значение бита в bitValue
		Значение bitValue будет находится в выбранной
		позиции bitIndex
	*/
	bit := int64(bitValue) << bitIndex
	// Установка бита
	num |= bit
	return num // Возвращаем результат
}
