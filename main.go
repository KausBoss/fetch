package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receipts map[uuid.UUID]Receipt

func init() {
	receipts = make(map[uuid.UUID]Receipt)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
