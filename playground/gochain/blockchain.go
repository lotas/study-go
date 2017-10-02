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
	Chain               []Block
	CurrentTransactions []Transaction
	LastBlock           *Block
}

func NewBlockchain() Blockchain {
	bc := Blockchain{}

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

func (bc *Blockchain) NewTransaction(s Sender, r Recipient, a Amount) BlockIndex {
	bc.CurrentTransactions = append(bc.CurrentTransactions, Transaction{
		Sender:    s,
		Recipient: r,
		Amount:    a,
	})

	return bc.LastBlock.Index + 1
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
