package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Learn the basic components of any Go program."
	r := strings.NewReader(str)

	b := make([]byte, 16)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
