package main

import "fmt"

func AddSum(sli []int, c chan int) {
	sum := 0
	for _, v := range sli {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	sli := []int{1, 2, 3, 4}
	c := make(chan int)
	go AddSum(sli, c)
	for i := range c {
		fmt.Println(i)
	}
}
