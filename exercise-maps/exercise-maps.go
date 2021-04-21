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

// original code of wc.Test
func wcTest(f func(string) map[string]int) {
	ok := true
	for _, c := range testCases {
		// use WordCount function to calculate the frequency of
		// words
		got := f(c.in)
		if len(c.want) != len(got) {
			ok = false
		} else {
			for k := range c.want {
				if c.want[k] != got[k] {
					ok = false
				}
			}
		}
		if !ok {
			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v\n",
				c.in, got, c.want)
			break
		}
		fmt.Printf("PASS\n f(%q) =\n  %#v\n", c.in, got)
	}
}

func main() {
	wc.Test(WordCount)
	fmt.Println("---")
	test := WordCount("an tour of go, a journey of fun")
	fmt.Println(test)
	fmt.Println("---")
	wcTest(WordCount)
}

var testCases = []struct {
	in		string
	want	map[string]int
}{
	{"I am learning Go!", map[string]int{
		"I": 1, "am": 1, "learning": 1, "Go!": 0,
	}},
	{"The quick brown fox jumped over the lazy dog.", map[string]int{
		"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
		"over": 1, "the": 1, "lazy": 1, "dog.": 1,
	}},
}
