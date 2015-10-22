package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorTextFound(t *testing.T) {
	expected := "Record Not Found"
	actual := ErrorText(ErrorRecordNotFound)

	assert.Equal(t, expected, actual)
}

func TestErrorTextNotFound(t *testing.T) {
	expected := ""
	actual := ErrorText(999999)

	assert.Equal(t, expected, actual)
}
