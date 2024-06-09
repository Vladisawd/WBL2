package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Создадим структуру флагов которые будут нам нужны
type flags struct {
	column         int
	byNumber       bool
	inReverseOrder bool
	noRepetitive   bool
}

// Функция определяющая по какому флагу будем сортировать файл
func parseFlags() *flags {
	f := flags{}

	flag.IntVar(&f.column, "k", -1, "specifying the column to sort")
	flag.BoolVar(&f.byNumber, "n", false, "sort by numeric value")
	flag.BoolVar(&f.inReverseOrder, "r", false, "sort in reverse order")
	flag.BoolVar(&f.noRepetitive, "u", false, "do not print duplicate lines")
	flag.Parse()

	return &f
}

// Функция определяющая как мы будем сортировать ориентируясь на флаги
func flagsSort(lines []string, f *flags) []string {
	var sortLines []string
	//Если подается значение столбика, то пользуемся функцией сортировки именно с указанием столбика
	//по умолчанию столбик -1, так что если не подается, то сортируем просто встроенной функцией
	if f.column >= 0 {
		sortLines = sortByColumn(lines, f.column)
	} else {
		sort.Strings(lines)
		sortLines = lines
	}

	if f.inReverseOrder {
		sortByReverseOrder(sortLines)
	}

	if f.noRepetitive {
		sortLines = sortByNoRepetitive(sortLines)
	}

	return sortLines
}

// Сортироска с указанием столбика
// Здесь я сначала разбиваю строку на отдельные элементы, чтобы взять те элементы, по которым мы будем сортировать
// Например взять именно вторые элементы, и занести их в качестве ключей мапы, а в качестве значений занести соответствующие
// им строки, таким образом мы получим отсортированный по нужным элементам массив строк
func sortByColumn(lines []string, col int) []string {
	m := make(map[string]string)

	for i := 0; i < len(lines); i++ {
		splittedLine := strings.Split(lines[i], " ")
		m[splittedLine[col]] = lines[i]
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedLines := make([]string, len(lines))
	for i, key := range keys {
		sortedLines[i] = m[key]
	}
	return sortedLines
}

// Сортировка в обратном порядке
func sortByReverseOrder(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// Сортировать не выводя повторяющиеся строки
func sortByNoRepetitive(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// Функция на чтение файла
func readFile(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error closing file", err.Error())
		}
	}(file)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

// Функция на запись файла
func writeFile(fileName string, lines []string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file", err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error closing file", err.Error())
		}
	}(file)

	for i := 0; i < len(lines); i++ {
		_, err := fmt.Fprintln(file, lines[i])
		if err != nil {
			log.Fatal("Error writing to the file", err.Error())
		}
	}
}

// Названия наших файлов
const (
	fileinput  = "input.txt"
	fileoutput = "output.txt"
)

// Проверяем
func main() {
	flags := parseFlags()
	lines := readFile(fileinput)
	sortedLines := flagsSort(lines, flags)
	writeFile(fileoutput, sortedLines)
}
