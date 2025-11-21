package main

import (
	"net/http"
	"github.com/eucaiot/golang_code_exemplos/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Start up!"))
	})
	mux.HandleFunc("/quotes", handler.GetQuote)
	http.ListenAndServe(":8080", mux)
}
