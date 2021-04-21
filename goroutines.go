package main

import (
	"fmt"
	"time"
)

func say(num int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(num)
	}
}

func main() {
	go say(1)
	say(2)
}
