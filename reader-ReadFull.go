package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}

    r = strings.NewReader("")

    buf = make([]byte, 5)
    if _, err := io.ReadFull(r, buf); err != nil {
        // log.Fatal(err)
        fmt.Println("error:", err)
    }
    fmt.Printf("%s\n", buf)
}
