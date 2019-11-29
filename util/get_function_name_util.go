package util

import (
	"reflect"
	"runtime"
)

// GetFunctionName returns the name of the given function interface at runtime.
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
