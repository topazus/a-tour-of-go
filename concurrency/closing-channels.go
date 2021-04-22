package main

import "fmt"

func main() {
    jobs := make(chan int)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
				fmt.Println("sent job", j)
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}
