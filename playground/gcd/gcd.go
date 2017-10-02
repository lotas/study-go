package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	fmt.Printf("gcd(30,15)=%d", gcd(30, 15))
}
