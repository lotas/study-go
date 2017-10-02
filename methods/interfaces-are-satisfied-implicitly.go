package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
	b int
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t)
}

func main() {
	var i I = T{"hello", 1}
	i.M()
}
