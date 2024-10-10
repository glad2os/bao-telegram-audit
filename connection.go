package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Ошибка при закрытии соединения:", err)
		}
	}(conn)

	fmt.Println("Новое соединение от:", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		processData(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
	}
}
