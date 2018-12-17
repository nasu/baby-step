package main

import (
	"bytes"
)

func main() {
	a := []byte{'a', 'b'}
	b := []byte{'a', 'b'}
	_op(a, b)
	_string(a, b)
	_bytesEqual(a, b)
}

func _op(a, b []byte) bool {
	// This occurs an error.
	//return a == b
	return false
}

func _string(a, b []byte) bool {
	return string(a) == string(b)
}

func _bytesEqual(a, b []byte) bool {
	return bytes.Equal(a, b)
}
