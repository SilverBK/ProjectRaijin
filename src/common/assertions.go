package common

import "reflect"

func IsNil(result interface{}) bool {
	if result == nil {
		return true
	}

	if reflected := reflect.ValueOf(result); reflected.Kind() != reflect.Struct && reflected.IsNil() {
		return true
	}

	return false
}
