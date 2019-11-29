package util

import (
	"reflect"
	"runtime"
)

func getErrorMessageHeaderContainingFunctionName(functionInterface interface{}) string {
	return "Error occurred at function [" + getFunctionName(functionInterface) + "]" + ColonSpace
}

func getFunctionName(functionInterface interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(functionInterface).Pointer()).Name()
}
