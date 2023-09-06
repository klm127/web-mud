package arg

import (
	"flag"
	"fmt"
	"os"

	"github.com/pwsdc/web-mud/shared"
)

// Database connection settings
type dbConfig struct {
	shared.HasLogs
	host     *string `env:"db_container"`
	port     *string `env:"db_port"`
	user     *string `env:"db_user"`
	password *string `env:"db_pw"`
	db       *string `env:"db"`
}

func (self *dbConfig) setFlags() {
	shared.HasLogsInit(self)
	self.host = flag.String("db_container", "localhost", "Set the host for the database connection.")
	self.port = flag.String("db_port", "5432", "Set the port for the database connection.")
	self.user = flag.String("db_user", "postgres", "Set the user for the database connection.")
	self.password = flag.String("db_pw", "password", "Set the password for the database connection.")
	self.db = flag.String("db", "bwgdb", "Set the name of the database for the database connection.")
}

// checks the environment variables when no command-line argument was given
func (self *dbConfig) parseEnv() {

	os_host := os.Getenv("db_container")
	if os_host == "" {
		configWarn("db_container", *self.host)
	} else {
		*self.host = os_host
	}
	self.Log("Host set to: " + *self.host + " from environment.")

	os_port := os.Getenv("db_port")
	if os_host == "" {
		configWarn("db_port", *self.port)
	} else {
		*self.port = os_port
	}
	self.Log("Port set to: " + *self.port + " from environment.")

	os_user := os.Getenv("db_user")
	if os_host == "" {
		configWarn("db_user", *self.user)
	} else {
		*self.user = os_user
	}
	self.Log("User set to: " + *self.user + " from environment.")

	os_pw := os.Getenv("db_pw")
	if os_pw == "" {
		configWarn("db_pw", *self.password)
	} else {
		*self.password = os_pw
	}
	self.Log("Password set to: ********" + " from environment.")

	os_db := os.Getenv("db")
	if os_db == "" {
		configWarn("db", *self.db)
	} else {
		*self.db = os_db
	}
	self.Log("Database set to: " + *self.db + " from environment.")
}

func (self *dbConfig) Host() string {
	return *self.host
}
func (self *dbConfig) Port() string {
	return *self.port
}
func (self *dbConfig) User() string {
	return *self.user
}
func (self *dbConfig) Password() string {
	return *self.password
}
func (self *dbConfig) Database() string {
	return *self.db
}

// gets a connection string for postgres
func (self *dbConfig) ConnectString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *self.user, *self.password, *self.host, *self.port, *self.db)
}

func (self *dbConfig) ConnectStringLocalhost() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *self.user, *self.password, "localhost", *self.port, *self.db)
}
