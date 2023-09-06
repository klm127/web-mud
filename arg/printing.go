package arg

import "fmt"

func configPanic(flag string, env string) {
	s := fmt.Sprintf("Error configuring. Set the flag '%s' or the environment variable '%s'.", flag, env)
	panic(s)
}

func configWarn(envvar string, using any) {
	s := fmt.Sprintf("Warning, environmment variable %s not set. Using value of : %s", envvar, using)
	fmt.Println(s)
}
