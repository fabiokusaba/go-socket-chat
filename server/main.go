package main

import (
	"bufio"
	"encoding/json/v2"
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

		message, err := FromJsonString(clientMessage)
		if err != nil {
			return
		}

		fmt.Printf("Mensagem recebida [%s] - Mensagem [%s]\n", message.SenderName, message.MessageText)

		_, err = connection.Write([]byte(message.ToJsonString()))
		if err != nil {
			fmt.Println("Error ao enviar a mensagem para o cliente")
			break
		}
	}
}

type Message struct {
	SenderName string
	MessageText string
}

func FromJsonString(data string) (Message, error) {
	var message Message

	if len(data) <= 0 {
		return Message{}, fmt.Errorf("invalid data")
	}

	err := json.Unmarshal([]byte(data), &message)
	if err != nil {
		return Message{}, fmt.Errorf("invalid data")
	}

	return message, nil
}

func(m Message) ToJsonString() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s\n", string(data))
}