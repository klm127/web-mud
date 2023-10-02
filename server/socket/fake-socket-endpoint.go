package socket

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/token"
	"github.com/pwsdc/web-mud/server/user/actor"
)

const contextActorKey = "actor"

/*
This is the structure of data sent back to the client, who is using their fake HTTP socket endpoint. These parameters tell the client which registered callbacks to call.

EType can be "message", "error", "open", or "close". It will map to 'event' in the JSON object and trigger the corresponding handlers on the front end.
Message is any data contained within. It is mapped to 'data' in the json.
*/
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
		// check the cookies for a JWT token we can extract an ID from.
		id, err := token.ExtractID(c)
		// if we were succesful, get the actor.
		if err == nil {
			actor, ok := actor.GetActor(id)
			if ok {
				// if the actor exists, place a reference to the actor and finish.
				c.Set(contextActorKey, actor)
				c.Next()
				return
			}
		}
		// if we were unsuccesful finding our actor, we need to register a fake socket.
		sock := NewFakeSocket()
		act := actor.StartActor(sock)
		// The front end is expecting an 'open' event.
		sr := sock_response{"open", ""}
		sra := make([]sock_response, 1)
		sra[0] = sr
		token.WriteID(act.GetId(), c)
		c.JSON(200, sra)
		c.Abort()
		IpLogger.Logf("Created fake socket http connection for ip %s. Assigned id: %d", c.ClientIP(), act.GetId())
	}
}

/* Messages from the front-end's fake socket come in this format. */
type clientMsg struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// Extracts
func extractActorMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		r_val, ok := c.Get(contextActorKey)
		if !ok { // this should never trigger
			fmt.Println("Tried to extract an actor that didn't exist!")
			c.JSON(200, quickErrResponse("ERR-NO-ACTOR"))
			IpLogger.Errorf("Fake socket error - no actor in extractActorMessages. Ip: %s", c.ClientIP())
			return
		}
		actor := r_val.(iactor.IActor)
		fs := actor.GetConnection().(*fakeSocket)
		pf := make([]clientMsg, 0, 1)
		c.Bind(&pf) // extract the messages the client sent
		for _, v := range pf {
			if v.Type == "message" {
				fs.addMessage(v.Data) // add the messages for the actor loop to pick up.
			} else if v.Type == "close" {
				actor.Disconnect()
			}
		}
		c.Next()
	}
}

func respondToActor(c *gin.Context) {
	raw_val, ok := c.Get(contextActorKey)
	if !ok { // should never trigger
		c.JSON(200, quickErrResponse("ERR-NO-ACTOR"))
		IpLogger.Errorf("Fake socket error - no actor in respondToActor. Ip: %s", c.ClientIP())
		return
	}
	actor := raw_val.(iactor.IActor)
	fake_socket := actor.GetConnection().(*fakeSocket)
	if fake_socket.closed {
		respondClosed(c)
		return
	}
	to_send := fake_socket.popPending() // get messages waiting for the actor since last communication
	send_arr := make([]sock_response, len(to_send))
	for i, v := range to_send {
		send_arr[i].EType = "message"
		send_arr[i].Message = string(v)
	}
	c.JSON(200, send_arr)
}

func respondClosed(c *gin.Context) {
	send_arr := make([]sock_response, 1)
	cm := sock_response{"close", ""}
	send_arr[0] = cm
	c.JSON(200, send_arr)
}
