package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Поиск анаграмм по словарю

Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/

// Сама функция
func anagramSearch(words *[]string) *map[string][]string {
	result := make(map[string][]string)
	toLower(*words)      //Сначала переводим все буквы слов в нихний регистр
	sort.Strings(*words) //Сортируем все слова

	for _, word := range *words {
		flag := false //Флаг нужен для того, чтобы точно был ключ в мапе
		for key, val := range result {
			if anagram(key, word) {
				result[key] = append(val, word)
				flag = true
				break
			}
		}
		if !flag {
			result[word] = append(result[word], word)
		}
	}

	//Если будет одно слово, удаляем ключ
	for k, v := range result {
		if len(v) <= 1 {
			delete(result, k)
		}
	}

	return &result
}

func toLower(words []string) {
	for i, x := range words {
		words[i] = strings.ToLower(x)
	}
}

// анаграмма - это слово в котором есть определенное количество букв которое не изменяется, но переставляется в другом порядке
// таким образом можно просто отсортировать все имеющиеся буквы и сравнить два слова
func anagram(key, word string) bool {
	keyChars := strings.Split(key, "")
	sort.Strings(keyChars)
	resultKey := strings.Join(keyChars, "")

	wordChars := strings.Split(word, "")
	sort.Strings(wordChars)
	resultWord := strings.Join(wordChars, "")

	return resultKey == resultWord
}

// проверяем работу
func main() {
	words := []string{"Тяпка", "пятак", "лИсток", "столик", "пятка", "слиток"}
	m := anagramSearch(&words)
	fmt.Println(m)
}
