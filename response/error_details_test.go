package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorTextFound(t *testing.T) {
	expected := "Missing parameter `name`"
	actual := ErrorDetailText(ErrorMissingNameParameter)
	assert.Equal(t, expected, actual)
}

func TestErrorTextNotFound(t *testing.T) {
	expected := ""
	actual := ErrorDetailText(999999)
	assert.Equal(t, expected, actual)
}
