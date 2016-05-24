package response

import (
	"github.com/carrot/burrow/constants"
	"github.com/carrot/burrow/util"
)

// Enums
//
// These are enums to be used with the `invalidEnumParam()`
// helper function, which generates a string to describe
// that the input did not match one of the valid enums.
//
// `invalidBoolParam()` is a convenience wrapper for booleans
// so you should keep `validBooleans` here in this list,
// `exampleEnum` is safe to remove.
var (
	validBooleans = []string{"true", "false"}
	exampleEnum   = []string{"dogs", "cats", "birds"}
)

// Custom error details
//
// It is suggested you actually type the int value here, as opposed
// to using the iota pattern generally used with constants.
// These integer values actually have meaning, and API consumers
// are to use the values.  Using actual ints will prevent accidental
// value changes in the event something is inserted in the middle of
// the list.  In addition, this makes it easier to lookup what an error
// code means from it's int value.
//
// After adding a value to this list, you must always add it to the
// errorDetailText map.
const (
	ErrorMissingNameParameter   = 1
	ErrorInvalidIdParameter     = 2
	ErrorInvalidEnumParameter   = 3
	ErrorInvalidLimitParameter  = 4
	ErrorInvalidOffsetParameter = 5
)

var errorDetailText = map[int]string{
	ErrorMissingNameParameter:   missingParam("name"),
	ErrorInvalidIdParameter:     invalidIntParam("id"),
	ErrorInvalidEnumParameter:   invalidEnumParam("access_level", exampleEnum),
	ErrorInvalidOffsetParameter: invalidNonNegativeIntParam(constants.OFFSET),
}

// ErrorText returns a code's associated error text
func ErrorDetailText(code int) string {
	return errorDetailText[code]
}

// ===== Error String Helpers =====

func missingParam(name string) string {
	return "Missing parameter `" + name + "`"
}

func missingHeader(name string) string {
	return "Missing header `" + name + "`"
}

func invalidParam(name string, mustBe string) string {
	return "Invalid `" + name + "` parameter, `" + name + "` must be " + mustBe
}

func invalidHeader(name string, mustBe string) string {
	return "Invalid `" + name + "` header, `" + name + "` must be " + mustBe
}

func collisionParam(name string) string {
	return "Duplicate key value violates unique constraint `" + name + "`"
}

func deleteRestricted(dependentModelNameSingular string) string {
	return "Resource cannot be deleted as it is still referenced by a `" +
		dependentModelNameSingular + "`"
}

func invalidDateParam(name string) string {
	return invalidParam(name, "an ISO8601 Date ("+util.ISO8601_LAYOUT+")")
}

func invalidIntParam(name string) string {
	return invalidParam(name, "an integer")
}

func invalidNonNegativeIntParam(name string) string {
	return invalidParam(name, "a non-negative integer")
}

func invalidUuidParam(name string) string {
	return invalidParam(name, "a valid RFC 4122 UUID")
}

func invalidFloatParam(name string) string {
	return invalidParam(name, "a float")
}

func invalidEnumParam(name string, valids []string) string {
	return invalidParam(name, "one of: "+arrayToString(valids))
}

func invalidBoolParam(name string) string {
	return invalidEnumParam(name, validBooleans)
}

func invalidForeignKey(key, modelNameSingular string) string {
	return "The value for `" + key + "` does not reference a valid `" + modelNameSingular + "`"
}

func arrayToString(array []string) string {
	stringArray := "["
	for index, value := range array {
		stringArray += "'" + value + "'"
		if index < (len(array) - 1) {
			stringArray += ", "
		}
	}
	stringArray += "]"
	return stringArray
}
