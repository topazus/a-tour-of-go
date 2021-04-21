package main

import (
	"fmt"
	"math/cmplx"
)

var (
	a		bool		= false
	maxInt	uint64		= 1<<64 -1
	b		complex128	= cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("type: %T value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", b, b)
}
