package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// If you only want the index, you can omit the second variable.
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
