package main

import (
	"fmt"
	"math"
)
func main() {
	// int and float
	r := 1
	area := math.Pi * float64(r) * float64(r)
	fmt.Printf("the radius: %d, the area: %.2f\n", r, area)

	// string
    word := "hello"
    fmt.Println(word)
    fmt.Printf("%s world", word)

	// Sprintf to format the string without printing it.
	s := fmt.Sprintf("the binary number of 4 is: %b", 4)
	fmt.Println(s)
}
