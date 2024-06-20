package pattern

//package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Это структурный паттерн проектирования, уровня объекта
Представляет собой высокоуровневый интерфейс(т.е. доступ) к сложной системе типов

Разбиение системы на подсистемы позволяет упростить процесс разработки,
а также помогает максимально снизить зависимости одной подсистемы от другой.
Однако использовать такие подсистемы становиться довольно сложно.
Один из способов решения этой проблемы является паттерн Facade.
Задача, сделать единый объект, через который можно было бы взаимодействовать с подсистемами.

+ изолирует клиентов от поведения сложной подсистемы давая доступ к высокоуровневому интерфейсу
- сам интерфейс фасада может стать суперобъектом
*/

// import (
// 	"fmt"
// )

type Sumator struct {
	A, B float32
}

func (s *Sumator) Sum() float32 {
	return s.A + s.B
}

type Substractor struct {
	A, B float32
}

func (s *Substractor) Sub() float32 {
	return s.A - s.B
}

type Multiplyer struct {
	A, B float32
}

func (s *Multiplyer) Multi() float32 {
	return s.A * s.B
}

type Devider struct {
	A, B float32
}

func (s *Devider) Devide() float32 {
	return s.A / s.B
}

type Calculator struct {
	Sm *Sumator
	Sb *Substractor
	M  *Multiplyer
	D  *Devider
}

func (c *Calculator) Sum() float32 {
	return c.Sm.Sum()
}

func (c *Calculator) Sub() float32 {
	return c.Sb.Sub()
}

func (c *Calculator) Multi() float32 {
	return c.M.Multi()
}

func (c *Calculator) Devide() float32 {
	return c.D.Devide()
}

func NewCalculator(a, b float32) *Calculator {
	return &Calculator{
		Sm: &Sumator{a, b},
		Sb: &Substractor{a, b},
		M:  &Multiplyer{a, b},
		D:  &Devider{a, b},
	}
}

// func main() {
// 	calc := NewCalculator(10, 20)
// 	fmt.Printf("\nCalculated sum: %0.2f\n", calc.Sum())
// 	fmt.Printf("\nCalculated sub: %0.2f\n", calc.Sub())
// 	fmt.Printf("\nCalculated multi: %0.2f\n", calc.Multi())
// 	fmt.Printf("\nCalculated Dev: %0.2f\n", calc.Devide())
// }
