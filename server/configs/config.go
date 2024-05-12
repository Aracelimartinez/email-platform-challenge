package configs

import (
	"os"
)

type Credentials struct {
	Username string
	Password string
	Host     string
}

type Config struct {
	//Data to connect with zincsearch
	ZincSearchCredentials Credentials
	APIPort               string
}

var (
	GlobalConfig Config
)

func init() {
	LoadEnv()
}

// Initialize the env variables
func LoadEnv() Config {

	GlobalConfig = Config{
		ZincSearchCredentials: Credentials{
			Username: os.Getenv("ZINCSEARCH_USERNAME"),
			Password: os.Getenv("ZINCSEARCH_PASSWORD"),
			Host:     os.Getenv("ZINCSEARCH_HOST"),
		},
		APIPort: os.Getenv("API_PORT"),
	}
	return GlobalConfig
}
