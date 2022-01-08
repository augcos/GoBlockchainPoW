package blockchain

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
)

// var Blockchain contains the blockchain data
var Blockchain []Block
type Block struct {
	BlockNumber int
	BlockTime 	string
	Data		string
	Hash		string
	PrevHash	string
}

// CalculateHash() returns the hash of a given block
func CalculateHash(block Block) string{
	record := string(block.BlockNumber) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	hash := hex.EncodeToString(hashed)
	return hash
}

// GenerateBlock() returns a pointer to a newly create block
func GenerateBlock(prevBlock Block, data string) (Block, error) {
	var nextBlock Block
	t := time.Now()

	nextBlock.BlockNumber = prevBlock.BlockNumber + 1
	nextBlock.BlockTime = t.String()
	nextBlock.Data = data
	nextBlock.PrevHash = prevBlock.Hash
	nextBlock.Hash = CalculateHash(nextBlock)

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

// IsBlockValid() checks if all blocks in a blockchain are valid
func IsBlockchainValid(testedChain []Block) bool {
	for i:=1; i<(len(testedChain)-1); i++ {
		if !IsBlockValid(testedChain[i-1],testedChain[i]) {
			return false
		}
	}
	return true
}

// ReplaceChain() replaces the blockchain if it is given a longer alternative
func ReplaceChain(newBlockchain []Block) {
	if len(newBlockchain) > len(Blockchain) {
		Blockchain = newBlockchain
	}
}