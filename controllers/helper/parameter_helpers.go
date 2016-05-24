package helper

import (
	"github.com/carrot/burrow/constants"
	"github.com/carrot/burrow/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetLimit(c echo.Context) (int64, *HelperError) {
	limit := int64(10)
	limitParam := c.QueryParam(constants.LIMIT)
	if limitParam != "" {
		limitParamInt, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			helperError := NewHelperError(http.StatusBadRequest, err)
			helperError.AddErrorDetailCode(response.ErrorInvalidLimitParameter)
			return 0, helperError
		} else {
			if limitParamInt < 0 {
				helperError := NewHelperError(http.StatusBadRequest, nil)
				helperError.AddErrorDetailCode(response.ErrorInvalidLimitParameter)
				return 0, helperError
			}
			limit = limitParamInt
		}
	}
	return limit, nil
}

func GetOffset(c echo.Context) (int64, *HelperError) {
	offset := int64(0)
	offsetParam := c.QueryParam(constants.OFFSET)
	if offsetParam != "" {
		offsetParamInt, err := strconv.ParseInt(offsetParam, 10, 64)
		if err != nil {
			helperError := NewHelperError(http.StatusBadRequest, err)
			helperError.AddErrorDetailCode(response.ErrorInvalidOffsetParameter)
			return 0, helperError
		} else {
			if offsetParamInt < 0 {
				helperError := NewHelperError(http.StatusBadRequest, err)
				helperError.AddErrorDetailCode(response.ErrorInvalidOffsetParameter)
				return 0, helperError
			}
			offset = offsetParamInt
		}
	}
	return offset, nil
}
