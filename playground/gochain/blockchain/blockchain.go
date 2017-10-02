package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Sender string
type Recipient string
type Amount uint64
type BlockIndex uint64
type Proof uint64

type Transaction struct {
	sender    Sender
	recipient Recipient
	amount    Amount
}

type Block struct {
	index         BlockIndex
	timestamp     int64
	transacations []Transaction
	proof         Proof
	previousHash  string
}

type Blockchain struct {
	chain               []Block
	currentTransactions []Transaction
	lastBlock           *Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}

	b := Block{
		previousHash: "1",
		proof:        100,
	}
	bc.chain = append(bc.chain, b)
	bc.lastBlock = &b

	return bc
}

// NewBlock creates new block
func (bc *Blockchain) NewBlock(proof Proof, previousHash string) *Block {
	prevHash := previousHash
	if prevHash == "" {
		prevHash = Hash(bc.lastBlock)
	}
	b := Block{
		index:         BlockIndex(len(bc.chain) + 1),
		timestamp:     time.Now().UnixNano(),
		transacations: bc.currentTransactions,
		proof:         proof,
		previousHash:  previousHash,
	}

	bc.currentTransactions = []Transaction{}

	bc.chain = append(bc.chain, b)
	bc.lastBlock = &b

	return &b
}

func (bc *Blockchain) NewTransaction(s Sender, r Recipient, a Amount) BlockIndex {
	bc.currentTransactions = append(bc.currentTransactions, Transaction{
		sender:    s,
		recipient: r,
		amount:    a,
	})

	return bc.lastBlock.index + 1
}

func Hash(block *Block) string {
	str, _ := json.Marshal(block)
	return fmt.Sprintf("%x", sha256.Sum256(str))
}
