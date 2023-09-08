package socket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/server/msg"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var sockets []*websocket.Conn

func init() {
	sockets = make([]*websocket.Conn, 0, 100)
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusFound, msg.ServerError("Couldn't upgrade web socket: "+err.Error()))
		return
	}
	sockets = append(sockets, conn)
	conn.WriteMessage(1, []byte("Hello, Socket!"))
	getSocketMessages(conn)
	conn.Close()

}

func getSocketMessages(conn *websocket.Conn) {
	close_requested := false
	for !close_requested {
		mtype, data, err := conn.ReadMessage()
		if err != nil {
			break
		}

		if mtype != 1 {
			conn.WriteMessage(1, []byte("I couldn't understand that message."))
			continue
		}

		response := "I didn't understand that. Try 'list' to see available commands."
		command := string(data)

		switch command {
		case "list":
			response = "Commands are list, conns, quit, underly, subprotocol, remote"
		case "conns":
			response = fmt.Sprintf("There are %d connections", len(sockets))
		case "quit":
			close_requested = true
			response = "Disconnecting."
		case "underly":
			response = conn.UnderlyingConn().LocalAddr().String()
		case "subprotocol":
			response = conn.Subprotocol()
		case "remote":
			response = conn.UnderlyingConn().RemoteAddr().String()
		}
		conn.WriteMessage(1, []byte(response))
	}
	_ = close_requested
}
