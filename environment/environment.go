package environment

import (
	"os"

	"github.com/joho/godotenv"
)

type IEnvProvider interface {
	Load(filenames ...string) (err error)
	GetEnv(key string) string
}
type EnvProvider struct{}

type EnvironmentData struct {
	Token   string
	BaseURL string
}

var env = EnvironmentData{}

func (p EnvProvider) GetEnv(key string) string {
	return os.Getenv(key)
}

func LoadDotEnv(provider IEnvProvider) error {
	return provider.Load(".env")
}

func InitializeEnvVariables(provider IEnvProvider) {
	env.Token = provider.GetEnv("token")
	env.BaseURL = provider.GetEnv("baseUrl")
}

func (p EnvProvider) Load(filenames ...string) (err error) {
	return godotenv.Load(filenames...)
}
func GetEnv() EnvironmentData {
	return env
}
