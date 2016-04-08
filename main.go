package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"encoding/json"
)

type Wallet struct {
	CurrencyType string	`json:"currency_type"`
	Amount int	`json:"amount"`
}
type Wallets []Wallet

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/wallets", GetWallets)
	router.HandleFunc("/wallets/{currencyType}", WalletCurrency)

    log.Println("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    log.Printf("Index Request")
	fmt.Fprintf(w, "you are at %q", html.EscapeString(r.URL.Path))
}

func GetWallets(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get Wallets")
	// TODO get wallets
	wallets := Wallets {
		Wallet {  CurrencyType: "gold", Amount: 123},
		Wallet {  CurrencyType: "silver", Amount: 5555},
	}
	json.NewEncoder(w).Encode(wallets)
}

func WalletCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyType := vars["currencyType"]
	log.Printf("Get Wallet Currency: %s", currencyType)
	fmt.Fprintf(w, "currency type %s", currencyType)
}
