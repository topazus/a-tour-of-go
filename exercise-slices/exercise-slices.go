package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mypic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		mypic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			mypic[y][x] = uint8(x^y)
		}
	}
	return mypic
}

func main() {
	pic.Show(Pic)
}
