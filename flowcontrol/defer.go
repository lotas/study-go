package main

import "fmt"

func deferred(a string) {
	defer fmt.Println("deferred()", a)
}

func main() {
	defer fmt.Println("world")

	defer deferred("1")
	deferred("2")
	defer deferred("3")
	fmt.Println("hello")
}
