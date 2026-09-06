package main

import (
	"bufio"
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

	go func() {
		serverScanner := bufio.NewScanner(serverConnection)

		for serverScanner.Scan() {
			fmt.Println(serverScanner.Text())
		}
	}()

	scannerInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a sua mensagem:")

	for {
		if !scannerInput.Scan() {
			break
		}

		message := scannerInput.Text()

		_, err := serverConnection.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error ao enviar a mensagem")
			break
		}
	}
}