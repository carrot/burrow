package response

// Custom error codes
const (
	ErrorNoError                   = 0
	ErrorInvalidApiKey             = 1
	ErrorEmailPasswordInvalid      = 2
	ErrorAccountTemporarilyLocked  = 3
	ErrorEmailAddressNotAvailable  = 4
	ErrorMissingRequiredParameters = 5
	ErrorInvalidParameters         = 6
	ErrorRecordNotFound            = 7
	ErrorInvalidApplicationSecret  = 8
	ErrorSessionInvalid            = 9
	ErrorNoContent                 = 10
	ErrorRecordNotCreated          = 11
	ErrorRecordNotUpdated          = 12
	ErrorRecordNotDestroyed        = 13
)

var errorText = map[int]string{
	ErrorNoError:                   "No Error",
	ErrorInvalidApiKey:             "Invalid API Key",
	ErrorEmailPasswordInvalid:      "Email and/or Password Invalid",
	ErrorAccountTemporarilyLocked:  "Account is Temporarily Locked",
	ErrorEmailAddressNotAvailable:  "Email Address Already in Use",
	ErrorMissingRequiredParameters: "Missing Required Parameters",
	ErrorInvalidParameters:         "Invalid Parameters",
	ErrorRecordNotFound:            "Record Not Found",
	ErrorInvalidApplicationSecret:  "Invalid Application Secret",
	ErrorSessionInvalid:            "Session is Invalid/Expired",
	ErrorNoContent:                 "No Content to Return",
	ErrorRecordNotCreated:          "Record Not Created",
	ErrorRecordNotUpdated:          "Record Not Updated",
	ErrorRecordNotDestroyed:        "Record Not Removed",
}

// Returns the associated error text
func ErrorText(code int) string {
	return errorText[code]
}
