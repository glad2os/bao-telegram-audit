package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Ошибка при создании сервера:", err)
		os.Exit(1)
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("Ошибка при закрытии сервера:", err)
		}
	}(listen)

	fmt.Println("Сервер запущен, ожидание соединений...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении клиента:", err)
			continue
		}

		go handleConnection(conn)
	}
}
