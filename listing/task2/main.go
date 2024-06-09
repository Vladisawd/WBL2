package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

/*
Ответ:

2
1

Первое значение будет 2, так как в return не уканаз точно возвращаемый тип,
из за этого функция выполняется и дефер тоже выполняется до возвращение икса

Второе значение будет 1, так как дефер сохраняет функция с сотсоянием x=0б и после выхода из функции только инкрементирует его
*/
