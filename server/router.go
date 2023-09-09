package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/pwsdc/web-mud/server/actor/commands"
	"github.com/pwsdc/web-mud/server/socket"
)

func CreateServer() *gin.Engine {
	serv := gin.Default()
	serv.Static("static", "server/static")
	serv.LoadHTMLGlob("server/templates/**/*")
	serv.GET("/", home)
	socket.LoadRoutes(serv)
	return serv
}
