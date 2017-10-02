package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	sqrt := x
	for n := 1; n < 100; n++ {
		sqrt = sqrt - (sqrt*sqrt-x)/(2*sqrt)
	}
	return sqrt
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt2(x float64) float64 {
	sqrt, prev, delta := x, 0.0, 0.00000001

	for abs(prev-sqrt) > delta {
		prev = sqrt
		sqrt = sqrt - (sqrt*sqrt-x)/(2*sqrt)
	}
	return sqrt
}

func main() {
	fmt.Println(Sqrt(443))
	fmt.Println(Sqrt2(443))
}
