package socket

import "github.com/gin-gonic/gin"

func LoadRoutes(topRouter *gin.Engine) {
	group := topRouter.Group("sock")
	group.GET("connect", handleWebSocket)
	group.POST("connect-http", attachActor(), extractActorMessages(), respondToActor)
}
