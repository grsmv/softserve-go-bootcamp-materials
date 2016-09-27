package main

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, got interface{}, expected interface{}, description string) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("%s.\n\tExpected %v (type %v)\n\tGot %v (type %v)", description, expected, reflect.TypeOf(expected), got, reflect.TypeOf(got))
	}
}
