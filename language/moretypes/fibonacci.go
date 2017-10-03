package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := 0
	fib2 := 1
	return func() int {
		res := fib + fib2
		fib = fib2
		fib2 = res
		return res - fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
