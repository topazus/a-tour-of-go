package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
	test := WordCount("an tour of go, a journey of fun")
	fmt.Println(test)
}
