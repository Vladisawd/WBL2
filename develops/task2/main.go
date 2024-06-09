package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

// Фенкция распаковки строки
func stringUnpacking(str string) (string, error) {
	//проверяем на пустоту и присутствие только чисел
	if len(str) == 0 || unicode.IsDigit(rune(str[0])) {
		return "", errors.New("некорректная строка")
	}
	//Инициализируем стринг билдер для возвращаемой строки итоговой
	totalLine := strings.Builder{}
	//проходимся по получаемой строке
	for i, symbol := range str {
		//проверяем явзяется ли элемент который мы сейчас просматриваем в строке цифрой
		number, err := strconv.Atoi(string(symbol))
		//если нет, то просто записываем его в итоговую строку
		if err != nil {
			totalLine.WriteString(string(symbol))
		} else {
			//если все таки число, то вписываем элемент который у нас был до этого числа столько раз, какое число мы получили
			for iS := 1; iS < number; iS++ {
				totalLine.WriteString(string(str[i-1]))
			}
		}
	}
	//возвращаем строку
	return totalLine.String(), nil
}

func main() {
	//проверяем работу
	unpStr, err := stringUnpacking("")
	if err != nil {
		log.Print(err.Error())
		return
	}
	fmt.Println(unpStr)
}
