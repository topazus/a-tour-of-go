package main

import "fmt"

func main() {
	// queue := make(chan string), error!
	queue := make(chan string, 3)
	str := []string{"one", "two"}
	for _, v := range str {
		queue <- v
	}
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
