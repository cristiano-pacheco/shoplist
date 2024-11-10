// Package empty provides utilities for checking if values are empty
package empty

import "reflect"

// IsEmpty checks if the given value is considered empty
// Returns true if the value is zero/empty
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.String:
		return len(value.String()) == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Chan:
		return value.IsNil()
	case reflect.Struct:
		return false // Structs are considered non-empty by default
	case reflect.Ptr:
		if value.IsNil() {
			return true
		}
		return IsEmpty(value.Elem().Interface())
	default:
		return true
	}
}
