package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Feed struct {
	Items []Item
}

type Item struct {
	Preview string
}

type PhotoApi struct {}

func NewPhotoApi() *PhotoApi {
	return &PhotoApi{}
}

func (c PhotoApi) GetRandomItem(apiUrl string) (*Feed, error) {
	url := apiUrl + "0/1/random"
	feed, err := c.RequestItems(url)

	return feed, err
}

func (c PhotoApi) RequestItems(url string) (*Feed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var items []Item
	err = json.Unmarshal(body, &items)

	if err != nil {
		return nil, err
	}

	feed := new(Feed)
	feed.Items = items

	return feed, nil
}
