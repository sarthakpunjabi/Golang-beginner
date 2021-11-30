package main

import "pic"


func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, 0)
	for i := 0; i < dy; i++ {
		b := make([]uint8, 0)
		for j := 0; j < dx; j++ {
			b = append(b, uint8((i ^ j)))
		}
		a = append(a, b)
	}
	return a
}

func main() {
	pic.Show(Pic)
}
