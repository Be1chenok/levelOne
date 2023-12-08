package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

type BigInteger interface {
	Add() *big.Int
	Sub() *big.Int
	Div() *big.Int
	Mul() *big.Int
}

type bigInteger struct {
	a, b *big.Int
}

/*
Конструктор который создает новый экземпляр bigInteger и возвращает его в виде BigInteger(интерфейс)
это позволяет использовать интерфейс BigInteger для взаимодействия,
скрывая детали реализации
*/
func NewBigInteger(a, b int64) BigInteger {
	return &bigInteger{
		a: new(big.Int).SetInt64(a),
		b: new(big.Int).SetInt64(b),
	}
}

// Суммирование
func (b bigInteger) Add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

// Вычитание
func (b bigInteger) Sub() *big.Int {
	return new(big.Int).Sub(b.a, b.b)
}

// Деление
func (b bigInteger) Div() *big.Int {
	return new(big.Int).Div(b.a, b.b)
}

// Умножение
func (b bigInteger) Mul() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

func main() {
	nums := NewBigInteger(999999999999999999, 222222222222222222)

	add := nums.Add()
	fmt.Printf("add: %d\n", add)

	sub := nums.Sub()
	fmt.Printf("sub: %d\n", sub)

	div := nums.Div()
	fmt.Printf("div: %d\n", div)

	mul := nums.Mul()
	fmt.Printf("mul: %d\n", mul)
}
