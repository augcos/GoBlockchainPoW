package blockchainPoW

import (
	"os"
	"bufio"
	"io"
	"log"
	"time"
	"net"
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
)

var bcServer chan []Block

// RunTcp() starts the TCP server
func RunTcp() error {
	bcServer = make(chan []Block)
	server, err := net.Listen("tcp", ":" + os.Getenv("PORT"))
	log.Println("Listening on", os.Getenv("PORT"))
	if err != nil {
		return err
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
	return nil
}

// handleConn() starts the TCP server
func handleConn(conn net.Conn) {
	defer conn.Close()

	// go routine to scan for strings to create new blocks
	io.WriteString(conn, "Enter a new string: ")
	scanner := bufio.NewScanner(conn)
	go func() {
		for scanner.Scan() {
			newBlock, err := GenerateBlock(Blockchain[len(Blockchain)-1],scanner.Text())
			if err != nil {
				log.Println(err)
				continue
			}
			if IsBlockValid(Blockchain[len(Blockchain)-1], newBlock) {
				newBlockchain := append(Blockchain, newBlock)
				ReplaceChain(newBlockchain)
			}
			bcServer<-Blockchain
			io.WriteString(conn, "\nEnter a new string: ")
		}
	}()

	// go routine to broadcast the blockchain
	go func() {
		for {
			time.Sleep(30*time.Second)
			output,err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn,string(output))
		}
	}()
	
	// go routine to print the blockchain on the terminal
	for _ = range bcServer {
		spew.Dump(Blockchain)
	}
}