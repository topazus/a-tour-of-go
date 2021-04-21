package main

import "fmt"

func main() {
	// len=5, cap=5
	a := make([]int, 5)
	printSlice("a", a)

	// len=0, cap=5
	b := make([]int, 0, 5)
	printSlice("b", b)

	// len=2, cap=5
	c := b[0:2]
	printSlice("c", c)

	// len=3, cap=3
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
