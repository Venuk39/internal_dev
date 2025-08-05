package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/intralot", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("[Socket] Upgrading Error: ", err)
			return
		}
		fmt.Println("[Socket] New connection established from:", r.RemoteAddr)
		defer conn.Close()
		socketHandler(conn)
	})

	fmt.Println("WebSocket server started on :8005")
	err := http.ListenAndServe(":8005", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
