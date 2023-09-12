package socket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/server/user/actor"
)

type serverMsg struct {
	error bool
	msg   string
}

func serverError(msg string) serverMsg {
	return serverMsg{true, msg}
}
func serverGood(msg string) serverMsg {
	return serverMsg{false, msg}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusFound, serverError("Couldn't upgrade web socket: "+err.Error()))
		return
	}
	actor.StartActor(conn)
}
