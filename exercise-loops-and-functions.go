package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, int) {
	z := 10.0
	i := 0
	for ;i<10; i++ {
		z -= (z*z -x) / (2*z)
		if (z*z - x) / x <= 0.00001 {
			break
		}
	}
	return z, i
}

func main() {
	num := 2.0
	v, times := Sqrt(num)
	fmt.Println("after %v loops to guarantee the accuaracy of (z*z - x) / x <= 0.00001, the square root of %v is %v", times, num, v)
}
