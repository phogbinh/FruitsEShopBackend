package util

import (
	"net/http"
	"reflect"
	"runtime"

	. "backend/model"
)

func StatusOK() Status {
	return Status{
		HttpStatusCode: http.StatusOK,
		ErrorMessage:   ""}
}

func IsStatusOK(status Status) bool {
	return status.HttpStatusCode == http.StatusOK
}

func StatusInternalServerError(functionInterface interface{}, err error) Status {
	return getErrorStatus(http.StatusInternalServerError, functionInterface, err)
}

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
