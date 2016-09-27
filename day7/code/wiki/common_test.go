package main

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, got interface{}, expected interface{}, description string) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("%s. Expected %v (type %v) - Got %v (type %v)", description, expected, reflect.TypeOf(expected), got, reflect.TypeOf(got))
	}
}
