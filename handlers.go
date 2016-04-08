package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you are at %q", html.EscapeString(r.URL.Path))
}

func GetWallets(w http.ResponseWriter, r *http.Request) {
	wallets := Wallets{
		Wallet{CurrencyType: "gold", Amount: 123},
		Wallet{CurrencyType: "silver", Amount: 5555},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(wallets); err != nil {
		panic(err)
	}
}

func WalletCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyType := vars["currencyType"]

	requestedWallet := Wallet{CurrencyType: currencyType, Amount: 500}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(requestedWallet); err != nil {
		panic(err)
	}
}
