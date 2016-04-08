package main

type Wallet struct {
	CurrencyType string `json:"currency_type"`
	Amount       int    `json:"amount"`
}
type Wallets []Wallet
