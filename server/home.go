package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {

	c.HTML(http.StatusFound, "console.html", nil)
}
