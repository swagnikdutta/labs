package reflection

import (
	"fmt"
	"reflect"
)

// This function so far is bad. Yuck.
// Because, note that we are doing the same thing. We want to call walk()
// on each "thing" of x, whether x is a slice or a struct.
//
// But it looks clumsy. At first we have a check for slice, and then we
// assume that it's a slice (See the for loop with val.NumField()).
//
// We will rework on this code and check the type first.

// func walk(x interface{}, fn func(s string)) {
// 	val := getValue(x)
//
// 	// this is the case when x is not a struct but a slice
// 	if val.Kind() == reflect.Slice {
// 		for i := 0; i < val.Len(); i++ {
// 			walk(val.Index(i).Interface(), fn)
// 		}
// 		return
// 	}
//
// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
//
// 		switch field.Kind() {
// 		case reflect.String:
// 			fn(field.String())
// 		case reflect.Struct:
// 			walk(field.Interface(), fn)
// 		case reflect.Slice:
// 			for i := 0; i < field.Len(); i++ {
// 				walk(field.Index(i).Interface(), fn)
// 			}
// 		}
// 	}
// }

func walk(x interface{}, fn func(s string)) {
	val := getValue(x)
	fmt.Println(val.Kind())
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

// This function should be named pointer filter.
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
