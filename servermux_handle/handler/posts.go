package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eucaiot/servermux_handle/clients"
)

type PostsHandle struct {
	http.HandlerFunc
}

func (p PostsHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.HandlerFunc(w, r)
}

func FindPostById(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Post ID não informado"))
		return
	}

	res, err := clients.GetPostById(postId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Erro ao obter post. ID: %s", postId)))
		return
	}

	json.NewEncoder(w).Encode(res)
}

func SearchPost(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Paramentro 'q' não informado"))
		return
	}

	res, err := clients.SearchPost(q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Erro ao obter post. query: %s", q)))
		return
	}
	json.NewEncoder(w).Encode(res)
}
