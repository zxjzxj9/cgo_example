package main

// #cgo CXXFLAGS:
// #cgo LDFLAGS:
// #include "test.hh"
import "C"

import "fmt"

func main() {
	a := make([]float64, 10)
	for i := range a {
		a[i] = float64(i)
	}
	fmt.Println("test")
	fmt.Println(a)
	C.print()
	C.addOne((*C.double)(&a[0]), C.int(len(a)))
	fmt.Println(a)
}