package middleware

import (
	"fmt"
	"github.com/carrot/burrow/response"
	"github.com/labstack/echo"
	"net/http"
	"runtime"
)

// Recover returns a middleware which recovers from panics anywhere in the chain
// and handles the control to the centralized HTTPErrorHandler.
func Recover() echo.MiddlewareFunc {
	// TODO: Provide better stack trace `https://github.com/go-errors/errors` `https://github.com/docker/libcontainer/tree/master/stacktrace`
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					trace := make([]byte, 1<<16)
					n := runtime.Stack(trace, true)

					c.Error(fmt.Errorf("echo => panic recover\n %v\n stack trace %d bytes\n %s",
						err, n, trace[:n]))

					resp := response.New(c)
					resp.SetResponse(http.StatusInternalServerError, nil)
					resp.Render()
				}
			}()
			return h(c)
		}
	}
}
