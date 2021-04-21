package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("fields are: %q\n", strings.Fields(" foo bar  baz "))
}
