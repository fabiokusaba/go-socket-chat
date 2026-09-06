package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error ao subir o servidor")
		return
	}

	defer listener.Close()

	fmt.Println("Servidor rodando na porta :8080")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error ao aceitar a conexão com o servidor")
			return
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	fmt.Println("Novo cliente conectado")

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		clientMessage := scanner.Text()
		fmt.Println("Mensagem recebida:" + clientMessage)

		response := fmt.Sprintf("Servidor: mensagem %s recebida com sucesso!", clientMessage)

		_, err := connection.Write([]byte(response + "\n"))
		if err != nil {
			fmt.Println("Error ao enviar a mensagem para o cliente")
			break
		}
	}
}