package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

/*
Ответ:

[77 78 79]
[1:4] создает срез, включающий элементы с 1 по 4, но 4 не включительно
*/
