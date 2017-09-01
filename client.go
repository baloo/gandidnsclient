package gandidnsclient

import (
	"io"
	"net/http"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return apiKey
}

const defaultBaseUrl = "https://dns.beta.gandi.net/api/v5/"

func (c Client) request(method string, urlPart string, body io.Reader) *http.Request {
	req := http.NewRequest(method, defaultBaseUrl+urlPart, body)
	req.Header.Add("X-Api-Key", c.apiKey)
	return req
}
