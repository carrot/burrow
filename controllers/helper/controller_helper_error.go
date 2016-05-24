package helper

import "github.com/carrot/burrow/response"

type HelperError struct {
	HttpStatusCode   int
	ErrorDetailCodes []int
	Error            error
}

func NewHelperError(httpStatusCode int, err error) *HelperError {
	return &HelperError{
		HttpStatusCode: httpStatusCode,
		Error:          err,
	}
}

func (he *HelperError) AddErrorDetailCode(code int) {
	he.ErrorDetailCodes = append(he.ErrorDetailCodes, code)
}

// PrepareResponse mutates a response object from a HelperError
// This function also returns the helper error so this function can be used:
// `return helper.PrepareResponse(resp, helperError)` in a controller
func PrepareResponse(resp *response.Response, helperError *HelperError) error {
	resp.SetResponse(helperError.HttpStatusCode, nil)
	resp.AddErrorDetails(helperError.ErrorDetailCodes)
	return helperError.Error
}
