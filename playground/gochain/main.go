package main

import (
	"encoding/json"
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

	http.HandleFunc("/mine", handleMine)
	http.HandleFunc("/transactions/new", handleTransactionsNew)
	http.HandleFunc("/chain", handleChain)

	http.HandleFunc("/nodes/register", handleNodeRegister)
	http.HandleFunc("/nodes/resolve", handleNodeResolve)

	fmt.Printf("Listening on localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func handleMine(w http.ResponseWriter, r *http.Request) {
	// 1. Calculate next proof
	lastBlock := blockchain.LastBlock
	lastProof := lastBlock.Proof
	newProof := ProofOfWork(lastProof)

	// 2. Receive a reward for mining
	blockchain.NewTransaction("0", "me", 1)

	// 3. Forge new block by adding it to the chain
	block := blockchain.NewBlock(newProof)

	fmt.Fprintf(w,
		"Mined new block #%d\ntransactions: %v\nproof: %d\nprevHash: %s\n",
		block.Index,
		block.Transacations,
		block.Proof,
		block.PreviousHash,
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
	w.Header().Set("Content-Type", "application/json")
	str, _ := json.Marshal(blockchain)
	w.Write(str)
}

func handleNodeRegister(w http.ResponseWriter, r *http.Request) {

}

func handleNodeResolve(w http.ResponseWriter, r *http.Request) {

}
