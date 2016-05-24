package response

import (
	"github.com/labstack/echo"
	"net/http"
)

type ErrorDetail struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Response struct {
	Context      echo.Context  `json:"-"`
	Success      bool          `json:"success"`
	StatusCode   int           `json:"status_code"`
	StatusText   string        `json:"status_text"`
	ErrorDetails []ErrorDetail `json:"error_details"`
	Content      interface{}   `json:"content"`
}

// New instantiates a new Response struct and attaches the Echo context.
// It returns the Response struct.
func New(c echo.Context) *Response {
	r := new(Response)
	r.Context = c
	r.Success = false
	r.StatusCode = http.StatusInternalServerError
	return r
}

// AddError appends an error to the response via an Error Code.
func (r *Response) AddErrorDetail(code int) {
	r.ErrorDetails = append(r.ErrorDetails, ErrorDetail{code, ErrorDetailText(code)})
}

// AddErrorDetails appends multiple errors to the response via Error Codes.
func (r *Response) AddErrorDetails(codes []int) {
	for _, code := range codes {
		r.ErrorDetails = append(r.ErrorDetails, ErrorDetail{code, ErrorDetailText(code)})
	}
}

// SetResponse sets the response status code and content.
func (r *Response) SetResponse(code int, content interface{}) {
	r.StatusCode = code
	r.Content = content
}

// Render sets the appropriate status, converts the content to JSON and
// passes it to Echo's context for rendering.
func (r *Response) Render() {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		r.Success = true
	}

	r.StatusText = http.StatusText(r.StatusCode)
	r.Context.JSON(r.StatusCode, r)
}
