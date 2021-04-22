package main

import "fmt"

// use a channel for receiving values.
func rec(recs chan<- string, msg string) {
    recs <- msg
}

// use a channel for receiving and sending values.
func recSend(recs <-chan string, recSends chan<- string) {
    msg := <-recs
    recSends <- msg
}

func main() {
    recs := make(chan string, 1)
    recSends := make(chan string, 1)
    rec(recs, "passed message")
    recSend(recs, recSends)
    fmt.Println(<-recSends)
}
