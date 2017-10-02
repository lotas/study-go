package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	var str []string
	str = append(str, "one")
	str = append(str, "two")
	printSliceStr(str)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSliceStr(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
