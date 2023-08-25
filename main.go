package main

import (
	"encoding/json"
	"fmt"
	"gymshark/store"
	"log"
	"net/http"
)

func main() {
	// Adding some seed data for packs
	store.SeedData()

	mux := http.NewServeMux()
	mux.HandleFunc("/store/packs", getNumOfPacks)
	mux.HandleFunc("/healthcheck", healthCheck)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getNumOfPacks(w http.ResponseWriter, r *http.Request) {
	var body struct {
		NumOfItems int `json:"num_of_items"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	numOfPacks, err := store.GetNumOfPacks(body.NumOfItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(numOfPacks)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(res))
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, string("Ok"))
}
