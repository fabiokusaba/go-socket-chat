package main

import (
	"bufio"
	"encoding/json/v2"
	"fmt"
	"net"
	"os"
)

func main() {
	serverConnection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error ao se conectar com o servidor")
		return
	}

	defer serverConnection.Close()

	fmt.Println("Escreva o seu nome para entrar no chat")
	var user string
	usernameInput := bufio.NewScanner(os.Stdin)
	if usernameInput.Scan() {
		user = usernameInput.Text()
	}

	go func() {
		serverScanner := bufio.NewScanner(serverConnection)

		for serverScanner.Scan() {
			serverText := serverScanner.Text()

			serverMessage, _ := FromJsonString(serverText)
			fmt.Printf("[%s]: %s\n", serverMessage.SenderName, serverMessage.MessageText)
		}
	}()

	scannerInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a sua mensagem:")

	for {
		if !scannerInput.Scan() {
			break
		}

		msgText := scannerInput.Text()

		message := Message{
			SenderName: user,
			MessageText: msgText,
		}

		_, err := serverConnection.Write([]byte(message.ToJsonString()))
		if err != nil {
			fmt.Println("Error ao enviar a mensagem")
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