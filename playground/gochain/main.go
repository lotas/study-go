package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var blockchain Blockchain

type ChainResponse struct {
	Length int
	Chain  []Block
}

func init() {
	blockchain = NewBlockchain()
}

func main() {
	port := 3377

	if len(os.Args) > 1 {
		if port2, err := strconv.ParseInt(os.Args[1], 10, 32); err == nil {
			port = int(port2)
		}
	}

	http.HandleFunc("/mine", handleMine)
	http.HandleFunc("/transactions/new", handleTransactionsNew)
	http.HandleFunc("/chain", handleChain)

	http.HandleFunc("/nodes/resolve", handleNodeResolve)
	http.HandleFunc("/nodes/register", handleNodeRegister)

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
		fmt.Fprintf(w, "Required fields: sender, recipient, amount")
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
	sendJson(ChainResponse{
		len(blockchain.Chain),
		blockchain.Chain,
	}, w)
}

func handleNodeRegister(w http.ResponseWriter, r *http.Request) {
	node := r.FormValue("node")
	if node == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No node given")
		return
	}

	newNode := blockchain.RegisterNode(node)

	sendJson(struct {
		Message string
		Node    Node
	}{"New node added", newNode}, w)
}

func handleNodeResolve(w http.ResponseWriter, r *http.Request) {
	// resolving conflicts - find consensus
	fmt.Println("Resolving conflicts")
	result := blockchain.ResolveConflicts()
	fmt.Printf("Conflict resolved: %v\n", result)

	msg := "Our chain is authoritative :)"

	if result {
		msg = "Our chain was replaced :("
	}

	sendJson(struct {
		Message string
	}{
		msg,
	}, w)
}

func sendJson(data interface{}, w http.ResponseWriter) {
	str, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Error serializing object to json %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(str)
}
