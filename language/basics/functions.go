package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func mult(a int, b int) int {
	return a * b
}

type Point struct {
	x float64
	y float64
}

func squared(a float64) float64 {
	return a * a
}

func distance(a Point, b Point) float64 {
	return math.Sqrt(squared(a.x-b.x) + squared(a.y-b.y))
}

func main() {
	a := Point{1, 1}
	b := Point{1, 3}

	fmt.Println("The distance is", distance(a, b))
}
