package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                fmt.Println("Ticker stopped")
				return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1500 * time.Millisecond)
    ticker.Stop()
    done <- true // comment this line, the result will be 2 tickers.
}
