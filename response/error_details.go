package response

// Custom error codes
const (
	ErrorMissingNameParameter = 0
	ErrorInvalidIdParameter = 1
)

var errorDetailText = map[int]string{
	ErrorMissingNameParameter: "Missing parameter `name`",
	ErrorInvalidIdParameter: "Invalid `id` parameter, `id` must be an integer",
}

// ErrorText returns a code's associated error text
func ErrorDetailText(code int) string {
	return errorDetailText[code]
}
