package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	. "github.com/augcos/GoBlockchain/blockchain"
)



func main() {
	// loads the enviroment variables (.env file)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}



	// creates the genesis block and starts the blockchain
	go func() {
		genesisBlock := Block{0, time.Now().String(), "test", "", ""}
		genesisBlock.Hash = CalculateHash(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	// runs the server
	log.Fatal(RunTcp())
}
