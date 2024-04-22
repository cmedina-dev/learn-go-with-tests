package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkInterface := func(val reflect.Value) {
		Walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkInterface(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkInterface(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkInterface(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkInterface(v)
			} else {
				break
			}
		}
	case reflect.Func:
		values := val.Call(nil)
		for _, res := range values {
			walkInterface(res)
		}
	case reflect.String:
		fn(val.String())
	default: // Unknown value found
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
