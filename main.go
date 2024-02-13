package main

import (
	"math/rand"
	"fmt"
	"log"
	"net/http"
    
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Output)
	http.Handle("/",r)

	log.Fatal(http.ListenAndServe(":9000", r))
}

func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random Number is : %v", randomNumber())
}
func randomNumber() int {
	return rand.Intn(1000)
}
