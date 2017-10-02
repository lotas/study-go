package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func swapInt(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	c, d := swapInt(50, 60)
	fmt.Println(a, b, c, d)
}
