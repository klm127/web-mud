package arg

import (
	"flag"
	"fmt"
	"os"

	"github.com/pwsdc/web-mud/shared"
)

// HTTP Server settings
type httpConfig struct {
	shared.HasLogs
	port    *string `env:"http_port"`
	docker  *bool   `env:"http_docker"`
	jwt_key *string `env:"jwt_key"`
	cookie  *string `env:"auth_cookie"`
}

func (hc *httpConfig) setFlags() {
	shared.HasLogsInit(hc)
	hc.port = flag.String("http_port", "80", "Set the port for the http server to listen on.")
	hc.docker = flag.Bool("http_docker", false, "Configures whether the http server is in a docker container using a docker network or should use localhost.")
	hc.jwt_key = flag.String("jwt_key", "default-key", "Set the encrpytion key for json web tokens.")
	hc.cookie = flag.String("auth_cookie", "sdcmcook", "Set the cookie name that will be used for the JWT token.")
}

func (hc *httpConfig) parseEnv() {

	os_port := os.Getenv("http_port")
	if os_port == "" {
		configWarn("http_port", *hc.port)
	} else {
		*hc.port = os_port
	}
	hc.Log("Port set to: " + *hc.port + " from environment.")

	os_docker := os.Getenv("http_docker")
	if os_docker == "" {
		configWarn("http_docker", *hc.docker)
	} else {
		if os_docker == "true" {
			*hc.docker = true
		} else if os_docker == "false" {
			*hc.docker = false
		} else {
			configWarn("http_docker", *hc.docker)
		}
	}
	hc.Log(fmt.Sprintf("Docker mode set to: %v from environment.", *hc.docker))

	jwt_env := os.Getenv("jwt_key")
	if jwt_env == "" {
		configWarn("jwt_key", *hc.jwt_key)
	} else {
		*hc.jwt_key = jwt_env
	}
	hc.Logf("JWT key set to: %s from environment.", "*****")

	cook_env := os.Getenv("auth_cookie")
	if cook_env == "" {
		configWarn("auth_cookie", *hc.cookie)
	} else {
		*hc.cookie = cook_env
	}
	hc.Logf("Auth cookie name set to: %s from environment.", *hc.cookie)
}

func (hc *httpConfig) Port() string {
	return *hc.port
}

func (hc *httpConfig) DockerMode() bool {
	return *hc.docker
}

func (hc *httpConfig) JWTKey() string {
	return *hc.jwt_key
}

func (hc *httpConfig) AuthCookie() string {
	return *hc.cookie
}
