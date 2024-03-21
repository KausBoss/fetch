package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	receipt.ID = uuid.New()
	receipt.SetPoints()

	// save to main database(here a global variable)
	receipts[receipt.ID] = receipt

	response, err := json.Marshal(map[string]interface{}{"id": receipt.ID})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		return
	}
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	receiptID := params["id"]

	parsedUUID, err := uuid.Parse(receiptID)
	if err != nil {
		http.Error(w, "Error in ID Parsing", http.StatusBadRequest)
		return
	}

	receipt, ok := receipts[parsedUUID]
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(map[string]interface{}{"points": receipt.Points})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		return
	}
}
