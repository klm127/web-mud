package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/server"
	admincommands "github.com/pwsdc/web-mud/server/admin-commands"
	"github.com/pwsdc/web-mud/server/user/actor"
	basecommands "github.com/pwsdc/web-mud/server/user/base-commands"
	"github.com/pwsdc/web-mud/server/workers"
	"github.com/pwsdc/web-mud/util/console"
	buildercommands "github.com/pwsdc/web-mud/world/builder-commands"
)

func main() {
	green := console.GetFgSprintf(0, 245, 0)
	blue := console.GetFgSprintf(0, 0, 245)
	red := console.GetFgSprintf(245, 0, 0)
	fmt.Println(green("Starting sdcmud"))
	gin.SetMode(gin.ReleaseMode)

	// Parse the command line arguments
	arg.Parse()
	arg.Config.PrintLogs()

	// Create the server and connect to the database.
	server := server.CreateServer()
	fmt.Println(green("Connecting to database on port "+arg.Config.Db.Port()) + ".")
	err := db.Store.Connect()
	db.Store.PrintLogs()
	if err != nil {
		fmt.Println(red("Failed to connect to postgres! Error: %s", err.Error()))
		return
	}

	// Ensure the start room exists
	db.Store.EnsureStartRoom()

	// Set the default command groups for a logged-out connection
	actor.SetDefaultCommandGroups(basecommands.UserLoggedOutCommands, basecommands.UserInfoCommands, basecommands.ConnectionCommands, admincommands.AdminCommands, buildercommands.BuilderCommands)
	// Remember to remove admin commands from above

	// Start all the background workers.
	workers.StartAllWorkers()

	// Start the server
	fmt.Println(green("HTTP Server listening on port "+arg.Config.Http.Port()) + ".")
	fmt.Println(blue("http://localhost:" + arg.Config.Http.Port()))
	err = server.Run("localhost:" + arg.Config.Http.Port())
	if err != nil {
		fmt.Println(red("Server exited. %s", err.Error()))
	}
}
