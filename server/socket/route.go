package socket

import "github.com/gin-gonic/gin"

func LoadRoutes(topRouter *gin.Engine) {
	group := topRouter.Group("sock")
	group.GET("connect", handleWebSocket)
	p_http := group.POST("connect-http", respondToActor)
	p_http.Use(attachActor(), extractActorMessages())
}
