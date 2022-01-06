package blockchain

import (
	"time"
	"crypto/sha256"
	"encoding/hex"
)

var Blockchain []*Block
type Block struct {
	BlockNumber int
	BlockTime 	string
	Data		[]byte
	Hash		string
	PrevHash	string

}

func CalculateHash(block *Block) string{
	record := string(block.BlockNumber) + block.BlockTime + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	hash := hex.EncodeToString(hashed)
	return hash
}

func GenerateBlock(prevBlock *Block, data []byte) (*Block, error) {
	var nextBlock *Block
	t := time.Now()

	nextBlock.BlockNumber = prevBlock.BlockNumber + 1
	nextBlock.BlockTime = t.String()
	nextBlock.Data = data
	nextBlock.Hash = CalculateHash(nextBlock)
	nextBlock.PrevHash = prevBlock.Hash

	return nextBlock, nil
}

func IsBlockValid(prevBlock, nextBlock *Block) bool {
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

func ReplaceChain(newBlockchain []*Block) {
	if len(newBlockchain) > len(Blockchain) {
		Blockchain = newBlockchain
	}
}