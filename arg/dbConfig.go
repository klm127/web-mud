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

func (dbc *dbConfig) setFlags() {
	shared.HasLogsInit(dbc)
	dbc.host = flag.String("db_container", "localhost", "Set the host for the database connection.")
	dbc.port = flag.String("db_port", "5432", "Set the port for the database connection.")
	dbc.user = flag.String("db_user", "postgres", "Set the user for the database connection.")
	dbc.password = flag.String("db_pw", "password", "Set the password for the database connection.")
	dbc.db = flag.String("db", "bwgdb", "Set the name of the database for the database connection.")
}

// checks the environment variables when no command-line argument was given
func (dbc *dbConfig) parseEnv() {

	os_host := os.Getenv("db_container")
	if os_host == "" {
		configWarn("db_container", *dbc.host)
	} else {
		*dbc.host = os_host
	}
	dbc.Log("Host set to: " + *dbc.host + " from environment.")

	os_port := os.Getenv("db_port")
	if os_host == "" {
		configWarn("db_port", *dbc.port)
	} else {
		*dbc.port = os_port
	}
	dbc.Log("Port set to: " + *dbc.port + " from environment.")

	os_user := os.Getenv("db_user")
	if os_host == "" {
		configWarn("db_user", *dbc.user)
	} else {
		*dbc.user = os_user
	}
	dbc.Log("User set to: " + *dbc.user + " from environment.")

	os_pw := os.Getenv("db_pw")
	if os_pw == "" {
		configWarn("db_pw", *dbc.password)
	} else {
		*dbc.password = os_pw
	}
	dbc.Log("Password set to: ********" + " from environment.")

	os_db := os.Getenv("db")
	if os_db == "" {
		configWarn("db", *dbc.db)
	} else {
		*dbc.db = os_db
	}
	dbc.Log("Database set to: " + *dbc.db + " from environment.")
}

func (dbc *dbConfig) Host() string {
	return *dbc.host
}
func (dbc *dbConfig) Port() string {
	return *dbc.port
}
func (dbc *dbConfig) User() string {
	return *dbc.user
}
func (dbc *dbConfig) Password() string {
	return *dbc.password
}
func (dbc *dbConfig) Database() string {
	return *dbc.db
}

// gets a connection string for postgres
func (dbc *dbConfig) ConnectString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *dbc.user, *dbc.password, *dbc.host, *dbc.port, *dbc.db)
}

func (dbc *dbConfig) ConnectStringLocalhost() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *dbc.user, *dbc.password, "localhost", *dbc.port, *dbc.db)
}
