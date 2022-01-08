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

	// runs the blockchain and the server
	go func() {
		genesisBlock := Block{0, time.Now().String(), "test", "", ""}
		genesisBlock.Hash = CalculateHash(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	log.Fatal(Run())
}
