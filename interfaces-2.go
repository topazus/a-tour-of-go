package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	a := MyFloat(-math.Sqrt(2))  // a MyFloat implements Abser
	fmt.Println(a.Abs())

	b := &Vertex{3, 4} // a *Vertex implements Abser
	fmt.Println(b.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v
	// fmt.Println(a.Abs())
}
