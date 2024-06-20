Что выведет программа? Объяснить вывод программы.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func asChan(vs ...int) <-chan int {
    c := make(chan int)

    go func() {
        for _, v := range vs {
            c <- v
            time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        }

        close(c)
    }()
    return c
}

func merge(a, b <-chan int) <-chan int {
    c := make(chan int) 
    go func() {
        for {
            select {
                case v := <-a:
                    c <- v
                case v := <-b:
                    c <- v 
            } 
        }
    }()
    return c
}

func main() {
    a := asChan(1, 3, 5, 7)
    b := asChan(2, 4 ,6, 8)
    c := merge(a, b )
    for v := range c {
        fmt.Println(v)
    }
}

Ответ: Выведутся через '\n' числа от 1 до 8 , а затем пойдут бесконечные 0

При вызове функции asChan в параллельной горутине происходит отправка в канал данных и возвращение этого канала При вызове функции merge в параллельной горутине происходит чтение из каналов Когда цикл в функции asChan завершается, функция закрывает канал Когда оба канала закрываются, в функции merge конструкция for select начинает бесконечно читать zero value из закрытых каналов.