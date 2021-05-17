package main

// #cgo CXXFLAGS: -I/usr/local/Cellar/libtorch/1.8.1/include
// #cgo CXXFLAGS: -I/usr/local/Cellar/libtorch/1.8.1/include/torch/csrc/api/include
// #cgo CXXFLAGS: -I/Users/stardust/miniconda3/envs/nnpot/include/python3.7m -std=c++14
// #cgo LDFLAGS: -L/usr/local/Cellar/libtorch/1.8.1/lib -ltorch -lc10 -ltorch_cpu -ltorch_global_deps
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
	C.fillRandn((*C.double)(&a[0]), C.int(len(a)))
	fmt.Println(a)
}