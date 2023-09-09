package arg

import (
	"fmt"

	"github.com/pwsdc/web-mud/util/console"
)

var SpYellow func(string, ...any) string
var SpRed func(string, ...any) string

func init() {
	SpRed = console.GetFgSprintf(250, 0, 0)
	SpYellow = console.GetFgSprintf(255, 255, 0)
}

func configPanic(flag string, env string) {
	s := SpRed("Error configuring. Set the flag '%s' or the environment variable '%s'.", flag, env)
	panic(s)
}

func configWarn(envvar string, using any) {
	s := SpYellow("Warning, environmment variable %s not set. Using value of : %v", envvar, using)
	fmt.Println(s)
}
