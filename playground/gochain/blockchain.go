package main

import (
	"fmt"
	"time"
)

type Sender string
type Recipient string
type Amount uint64
type BlockIndex uint64
type Proof uint64

type Node string

type Transaction struct {
	Sender    Sender
	Recipient Recipient
	Amount    Amount
}

type Block struct {
	Index         BlockIndex
	Timestamp     int64
	Transacations []Transaction
	Proof         Proof
	PreviousHash  string
}

type Blockchain struct {
	Nodes               map[string]Node
	Chain               []Block
	CurrentTransactions []Transaction
	LastBlock           *Block
}

func NewBlockchain() Blockchain {
	bc := Blockchain{
		Nodes: make(map[string]Node),
	}

	b := Block{
		PreviousHash: "1",
		Proof:        1,
	}
	bc.Chain = append(bc.Chain, b)
	bc.LastBlock = &b

	return bc
}

// NewBlock creates new block
func (bc *Blockchain) NewBlock(proof Proof) *Block {
	previousHash := Hash(bc.LastBlock)

	b := Block{
		Index:         BlockIndex(len(bc.Chain) + 1),
		Timestamp:     time.Now().UnixNano(),
		Transacations: bc.CurrentTransactions,
		Proof:         proof,
		PreviousHash:  previousHash,
	}

	bc.CurrentTransactions = []Transaction{}

	bc.Chain = append(bc.Chain, b)
	bc.LastBlock = &b

	return &b
}

func (bc *Blockchain) RegisterNode(address string) Node {
	node := Node(address)
	bc.Nodes[address] = node
	return node
}

func (bc *Blockchain) NewTransaction(s Sender, r Recipient, a Amount) BlockIndex {
	bc.CurrentTransactions = append(bc.CurrentTransactions, Transaction{
		Sender:    s,
		Recipient: r,
		Amount:    a,
	})

	return bc.LastBlock.Index + 1
}

// ResolveConflicts is the Consensus Algorithm,
// it resolves conflicts by replacing our chain with the longest one on the network
func (bc *Blockchain) ResolveConflicts() bool {
	var newChain []Block

	// we are only looking for chains longer than ours
	maxLen := len(bc.Chain)

	for _, node := range bc.Nodes {
		fmt.Printf("[%s]\n", node)

		if otherChain, err := fetchChain(node); err == nil {
			fmt.Printf("[%s] Recieved: %+v\n", node, otherChain)
			if otherChain.Length > maxLen && ValidChain(otherChain.Chain) {
				maxLen = otherChain.Length
				newChain = otherChain.Chain
				fmt.Printf("[%s] has longer chain: %d\n", node, maxLen)
			} else {
				fmt.Printf("[%s] is invalid or smaller: %d\n", node, otherChain.Length)
			}
		}
	}

	if newChain != nil {
		bc.Chain = newChain
		return true
	}

	return false
}
