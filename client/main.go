package main

import (
	"bufio"
	"client/ui"
	"encoding/json/v2"
	"fmt"
	"net"
	"os"

	tea "charm.land/bubbletea/v2"
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

	model := ui.InitModel(serverConnection, user)
	p := tea.NewProgram(model)

	go func() {
		serverScanner := bufio.NewScanner(serverConnection)

		for serverScanner.Scan() {
			serverText := serverScanner.Text()
			serverMessage, _ := ui.FromJsonString(serverText)
			p.Send(serverMessage)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error ao inicializar a UI")
		return
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