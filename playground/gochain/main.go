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

	sendJson(struct {
		Msg   string
		Block *Block
	}{
		fmt.Sprintf("Mined new block #%d, proof: %d, prevHash: %s", block.Index, block.Proof, block.PreviousHash),
		block,
	}, w)
}

func handleTransactionsNew(w http.ResponseWriter, r *http.Request) {
	sender := r.FormValue("sender")
	recipient := r.FormValue("recipient")
	amount, _ := strconv.Atoi(r.FormValue("amount"))

	if sender == "" || recipient == "" || amount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	index := blockchain.NewTransaction(Sender(sender), Recipient(recipient), Amount(amount))

	sendJson(struct {
		Index     BlockIndex
		Sender    string
		Recipient string
		Amount    int
	}{
		index,
		sender,
		recipient,
		amount,
	}, w)
}

func handleChain(w http.ResponseWriter, r *http.Request) {
	sendJson(blockchain, w)
}

func handleNodeRegister(w http.ResponseWriter, r *http.Request) {

}

func handleNodeResolve(w http.ResponseWriter, r *http.Request) {

}

func sendJson(data interface{}, w http.ResponseWriter) {
	str, err := json.Marshal(blockchain)

	if err != nil {
		fmt.Printf("Error serializing object to json %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(str)
}
