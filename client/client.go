package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiSearchEndpoint = "https://jisho.org/api/v1/search/words?keyword="
)

func Search(query string) ([]Word, error) {
	resp, err := http.Get(apiSearchEndpoint + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var d struct {
		Data []Word `json:"data"`
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	return d.Data, nil
}
