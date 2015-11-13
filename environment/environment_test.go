package environment_test

import (
	"github.com/carrot/go-base-api/environment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

func TestEnvironmentSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentSuite))
}

type EnvironmentSuite struct {
	suite.Suite
}

func (suite *EnvironmentSuite) TestTestingEnvironment() {
	testingIsValid := environment.IsValid(environment.TESTING)
	assert.True(suite.T(), testingIsValid, "Testing must be a valid environment for test suites to work")
}

func (suite *EnvironmentSuite) TestSet() {
	// Checking if file exists
	_, err := os.Stat("../.env." + environment.TESTING)
	fileExists := err == nil

	// Testing Set (We are assuming environment.TESTING isValid)
	err = environment.SetWithRelativeDirectory("../", environment.TESTING)
	if err != nil {
		assert.False(suite.T(), fileExists)
	} else {
		assert.True(suite.T(), fileExists)
	}
}
