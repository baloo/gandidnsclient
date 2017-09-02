package gandidnsclient

import (
	"io"
	"log"
	"net/http"
)

type Client struct {
	apiKey string
	l      log.Logger
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey}
}

const defaultBaseUrl = "https://dns.beta.gandi.net/api/v5/"

func (c Client) request(method string, urlPart string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, defaultBaseUrl+urlPart, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", c.apiKey)
	return req, nil
}

func (c Client) log(message string, objects ...interface{}) {
	if nil != c.l {
		c.l.Printf(message, objects...)
	}
}

func (c *Client) SetLogger(l log.Logger) {
	c.l = l
}
