package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = 3377

var blockchain Blockchain

func init() {
	blockchain = NewBlockchain()
}

func main() {

	// c.NewTransaction("me", "you", 999)
	// c.NewTransaction("me2", "you2", 111)
	// c.NewTransaction("me3", "you3", 222)
	// b := c.NewBlock(2, "")

	// fmt.Printf("New blockchain %v\n", c)
	// fmt.Printf("Last block index %v\n", b)

	// fmt.Printf("Valid proof for 1: %d\n", ProofOfWork(1))

	http.HandleFunc("/mine", handleMine)
	http.HandleFunc("/transactions/new", handleTransactionsNew)
	http.HandleFunc("/chain", handleChain)

	fmt.Printf("Listening on localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func handleMine(w http.ResponseWriter, r *http.Request) {
	// 1. Calculate next proof
	lastBlock := blockchain.lastBlock
	lastProof := lastBlock.proof
	newProof := ProofOfWork(lastProof)

	// 2. Receive a reward for mining
	blockchain.NewTransaction("0", "me", 1)

	// 3. Forge new block by adding it to the chain
	block := blockchain.NewBlock(newProof)

	fmt.Fprintf(w,
		"Mined new block #%d\ntransactions: %v\nproof: %d\nprevHash: %s\n",
		block.index,
		block.transacations,
		block.proof,
		block.previousHash,
	)
}

func handleTransactionsNew(w http.ResponseWriter, r *http.Request) {
	sender := r.FormValue("sender")
	recipient := r.FormValue("recipient")
	amount, _ := strconv.Atoi(r.FormValue("amount"))

	if sender == "" || recipient == "" || amount == 0 {
		fmt.Fprintf(w, "Invalid request")
		return
	}

	index := blockchain.NewTransaction(Sender(sender), Recipient(recipient), Amount(amount))

	fmt.Fprintf(w, "Transaction will be added to Block %d\n", index)
}

func handleChain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Len: %d\nChain: %v\nUnconfirmed: %v\n",
		len(blockchain.chain),
		blockchain.chain,
		blockchain.currentTransactions,
	)
}
