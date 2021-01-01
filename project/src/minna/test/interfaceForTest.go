package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gist struct {
	Id string `json:"id"`
	RawUrl string `json:"url"`
	File map[string]interface{} `json:"files"`
}

type Getter interface {
	getGists() (io.Reader, error)
}

// Gist の List api を扱うためのクライアントを実装
type Client struct {
	GistGetter Getter
}

type Gister struct {
	user string
}

func (g *Gister) getGists() (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", g.user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}
	return &buf, nil
}

func (c *Client) ListGists() ([]Gist, error) {
	r, err := c.GistGetter.getGists()
	if err != nil {
		return nil, err
	}

	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	data := make([]Gist, 0, len(gists))
	for _, url := range gists {
		data = append(data, url)
	}

	return data, nil
}
