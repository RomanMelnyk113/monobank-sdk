package monobank

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.monobank.ua"
)

type Client struct {
	restClient *http.Client
	baseUrl    string
	token      string
}

func NewClient(token string) Client {
	restClient := &http.Client{}
	return Client{restClient, baseURL, token}
}

func (c *Client) doReq(path string, method string, body io.Reader) ([]byte, int, error) {
	// TODO: add method validation
	// TODO: add context support to be able to cancel long requests
	// TODO: add support for custom headers
	reqURL, err := url.Parse(c.baseUrl + path)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build URL: %w", err)
	}
	req, err := http.NewRequest(method, reqURL.String(), body)
	req.Header.Add("X-Token", c.token)
	res, err := c.restClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	return data, res.StatusCode, err
}
