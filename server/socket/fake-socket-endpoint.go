package socket

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/token"
	"github.com/pwsdc/web-mud/server/user/actor"
)

const contextActorKey = "actor"

type sock_response struct {
	EType   string `json:"event"`
	Message string `json:"data"`
}

func quickErrResponse(s string) []sock_response {
	sr := sock_response{EType: "error", Message: s}
	sa := make([]sock_response, 0, 1)
	sa[0] = sr
	return sa
}

// Reads the JWT token to attach the initialized actor. Otherwise, starts a new one.
func attachActor() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := token.ExtractID(c)
		if err == nil {
			actor, ok := actor.GetActor(id)
			if ok {
				c.Set(contextActorKey, actor)
				c.Next()
				return
			}
		}
		sock := NewFakeSocket()
		act := actor.StartActor(sock)
		sr := sock_response{"open", ""}
		sra := make([]sock_response, 1)
		sra[0] = sr
		token.WriteID(act.GetId(), c)
		c.JSON(200, sra)
		c.Abort()

		// c.Set(contextActorKey, actor)
		// c.Next()
	}
}

/* A struct to extract data sent from user */
type clientMsg struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func extractActorMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		r_val, ok := c.Get(contextActorKey)
		if !ok {
			fmt.Println("Tried to extract an actor that didn't exist!")
			c.JSON(200, quickErrResponse("ERR-NO-ACTOR"))
			return
		}
		actor := r_val.(iactor.IActor)
		fs := actor.GetConnection().(*fakeSocket)
		pf := make([]clientMsg, 0, 1)
		c.Bind(&pf)
		for _, v := range pf {
			if v.Type == "message" {
				fmt.Println("adding", v.Type, "to received")
				fs.addMessage(v.Data)
			}
			// open, close events?
		}
		c.Next()
	}
}

func respondToActor(c *gin.Context) {
	raw_val, ok := c.Get(contextActorKey)
	if !ok {
		c.JSON(200, quickErrResponse("ERR-NO-ACTOR"))
		return
	}
	actor := raw_val.(iactor.IActor)
	fake_socket := actor.GetConnection().(*fakeSocket)
	to_send := fake_socket.popPending()
	send_arr := make([]sock_response, len(to_send))
	for i, v := range to_send {
		send_arr[i].EType = "message"
		send_arr[i].Message = string(v)
	}
	c.JSON(200, send_arr)
}
