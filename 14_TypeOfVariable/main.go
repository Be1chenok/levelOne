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
	/*
		reflect.TypeOf возвращает экземпляр типа reflect.Type
		reflect.Type содержит информацию о типе данных.
	*/
	return fmt.Sprintln(reflect.TypeOf(val))
}

func main() {
	var a int
	var b bool
	var c float32
	var s string
	ch := make(chan rune)

	fmt.Printf("a is %v", typeOfVariable(a))
	fmt.Printf("b is %v", typeOfVariable(b))
	fmt.Printf("c is %v", typeOfVariable(c))
	fmt.Printf("s is %v", typeOfVariable(s))
	fmt.Printf("ch is %v", typeOfVariable(ch))
}
