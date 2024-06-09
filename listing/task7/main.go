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
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

/*
Ответ:
Сначала мы получаем значения от 1 до 8 в случайном порядке, а потом бесконечно выводится 0
Так как в функции asChan мы закрыли наши каналы, то все корректно выводится, а в функции merge
мы поулчаем значения из каналов а и в, выводим их НО потом не закрываем канал и наш цикл бесконечно выводит
zero value для int то есть 0

*/
