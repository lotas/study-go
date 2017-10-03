package main

import "fmt"

// Go has only one looping construct, the for loop.

func optional() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	sum := 0
	for i := 0; i < 1<<16; i++ {
		sum += i
	}
	fmt.Println(sum)

	optional()
}
