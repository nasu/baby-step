package main

import (
	bar "bar/baz"
	baz "bar/baz/qux"
	"foo/v1"
)

func main() {
	foo.HelloV1()
	bar.HelloBar()
	baz.HelloBaz()
}
