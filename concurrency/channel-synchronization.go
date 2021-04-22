package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
	// send a bool value to notify that we are done
	done <- true

}

func main() {
    done := make(chan bool, 1)
    go worker(done)
    <-done
}
