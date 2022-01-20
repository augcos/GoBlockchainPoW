# GoBlockchainPoW
## Introduction
This project is my personal Go implementation of a proof-of-work blockchain proposed by [Coral Health](https://github.com/nosequeldeebee/blockchain-tutorial). If you want to see the proof-of-stake version of this blockchain, click [here](https://github.com/augcos/GoBlockchainPoS). This project was developed using Go v1.17.5 for Linux systems.

## How to install?
First, you will have to make sure to have preinstalled the required third-party packages. You can install them using the following commands:
```
go get github.com/joho/godotenv
go get github.com/gorilla/mux
go get github.com/davecgh/go-spew/spew
```
You can clone this repository to your local system using the command:
```
git clone github.com/augcos/GoBlockchainPoW
```

## How to run?
### Using a TCP connection
By default, GoBlockchainPoW will run using a TCP server. In order to launch both the server and the blockchain, run the main.go file:
```
go run main.go
```
Then, open a different terminal and connect to the TCP server:
```
nc localhost 8080
```
You will be prompted to input a string as data for a new block. Once a new block has been proposed, you will see the mining process in the first terminal until a nonce appropiate to the difficulty is found and the block is attached to the blockchain. 

### Using a HTTP conection
In order to run the blockchain using a HTTP server, you will need to change the log.Fatal(runTcp()) line at the end of the main.go file to log.Fatal(runHttp()). Then, launch the server and the blockchain using the command:
```
go run main.go
```
In order to propose a new block, you will need to do a POST request to localhost:8080 with the following structure:
```
{
    "Data": "[string]"
}
```
After the mining process is completed and the new block is attached, you can see the full blockchain doing a GET request to localhost:8080.