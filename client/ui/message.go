package ui

import (
	"encoding/json/v2"
	"fmt"
)

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