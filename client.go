package sendpulse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ResponseError struct {
	HttpCode int
	Url      string
	Body     string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("Http code: %d, url: %s, body: %s", e.HttpCode, e.Url, e.Body)
}

type client struct {
	userId  string
	secret  string
	token   string
	timeout int
}

const apiBaseUrl = "https://api.sendpulse.com"

func (c *client) makeRequest(path string, method string, data map[string]string, useToken bool) ([]byte, error) {
	q := url.Values{}
	for param, value := range data {
		q.Add(param, value)
	}

	method = strings.ToUpper(method)

	fullPath := fmt.Sprintf(apiBaseUrl+"%s", path)
	req, e := http.NewRequest(method, fullPath, bytes.NewBufferString(q.Encode()))
	if e != nil {
		return nil, errors.New(fmt.Sprintf("Url parse error: %s", fullPath))
	}

	if method == "GET" {
		req.URL.RawQuery = q.Encode()
		req.Body = nil
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}

	client := &http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}

	if useToken {
		req.Header.Add("Authorization", "Bearer "+c.token)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized && useToken {
		err := c.refreshToken()
		if err != nil {
			return nil, err
		}
		respData, err := c.makeRequest(path, method, data, useToken)
		if err != nil {
			return nil, err
		}
		return respData, nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &ResponseError{resp.StatusCode, fullPath, string(body)}
	}

	return body, nil
}

func (c *client) refreshToken() error {
	data := make(map[string]string)
	data["grant_type"] = "client_credentials"
	data["client_id"] = c.userId
	data["client_secret"] = c.secret
	path := "/oauth/access_token"

	body, err := c.makeRequest(path, "POST", data, false)

	if err != nil {
		return err
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		return &ResponseError{http.StatusOK, fmt.Sprintf(apiBaseUrl+"%s", path), string(body)}
	}

	accessToken, tokenExists := respData["access_token"]
	if !tokenExists {
		return &ResponseError{http.StatusOK, fmt.Sprintf(apiBaseUrl+"%s", path), string(body)}
	}

	c.token = accessToken.(string)
	return nil
}
