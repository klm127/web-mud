package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/server"
	"github.com/pwsdc/web-mud/server/workers"
	"github.com/pwsdc/web-mud/util/console"
)

func main() {
	green := console.GetFgSprintf(0, 245, 0)
	blue := console.GetFgSprintf(0, 0, 245)
	red := console.GetFgSprintf(245, 0, 0)
	fmt.Println(green("Starting sdcmud"))
	gin.SetMode(gin.ReleaseMode)
	arg.Parse()
	arg.Config.PrintLogs()

	server := server.CreateServer()
	fmt.Println(green("Connecting to database on port "+arg.Config.Db.Port()) + ".")
	err := db.Store.Connect()
	db.Store.PrintLogs()
	if err != nil {
		fmt.Println(red("Failed to connect to postgres! Error: %s", err.Error()))
		return
	}
	fmt.Println(green("HTTP Server listening on port "+arg.Config.Http.Port()) + ".")
	fmt.Println(blue("http://localhost:" + arg.Config.Http.Port()))

	workers.StartAllWorkers()
	err = server.Run("localhost:" + arg.Config.Http.Port())
	if err != nil {
		fmt.Println(red("Server exited. %s", err.Error()))
	}

}
