package util

import (
	"net/http"
	"reflect"
	"runtime"

	. "backend/model"
)

// StatusOK operates as a constant named StatusOK.
func StatusOK() Status {
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   ""}
}

// IsStatusOK returns true if the given status is StatusOK.
func IsStatusOK(status Status) bool {
	return status.HttpStatusCode == http.StatusOK
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
		ErrorMessage:   getErrorMessageHeaderContainingFunctionName(functionInterface) + err.Error()}
}

func getErrorMessageHeaderContainingFunctionName(functionInterface interface{}) string {
	return "Error occurred at function [" + getFunctionName(functionInterface) + "]" + ColonSpace
}

func getFunctionName(functionInterface interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(functionInterface).Pointer()).Name()
}
