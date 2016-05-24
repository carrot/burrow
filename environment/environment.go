package environment

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

// Environment Types
const (
	DEVELOPMENT string = "development"
	TESTING     string = "testing"
	STAGING     string = "staging"
	PRODUCTION  string = "production"
)

// Environment Variables
const (
	PORT              string = "PORT"
	PSQL_DATABASE_URL string = "POSTGRES_DATABASE_URL"
)

var (
	activeEnvironment string
)

func Set(env string) error {
	return SetWithRelativeDirectory("./", env)
}

func SetWithRelativeDirectory(relativeDirectory string, activeEnvironment string) error {
	if isValid(activeEnvironment) {
		// Require an env file
		err := godotenv.Load(relativeDirectory + ".env." + activeEnvironment)
		if err != nil && activeEnvironment != PRODUCTION {
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
