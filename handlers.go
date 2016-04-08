package main

import (
	"net/http"
	"fmt"
	"html"
	"encoding/json"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you are at %q", html.EscapeString(r.URL.Path))
}

func GetWallets(w http.ResponseWriter, r *http.Request) {
	wallets := Wallets {
		Wallet {  CurrencyType: "gold", Amount: 123},
		Wallet {  CurrencyType: "silver", Amount: 5555},
	}
	json.NewEncoder(w).Encode(wallets)
}

func WalletCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyType := vars["currencyType"]
	fmt.Fprintf(w, "currency type %s", currencyType)
}
