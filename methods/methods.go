package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v2 := make([]Vertex, 5)
	for i := range v2 {
		v2[i] = Vertex{float64(i), float64(i)}

		fmt.Println(Abs(v2[i]))
	}
}
