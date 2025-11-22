package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eucaiot/servermux_handle/domain"
)

func GetPostById(id string) (*domain.PostById, error) {
	res, err := http.Get("https://dummyjson.com/posts/" + id)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var post domain.PostById
	err = json.Unmarshal(data, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func SearchPost(q string) (*domain.PostsSearch, error) {
	res, err := http.Get(fmt.Sprintf("https://dummyjson.com/posts/?q=%s", q))
	
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var post domain.PostsSearch
	err = json.Unmarshal(data, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
