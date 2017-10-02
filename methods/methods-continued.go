package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type MyInt int

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyInt) AbsInt() int {
	if f < 0 {
		return int(-f)
	}
	return int(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	a := MyInt(-5)
	fmt.Print(a.AbsInt())
}
