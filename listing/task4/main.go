package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch) закрываем канал после записи чисел, и деадлока не будет
	}()

	for n := range ch {
		println(n)
	}
}

/*
Ответ:

0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

Так как передача данных закончена, а канал не закрыли, происходит деад лок
*/
