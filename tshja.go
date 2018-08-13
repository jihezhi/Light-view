package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	bo := [][]uint8{}
	for i := 0; i < dy; i++ {
		sr := make([]uint8, 0, dx)
		for j := 0; j < dx; j++ {
			sr = append(sr, uint8(i+j))
		}
		bo = append(bo, sr)
	}
	return bo
}
func printSlice(s []uint) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	pic.Show(Pic)
}
