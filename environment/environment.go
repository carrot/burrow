package environment
import "errors"

const (
	DEVELOPMENT string = "development"
	TESTING     string = "testing"
	STAGING     string = "staging"
	PRODUCTION  string = "production"
)

var (
	activeEnvironment string
)


func Set(env string) error {
	if isValid(env) {
		activeEnvironment = env
		return nil
	} else {
		return errors.New("Invalid environment parameter")
	}
}

func Get() string {
	return activeEnvironment
}

func isValid(env string) bool {
	return env == DEVELOPMENT ||
		env == TESTING ||
		env == STAGING ||
		env == PRODUCTION
}
