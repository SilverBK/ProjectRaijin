package common

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, but was %+v", expected, actual)
	}
}

func AssertNil(t *testing.T, result interface{}) {
	if !IsNil(result) {
		t.Fatalf("%+v was recieved, but nil was expected", result)
	}
}

func AssertNotNil(t *testing.T, result interface{}) {
	if IsNil(result) {
		t.Fatalf("provided value was nil")
	}
}
