package main

import "fmt"

func create_channel() {
	// Create a channel using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel 1:", mychannel)
	fmt.Printf("Type of the channel 1: %T\n", mychannel)

	// Create a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("Value of the channel 2:", mychannel1)
	fmt.Printf("Type of the channel 2: %T\n", mychannel1)
}

func add(ch chan int) {
	fmt.Println(234 + <-ch)
}

func send_and_receive() {
	fmt.Println("start a channel")
	ch := make(chan int)
	go add(ch)
	ch <- 23

	fmt.Println("end a channel")
}

func close_channel(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "hi"
	}
	close(ch)
}

func test_close_channel() {
	ch := make(chan string)
	go close_channel(ch)
	for {
		if res, ok := <-ch; ok == false {
			fmt.Println("channel close ", ok)
			break
		} else {
			fmt.Println("channel open ", res, ok)
		}
	}
}

func anonymous_goroutine() {
	ch := make(chan string)
	go func() {
		ch <- "hi"
		ch <- "hello"
		close(ch)
	}()
	for res := range ch {
		fmt.Println(res)
	}
}

func len_and_cap() {
	ch := make(chan string, 5)
	ch <- "apple"
	ch <- "linux"

	fmt.Println("the length of the channel is", len(ch))
	fmt.Println("the capacity of the channel is", cap(ch))
}

func main() {
	//create_channel()
	//send_and_receive()
	//test_close_channel()
	//anonymous_goroutine()
	len_and_cap()
}
