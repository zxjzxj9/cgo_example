package main

// #cgo CXXFLAGS:
// #cgo LDFLAGS:
// #include "test.hh"
import "C"

import "fmt"

func main() {
	fmt.Println("test")
	C.print()
}