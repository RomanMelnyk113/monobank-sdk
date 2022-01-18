package monobank

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.monobank.ua"
)

type Client struct {
	restClient *http.Client
	baseUrl    string
	token      string
}

func validateErrors(res []byte, err error, status int) error {
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		var msg Error
		if err := json.Unmarshal(res, &msg); err != nil {
			return errors.New("invalid error payload")
		}
		return msg
	}
	return nil
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

// returns client details
func (c *Client) GetUserInfo() (*UserInfo, error) {
	path := "/personal/client-info"
	res, status, err := c.doReq(path, "GET", nil)

	if err := validateErrors(res, err, status); err != nil {
		return nil, err
	}

	var user UserInfo
	if err = json.Unmarshal(res, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// returns client transactions based on {from} and {to} time
func (c *Client) GetTransactions(accountId string, from, to time.Time) ([]Transaction, error) {
	path := fmt.Sprintf("/personal/statement/%s/%d/%d", accountId, from.Unix(), to.Unix())
	res, status, err := c.doReq(path, "GET", nil)
	if err != nil {
		return nil, err
	}
	if err := validateErrors(res, err, status); err != nil {
		return nil, err
	}

	var data []Transaction
	if err = json.Unmarshal(res, &data); err != nil {
		return nil, err
	}

	return data, nil
}
