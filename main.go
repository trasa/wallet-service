package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/wallets", Wallets)
    log.Println("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    log.Printf("Index Request: %s", r)
	fmt.Fprintf(w, "you are at %q", html.EscapeString(r.URL.Path))
}

func Wallets(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get Wallets: %s", r)
	// TODO get wallets
	fmt.Fprint(w, "{}");
}
