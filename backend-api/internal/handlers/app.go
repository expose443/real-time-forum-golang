package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *Client) WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.logger.Error.Println("Ошибка создания соединения:", err)
		return
	}
	for {
		_, message, err := conn.ReadMessage()
		fmt.Println(r.Cookie("jwt_token"))
		if err != nil {
			app.logger.Error.Println("Ошибка чтения сообщения:", err)
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("hello from server %s ", message)))
		fmt.Println(string(message))
	}
}
