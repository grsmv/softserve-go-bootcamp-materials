package inmem

import "fmt"

var (
	errNotFound       = fmt.Errorf("Key not found")
	errNotImplemented = fmt.Errorf("Not implemented")
	errNotExists      = fmt.Errorf("Not exists")
)
