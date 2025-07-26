package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("ошибка чтения ответа:", err)
		return
	}

	if response == "OK\n" {
		fmt.Println(" подключение корректно (OK\\n)")
	} else {
		fmt.Printf(" нет подключения: %q\n", response)
	}
}
