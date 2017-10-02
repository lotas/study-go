package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		row := make([]uint8, dx)
		for j := range row {
			row[j] = uint8(i * j)
		}
		pic[i] = row
	}
	return pic
}

func main() {
	fmt.Print("Hi")
	pic.Show(Pic)
}
