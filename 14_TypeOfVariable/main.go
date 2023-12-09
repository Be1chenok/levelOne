package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func typeOfVariable(val interface{}) string {
	// Использую type switch

	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan int32:
		return "chan int32"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return "unknown type"
	}

}

func typeOfVariable1(val interface{}) reflect.Type {
	/*
		reflect.TypeOf возвращает экземпляр типа reflect.Type
		reflect.Type содержит информацию о типе данных.
	*/
	return reflect.TypeOf(val)
}

func main() {
	var a int
	var b bool
	var s string
	ch := make(chan rune)

	fmt.Printf("a is %v\n", typeOfVariable(a))
	fmt.Printf("b is %v\n", typeOfVariable(b))
	fmt.Printf("s is %v\n", typeOfVariable(s))
	fmt.Printf("ch is %v\n", typeOfVariable(ch))

	fmt.Printf("a is %v\n", typeOfVariable1(a))
	fmt.Printf("b is %v\n", typeOfVariable1(b))
	fmt.Printf("s is %v\n", typeOfVariable1(s))
	fmt.Printf("ch is %v\n", typeOfVariable1(ch))
}
