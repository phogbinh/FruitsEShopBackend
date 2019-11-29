package util

import (
	"net/http"

	. "backend/model"
)

// StatusOK operates as a constant named StatusOK.
func StatusOK() Status {
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   ""}
}

// StatusInternalServerError returns an http.StatusInternalServerError associated with the error message consisting of the function name and the given error.
func StatusInternalServerError(functionInterface interface{}, err error) Status {
	return getErrorStatus(http.StatusInternalServerError, functionInterface, err)
}

// StatusBadRequest returns an http.StatusBadRequest associated with the error message consisting of the function name and the given error.
func StatusBadRequest(functionInterface interface{}, err error) Status {
	return getErrorStatus(http.StatusBadRequest, functionInterface, err)
}

func getErrorStatus(httpStatusCode int, functionInterface interface{}, err error) Status {
	return Status{
		HttpStatusCode: httpStatusCode,
		ErrorMessage:   GetErrorMessageHeaderContainingFunctionName(functionInterface) + err.Error()}
}

// IsStatusOK returns true if the given status is StatusOK.
func IsStatusOK(status Status) bool {
	return status.HttpStatusCode == http.StatusOK
}
