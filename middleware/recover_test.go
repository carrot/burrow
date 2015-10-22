package middleware

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecover(t *testing.T) {
	e := echo.New()
	e.SetDebug(true)

	req, _ := http.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := echo.NewContext(req, echo.NewResponse(rec), e)
	h := func(c *echo.Context) error {
		panic("test")
	}

	Recover()(h)(c)

	body := rec.Body.String()

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, body, `"success":false`)
	assert.Contains(t, body, `"status_code":500`)
}
