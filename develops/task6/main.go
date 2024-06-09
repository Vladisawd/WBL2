package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Утилита cut

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN,
разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

type flags struct {
	fields    string
	delimiter string
	separated bool
}

func parseFlags() *flags {
	f := flags{}

	flag.StringVar(&f.fields, "f", "0", "select fields (columns)")
	flag.StringVar(&f.delimiter, "d", "\t", "use a different delimiter")
	flag.BoolVar(&f.separated, "s", false, "only delimited lines")
	flag.Parse()

	return &f
}

func cut(input string, f *flags) string {
	if f.separated && !strings.Contains(input, f.delimiter) {
		return ""
	}

	result := strings.Builder{}
	splitted := strings.Split(input, f.delimiter)
	columns := strings.Split(f.fields, ",")
	for i := 0; i < len(columns); i++ {
		column, err := strconv.Atoi(columns[i])
		if err != nil {
			log.Fatalln("cannot parse column to int: ", err.Error())
		}
		result.WriteString(splitted[column])
	}
	return result.String()
}

func main() {
	f := parseFlags()
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		fmt.Println("cut str: ", cut(text, f))
	}
}
