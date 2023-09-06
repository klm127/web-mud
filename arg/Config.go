package arg

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

type argConfig struct {
	Db   dbConfig
	Http httpConfig
}

// Global configuration - set from flags or env on program start
var Config argConfig

// parses environment file and command line variables
func Parse() {

	Config = argConfig{}
	Config.Db.setFlags()
	Config.Http.setFlags()
	flag.Parse()

	env_path := flag.Arg(0)
	if len(env_path) != 0 {
		err := godotenv.Load(env_path)
		if err != nil {
			fmt.Println("Failed to load environment file at: ", env_path)
			panic(err.Error())
		}
		Config.Db.parseEnv()
		Config.Http.parseEnv()
	}
}

func (self *argConfig) PrintLogs() {
	dbl := *self.Db.GetLogs()
	for _, v := range dbl {
		fmt.Printf("DB Config: %v\n", v)
	}
	httpl := *self.Http.GetLogs()
	for _, v := range httpl {
		fmt.Printf("HTTP Config: %v\n", v)
	}
}
