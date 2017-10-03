package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

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

// ValidChain checks if the given chain is valid
func ValidChain(chain []Block) bool {
	lastBlock := chain[0]

	for i := 1; i < len(chain); i++ {
		block := chain[i]

		if block.PreviousHash != Hash(&lastBlock) {
			fmt.Printf("Hash of the %v doesn't match prevHash: %s", lastBlock, block.PreviousHash)
			return false
		}

		lastBlock = block
	}

	return true
}

// Hash calculates the hash of the block
func Hash(block *Block) string {
	str, _ := json.Marshal(block)
	return fmt.Sprintf("%x", sha256.Sum256(str))
}

func fetchChain(node Node) (*ChainResponse, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/chain", node))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Cannot fetch status: %s", resp.Status)
	}

	var chainRes ChainResponse

	if err := json.NewDecoder(resp.Body).Decode(&chainRes); err != nil {
		return nil, err
	}

	return &chainRes, nil
}
