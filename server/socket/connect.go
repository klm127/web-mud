package socket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/server/msg"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusFound, msg.ServerError("Couldn't upgrade web socket: "+err.Error()))
		return
	}
	base.CreateActor(conn)

}
