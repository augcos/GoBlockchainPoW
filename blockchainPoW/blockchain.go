package blockchainPoW

import (
	"fmt"
	"strings"
	"time"
	"crypto/sha256"
	"encoding/hex"
)

// mining difficulty (number of zeros that the hash must have)
const Difficulty = 3

// var Blockchain contains the multiple Blocks
var Blockchain []Block
type Block struct {
	BlockNumber int
	BlockTime 	string
	Data		string
	Hash		string
	PrevHash	string
	Difficulty  int
	Nonce 		string
}

// CalculateHash() returns the hash of a given block. It gets the block number, the
// previous hash and the nonce and calculates the sha256 hash in string form
func CalculateHash(block Block) string{
	record := string(block.BlockNumber) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	hash := hex.EncodeToString(hashed)
	return hash
}

// GenerateBlock() creates a new block given the previous block and a string
func GenerateBlock(prevBlock Block, data string) (Block, error) {
	var nextBlock Block
	t := time.Now()

	nextBlock.BlockNumber = prevBlock.BlockNumber + 1
	nextBlock.BlockTime = t.String()
	nextBlock.Data = data
	nextBlock.PrevHash = prevBlock.Hash
	nextBlock.Difficulty = Difficulty

	// The nonce is tried over and over until it matches the network difficulty
	for i:=0;; i++ {
		hex := fmt.Sprintf("%x",i)
		nextBlock.Nonce = hex
		if !IsHashValid(CalculateHash(nextBlock), nextBlock.Difficulty) {
			fmt.Println(CalculateHash(nextBlock), "do more work!")
			continue
		} else {
			fmt.Println(CalculateHash(nextBlock), "work done!")
			nextBlock.Hash = CalculateHash(nextBlock)
			break
		}
	}
	return nextBlock, nil
}

// IsBlockValid() checks if a block's BlockNumber, PrevHash and Hash are correct
func IsBlockValid(prevBlock, nextBlock Block) bool {
	if prevBlock.BlockNumber+1 != nextBlock.BlockNumber {
		return false
	}
	if prevBlock.Hash != nextBlock.PrevHash {
		return false
	}
	if CalculateHash(nextBlock) != nextBlock.Hash {
		return false
	}
	return true
}

// IsBlockchainValid() checks if all blocks in a blockchain are valid
func IsBlockchainValid(testedChain []Block) bool {
	for i:=1; i<(len(testedChain)-1); i++ {
		if !IsBlockValid(testedChain[i-1],testedChain[i]) {
			return false
		}
	}
	return true
}

// ReplaceChain() replaces the blockchain if it is given a longer alternative
func ReplaceChain(nextBlockchain []Block) {
	if len(nextBlockchain) > len(Blockchain) {
		Blockchain = nextBlockchain
	}
}

// IsHashValid() checks if hash is right difficulty
func IsHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}