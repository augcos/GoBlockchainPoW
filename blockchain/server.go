package blockchain

import (
	"os"
	"io"
	"log"
	"time"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)


// Run() starts the blockchain server
func Run() error {
	mux := MuxHandler()
	httpAddr := os.Getenv("PORT")
	log.Println("Listening on ", httpAddr)
	server := &http.Server {
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

// MakeMuxRouter() returns the http mux with the appropiate functions
func MuxHandler() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", GetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", PostBlockchain).Methods("POST")
	return muxRouter
}

// GetBlockchain() is the answer for the GET request
func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// GetBlockchain() is the answer for the GET request
type PostMessage struct {
	Data string
}

// PostBlockchain() is the answer for the POST request
func PostBlockchain(w http.ResponseWriter, r *http.Request) {
	var msg PostMessage
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		jsonRespond(w, r, http.StatusBadRequest, r.Body)
		return
	}
	newBlock, err := GenerateBlock(Blockchain[len(Blockchain)-1], msg.Data)
	if err != nil {
		jsonRespond(w, r, http.StatusInternalServerError, msg)
		return
	}
	if IsBlockValid(Blockchain[len(Blockchain)-1],newBlock) {
		newBlockchain := append(Blockchain, newBlock)
		ReplaceChain(newBlockchain)
	}
	jsonRespond(w, r, http.StatusCreated, newBlock)
}

// respondWithJSON() is the answer for the POST request
func jsonRespond(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}