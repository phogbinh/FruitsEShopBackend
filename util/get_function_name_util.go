package util

import (
	"reflect"
	"runtime"
)

func getFunctionName(functionInterface interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(functionInterface).Pointer()).Name()
}
