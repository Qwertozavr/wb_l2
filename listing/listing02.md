Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

package main

import ( "fmt" )

func test() (x int) { defer func() { x++ }() x = 1 return }

func anotherTest() int { var x int defer func() { x++ }() x = 1 return x }

func main() { fmt.Println(test()) fmt.Println(anotherTest()) }

Ответ: 2 1

Блок defer выполняется перед выходом из функции, если несколько вызовов то в последовательности LIFO. В первом случае функция возвращает именно переменную х, после возвращения выполнится конструкция defer и х будет равен 2. Во втором случае defer так же будет выполнен и х будет увеличен, но функция вернёт просто значение int в конкретный момент вызова return, а не саму переменную х.