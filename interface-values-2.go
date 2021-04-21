package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	i := &T{"Hello"}
	describe(i)
	i.M()

	j := F(math.Pi)
	describe(j)
	j.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
