package environment

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnvironmentSuites(t *testing.T) {
	suite.Run(t, new(WithoutEnvironmentFileSuite))
	suite.Run(t, new(WithEnvironmentFileSuite))
}

// ==========================================
// ========== WithEnvironmentSuite ==========
// ==========================================

const (
	testFile string = "./.env." + TESTING;
	testVariable = "TEST_ENV_VAR"
	testValue = "value"
)

type WithEnvironmentFileSuite struct {
	suite.Suite
}

func (suite *WithEnvironmentFileSuite) SetupTest() {
	os.Remove(testFile)
	contents := []byte(testVariable + "=" + testValue)
	ioutil.WriteFile(testFile, contents, 0666)
}

func (suite *WithEnvironmentFileSuite) TearDownTest() {
	os.Remove(testFile)
}

func (suite *WithEnvironmentFileSuite) TestIsValid() {
	testingIsValid := isValid(TESTING)
	assert.True(suite.T(), testingIsValid, "Testing must be a valid environment for test suites to work")

	testingIsValid = isValid("SOME_OTHER_ENVIRONMENT")
	assert.False(suite.T(), testingIsValid)
}

func (suite *WithEnvironmentFileSuite) TestSet() {
	err := Set(TESTING)
	assert.Nil(suite.T(), err)
}

func (suite *WithEnvironmentFileSuite) TestSetWithRelativeDirectory() {
	err := SetWithRelativeDirectory("./", TESTING)
	assert.Nil(suite.T(), err)
}

func (suite *WithEnvironmentFileSuite) TestGetActiveEnvironment() {
	Set(TESTING)
	assert.Equal(suite.T(), GetActiveEnvironment(), TESTING)
}

func (suite *WithEnvironmentFileSuite) TestGetEnvVar() {
	Set(TESTING)
	assert.Equal(suite.T(), GetEnvVar(testVariable), testValue)
}

// =============================================
// ========== WithoutEnvironmentSuite ==========
// =============================================

type WithoutEnvironmentFileSuite struct {
	suite.Suite
}

func (suite *WithoutEnvironmentFileSuite) SetupTest() {
	os.Remove(testFile)
}

func (suite *WithoutEnvironmentFileSuite) TestSet() {
	err := Set(TESTING)
	assert.NotNil(suite.T(), err)
}
func (suite *WithoutEnvironmentFileSuite) TestSetWithRelativeDirectory() {
	err := SetWithRelativeDirectory("./", TESTING)
	assert.NotNil(suite.T(), err)
}

func (suite *WithoutEnvironmentFileSuite) TestGetEnvVar() {
	Set(TESTING)
	assert.Equal(suite.T(), GetEnvVar(testVariable), "")
}
