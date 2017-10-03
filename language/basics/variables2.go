package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	d := 5
	y := 6.6

	fmt.Println(i, j, c, python, java)
	fmt.Printf("d: %T = %v\n", d, d)
	fmt.Printf("y: %T = %v\n", y, y)
	fmt.Printf("y*d: %T = %v\n", d*int(y), d*int(y))

}
