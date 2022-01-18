package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	. "github.com/augcos/GoBlockchainPoW/blockchainPoW"
)



func main() {
	// loads the enviroment variables (.env file) with the port number
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	// creates the genesis block and starts the blockchain
	var genesisBlock Block
	genesisBlock = Block{0, time.Now().String(), "genesis", "", "", Difficulty, ""}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	// runs the server
	log.Fatal(RunTcp())
}