package utils

import "reflect"

// TypeOf Get struct name
func TypeOf(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}