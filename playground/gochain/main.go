package main

import (
	"fmt"

	"./blockchain"
)

func main() {
	c := blockchain.NewBlockchain()

	c.NewTransaction("me", "you", 999)
	c.NewTransaction("me2", "you2", 111)
	c.NewTransaction("me3", "you3", 222)
	b := c.NewBlock(2, "")

	fmt.Printf("New blockchain %v\n", c)
	fmt.Printf("Last block index %v\n", b)
}
