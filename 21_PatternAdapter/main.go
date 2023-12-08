package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// Интерфейс современной бд
type ModernDatabase interface {
	Add(value string) bool
}

// Интерфейс старой бд
type OldDatabase interface {
	Add(value string) int
}

// Структура старой бд
type oldDatabase struct{}

// Конструктор старой бд
func NewOldDatabase() OldDatabase {
	return &oldDatabase{}
}

// Метод старой бд
func (d oldDatabase) Add(value string) int {
	return 1
}

// Структура современной бд
type modernDatabase struct{}

// Конструктор современной бд
func NewModernDatabase() ModernDatabase {
	return &modernDatabase{}
}

// Метод современной бд
func (d modernDatabase) Add(value string) bool {
	return true
}

// Структура адаптера
type databaseAdapter struct {
	oldDatabase OldDatabase
}

// Метод адаптера
func (dba databaseAdapter) Add(value string) bool {
	// Используем старую бд адаптируя под новый интерфейс
	result := dba.oldDatabase.Add(value)
	return result == 1
}

func main() {
	// Создаем экземпляр modernDatabase
	modernDB := NewModernDatabase()
	// Создаем экземпляр oldDatabase
	oldDB := NewOldDatabase()
	// Адаптируем
	adapter := &databaseAdapter{
		oldDatabase: oldDB,
	}

	fmt.Printf("modernDB result: %t\n", modernDB.Add("value"))           // Вывод результата работы современной бд
	fmt.Printf("oldDB result without adapter: %d\n", oldDB.Add("value")) // Вывод результата работы старой бд без адаптера
	fmt.Printf("oldDB result with adapter\n: %t", adapter.Add("value"))  // Вывод результата работы старой бд с адаптером

}
