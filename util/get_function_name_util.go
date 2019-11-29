package util

import (
	"reflect"
	"runtime"
)

// GetErrorMessageHeaderContainingFunctionName returns the header of an error message which contains the given function name.
func GetErrorMessageHeaderContainingFunctionName(functionInterface interface{}) string {
	return "Error occurred at function [" + getFunctionName(functionInterface) + "]" + ColonSpace
}

func getFunctionName(functionInterface interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(functionInterface).Pointer()).Name()
}
