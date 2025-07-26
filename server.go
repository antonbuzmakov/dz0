package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("ошибка, сервер не запущен:", err)
		return
	}
	defer listener.Close()

	fmt.Println("сервер слушает порт :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ошибка подключения:", err)
			continue
		}

		conn.Write([]byte("OK\n"))
		conn.Close()
	}
}
