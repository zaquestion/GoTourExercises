package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dx)
	for i, v := range a {
		v = make([]uint8, dy)
		for iy := range v {
			a[i][iy] = uint8(i ^ iy)
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
