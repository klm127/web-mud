package socket

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/token"
	"github.com/pwsdc/web-mud/server/user/actor"
)

const contextActorKey = "actor"

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
		actor := actor.StartActor(sock)
		c.Set(contextActorKey, actor)
		c.Next()
	}
}

func extractActorMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		r_val, ok := c.Get(contextActorKey)
		if !ok {
			fmt.Println("Tried to extract an actor that didn't exist!")
			c.Writer.Write([]byte("ERROR-NO ACTOR"))
			return
		}
		/*
			Read the body of the HTTP request
			Send it to the actor
		*/
		actor := r_val.(iactor.IActor)
		_ = actor
		fmt.Println("Extracted actor!")
		c.Next()
	}
}

func respondToActor(c *gin.Context) {
	/*
		get actor messages
		write it to the HTTP response
	*/
}
