package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

/*
Ответ:

<nil>
false

Нил потому что передает как нил, а Фолс, потому что тип не соответствует
Тип err в мейне error, а получаемый при отрабатывании функции *fs.PathError
*/
