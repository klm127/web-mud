package server

import "github.com/gin-gonic/gin"

func CreateServer() *gin.Engine {
	serv := gin.Default()
	serv.Static("static", "server/static")
	serv.LoadHTMLGlob("server/templates/**/*")
	serv.GET("/", home)
	return serv
}
