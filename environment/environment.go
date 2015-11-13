package environment

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

const (
	DEVELOPMENT string = "development"
	TESTING     string = "testing"
	STAGING     string = "staging"
	PRODUCTION  string = "production"

	PORT              string = "PORT"
	PSQL_DATABASE_URL string = "POSTGRES_DATABASE_URL"
)

var (
	activeEnvironment string
)

func Set(env string) error {
	return SetWithRelativeDirectory("./", env)
}

func SetWithRelativeDirectory(relativeDirectory string, env string) error {
	if isValid(env) {
		activeEnvironment = env
		err := godotenv.Load(relativeDirectory + ".env." + activeEnvironment)
		if err != nil {
			return errors.New("Error loading the .env." + activeEnvironment + " file")
		}
		return nil
	} else {
		return errors.New("Invalid environment parameter")
	}
}

func GetActiveEnvironment() string {
	return activeEnvironment
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}

func isValid(env string) bool {
	return env == DEVELOPMENT ||
		env == TESTING ||
		env == STAGING ||
		env == PRODUCTION
}
