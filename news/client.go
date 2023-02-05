package news

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClinet(time time.Duration) (*Client, error) {
	if time == 0 {
		return nil, errors.New("time cannot be null")
	}
	return &Client{
		client: &http.Client{
			Timeout:   time,
			Transport: http.DefaultTransport,
		},
	}, nil

}

func (c Client) GetNews() ([]Data, error) {
	resp, err := c.client.Get("https://inshorts.deta.dev/news?category=science")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var r MyResponse

	if err = json.Unmarshal(response, &r); err != nil {
		return nil, err
	}
	return r.News, nil
}
