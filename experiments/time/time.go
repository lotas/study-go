package main

import (
	"fmt"
	"os"
	"time"
)

const shortForm = "2006-Jan-02"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s time", os.Args[0])
		return
	}

	switch arg := os.Args[1]; arg {
	case "now":
		fmt.Println(time.Now())
	default:
		t, _ := time.Parse(shortForm, arg)
		fmt.Printf("%s = %v", arg, t)
	}
}
