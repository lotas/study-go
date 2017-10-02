package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)

	for _, word := range strings.Fields(s) {
		_, ok := m[word]
		if ok == false {
			m[word] = 1
		} else {
			m[word]++
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
