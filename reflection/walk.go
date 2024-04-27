package reflection

import (
	"fmt"
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	// Elem() returns the value that the pointer points to.
	// We can use Value.Kind() to find out the kind of the value.
	value := getValue(x)
	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
		return
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.Chan:
		for {
			x, ok := value.Recv()
			if ok {
				walkValue(x)
			} else {
				break
			}
		}
	case reflect.Func:
		for _, v := range value.Call(nil) {
			walkValue(v)
		}
	default:
		fmt.Println("Not a string or a slice")
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	return value
}
