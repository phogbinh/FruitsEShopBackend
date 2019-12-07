package util

import (
	"net/http"
	"reflect"
	"runtime"

	. "backend/model"
)

// StatusOK is a constant function representing StatusOK.
func StatusOK() Status {
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   ""}
}

// IsStatusOK returns true if the given status's http status code is StatusOK.
func IsStatusOK(status Status) bool {
	return status.HttpStatusCode == http.StatusOK
}

// StatusInternalServerError is a constant function representing StatusInternalServerError.
func StatusInternalServerError(functionInterface interface{}, err error) Status {
	return getErrorStatus(http.StatusInternalServerError, functionInterface, err)
}

// StatusBadRequest is a constant function representing StatusBadRequest.
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
