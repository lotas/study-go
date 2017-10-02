package main

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

func NewBlockchain() Blockchain {
	bc := Blockchain{}

	b := Block{
		previousHash: "1",
		proof:        1,
	}
	bc.chain = append(bc.chain, b)
	bc.lastBlock = &b

	return bc
}

// NewBlock creates new block
func (bc *Blockchain) NewBlock(proof Proof) *Block {
	previousHash := Hash(bc.lastBlock)

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

// ProofOfWork finds a number p' such that hash(pp') contains leading 4 zeroes
func ProofOfWork(lastProof Proof) Proof {
	var proof Proof
	for validProof(lastProof, proof) == false {
		proof++
	}
	return proof
}

func validProof(lastProof, proof Proof) bool {
	str := fmt.Sprintf("%d%d", lastProof, proof)
	guess := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	return guess[:4] == "0000"
}

func Hash(block *Block) string {
	str, _ := json.Marshal(block)
	return fmt.Sprintf("%x", sha256.Sum256(str))
}
