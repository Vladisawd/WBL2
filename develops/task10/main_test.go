package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main_t() {

	// Подключаемся к сокету
	//conn, _ := net.Dial("tcp", "127.0.0.1:2000")
	l, _ := net.Listen("tcp", "127.0.0.1:2000")
	conn, _ := l.Accept()
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
