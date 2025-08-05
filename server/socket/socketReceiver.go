package socket

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func HandleIncomingMessages(msg []byte) {
	var message Message
	err := json.Unmarshal(msg, &message)

	if err != nil {
		fmt.Println("[Incoming] Message Unmarshal Error: ", err)
		return
	}
	switch message.Type {
	case "register":
		fmt.Println("[Incoming] Register Req")
	}
}
