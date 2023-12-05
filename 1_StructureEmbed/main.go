package main

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

import "fmt"

// Определяем структуру Human с произвольным набором полей и методов
type Human struct {
	Name    string
	Surname string
	Age     int
}

// Метод структуры Human
func (h Human) Eat() {
	fmt.Printf("%s is eating\n", h.Name)
}

// Метод структуры Human
func (h Human) Drink() {
	fmt.Printf("%s is drinking\n", h.Name)
}

// Определяем структуру Action с встраиванием структуры Human
type Action struct {
	Human      // Встраиваем структуру Human
	Profession string
}

/*
		Встраивая одну структуру в другую, внешняя структура получает
	доступ ко всем полям и методам внутренней(встроенной) структуры,
	но без использования имени этой встроенной структуры.
		Внешняя структура может обращаться к полям и методам внутренней
	структуры так, как если бы они были определенны прямо внутри этой
	внешней структуры.
*/

// Метод структуры Action
func (a Action) Work() {
	fmt.Printf("%s works as a %s\n", a.Name, a.Profession)
}

func main() {
	// Создаем экземпляр типа Action
	act := Action{
		Human: Human{
			Name:    "John",
			Surname: "Biller",
			Age:     20,
		},
		Profession: "Builder",
	}

	act.Eat()   // Вызываем метод встроенной структуры Human
	act.Drink() // Вызываем метод встроенной структуры Human
	act.Work()  // Вызываем метод структуры Action
}
