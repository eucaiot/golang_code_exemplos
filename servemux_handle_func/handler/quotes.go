package handler

import (
	"io"
	"log/slog"
	"net/http"
)

const (
	QUOTES_URL = "https://dummyjson.com/quotes/random"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(QUOTES_URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Um erro ocorreu ao obter dados"))
		slog.Error("[DUMMY JSON] - Falha ao chamar API. Resource: %s. Error: %s", QUOTES_URL, err.Error())
		return
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Um erro ocorreu ao obter dados"))
		slog.Error("[DUMMY JSON] - Falha ao ler corpo daresposta API. Resource: %s. Error: %s", QUOTES_URL, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
	slog.Info("[SERVMUX] - Sucess!", "Response Status:", res.Status)
}
