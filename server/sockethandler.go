package main

import (
	"fmt"
	"lnw/intralot/socket"

	"github.com/gorilla/websocket"
)

func socketHandler(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err,
				websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseNoStatusReceived) {
				fmt.Println("[Socket Handler] Connection closed by client (no status)")
			} else {
				fmt.Println("[Socket Handler] Read error:", err)
			}
			return
		}
		socket.HandleIncomingMessages(message)
	}
}
