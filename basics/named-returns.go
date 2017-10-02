package main

import "fmt"

func split(sum int) (x, y, z int) {
	z = x*y - sum
	x = sum * 4 / 9
	y = sum - x - z
	return
}

func main() {
	fmt.Println(split(17))
}
