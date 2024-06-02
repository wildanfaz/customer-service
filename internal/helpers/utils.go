package helpers

import "reflect"

func IsZeroStruct(x any) bool {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v.IsZero()
}