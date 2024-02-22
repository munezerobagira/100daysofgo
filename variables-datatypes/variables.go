package main

import (
	"fmt"
)

func main() {
	// 2^32=4294967296
	// var num float64 = 42.2
	// fmt.Printf("%f is the minimum negative number for float32\n, Its size is %d bytes", num, unsafe.Sizeof(num))
	var b complex128 = complex(12, 5)
	var a complex64 = complex(2, 3)
	var c byte = 255   // alias fo uint8
	var d rune = 65535 // alias for int32
	var e bool = true
	fmt.Println(a, b, c, d, e)
}
