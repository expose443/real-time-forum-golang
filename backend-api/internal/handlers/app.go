package handlers

import (
	"fmt"
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/jwt"
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
	c, err := r.Cookie("jwt_token")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.logger.Error.Println("Ошибка создания соединения:", err)
		return
	}
	for {
		status, datauser, err := jwt.VerifyJWT(c.Value)
		if !status || err != nil {
			f := conn.CloseHandler()
			err = f(401, "test")
			if err != nil {
				fmt.Println(err)
			}

			return
		}
		_, message, err := conn.ReadMessage()
		if err != nil {
			app.logger.Error.Println("Ошибка чтения сообщения:", err)
			break
		}
		idStr, ok := datauser["sub"]
		if !ok {
			app.logger.Error.Println("exp dont exist")
			conn.Close()
		}
		user, err := app.authService.GetUserById(idStr)
		if err != nil {
			app.logger.Error.Println(err)
		}
		fmt.Println(user)

		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("hello from server %s ", message)))
		app.logger.Debug.Println(string(message))
	}
}
