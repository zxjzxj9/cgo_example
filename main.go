package main

// #cgo CXXFLAGS: -I/usr/local/Cellar/libtorch/1.8.1_1/include
// #cgo CXXFLAGS: -I/usr/local/Cellar/libtorch/1.8.1_1/include/torch/csrc/api/include
// #cgo CXXFLAGS: -I/Users/stardust/miniconda3/envs/nnpot/include/python3.7m -std=c++14
// #cgo LDFLAGS: -L/usr/local/Cellar/libtorch/1.8.1_1/lib -ltorch -lc10 -ltorch_cpu -ltorch_global_deps
// #include "test.hh"
import "C"
import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

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
	sval, _ := C.sumVec((*C.double)(&a[0]), C.int(len(a)))
	fmt.Println(sval)

	rawImage := make([]uint8, 1*84*64*3)
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	C.genImage((*C.uchar)(&rawImage[0]), 1*84*64*3)
	for i:=0; i<64; i++ {
		for j:=0; j<64; j++ {
			img.Set(i, j, color.RGBA{rawImage[j*64+i], rawImage[j*64+i+64*64],rawImage[j*64+i+64*64*2], 255})
		}
	}
	f, _ := os.Create("./test.png")
	defer f.Close()
	png.Encode(f, img)
}