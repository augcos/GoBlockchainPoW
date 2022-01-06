package main
import (
	"log"
	"time"
	"github.com/joho/godotenv"
	"github.com/davecgh/go-spew/spew"
	. "github.com/augcos/GoBlockchain/blockchain"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), []byte("test"), "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, &genesisBlock)
	}()
	log.Fatal(Run())

}