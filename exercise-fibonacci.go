package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	num1 := 0
	num2 := 1
	return func() int {
		// sum := num1 + num2
		// num1 = num2
		// num2 = sum
		num1, num2 = num2, num1 + num2
		return num1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
