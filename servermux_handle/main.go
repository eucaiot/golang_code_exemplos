package main

import (
	"net/http"

	"github.com/eucaiot/servermux_handle/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("/posts/{id}", handler.PostsHandle{
		HandlerFunc: handler.FindPostById,
	})

	mux.Handle("/posts/search", handler.PostsHandle{
		HandlerFunc: handler.SearchPost,
	})
	http.ListenAndServe(":8080", mux)
}
