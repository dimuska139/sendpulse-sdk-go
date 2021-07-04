package sendpulse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

const apiBaseUrl = "https://api.sendpulse.com"

type SendpulseError struct {
	HttpCode int
	Url      string
	Body     string
	Message  string
}

func (e *SendpulseError) Error() string {
	return fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message)
}

type Client struct {
	client    *http.Client
	config    *Config
	token     string
	tokenLock *sync.RWMutex
	Emails    *EmailsService
	Balance   *BalanceService
	SMTP      *SmtpService
	Push      *PushService
	Sms       *SmsService
	Viber     *ViberService
	VkOk      *VkOkService
	Bots      *BotsService
}

func NewClient(client *http.Client, config *Config) *Client {
	cl := &Client{
		client:    client,
		config:    config,
		token:     "",
		tokenLock: new(sync.RWMutex),
	}
	cl.Emails = newEmailsService(cl)
	cl.Balance = newBalanceService(cl)
	cl.SMTP = newSmtpService(cl)
	cl.Push = newPushService(cl)
	cl.Sms = newSmsService(cl)
	cl.Viber = newViberService(cl)
	cl.VkOk = newVkOkService(cl)
	cl.Bots = newBotsService(cl)
	return cl
}

func (c *Client) getToken() (string, error) {
	c.tokenLock.RLock()
	token := c.token
	c.tokenLock.RUnlock()

	if token != "" {
		return token, nil
	}

	data := make(map[string]interface{})
	data["grant_type"] = "client_credentials"
	data["client_id"] = c.config.UserID
	data["client_secret"] = c.config.Secret
	path := "/oauth/access_token"

	var respData struct {
		AccessToken string `json:"access_token"`
	}

	_, err := c.NewRequest(http.MethodPost, path, data, &respData, false)
	if err != nil {
		return "", err
	}

	c.tokenLock.Lock()
	c.token = respData.AccessToken
	token = respData.AccessToken
	c.tokenLock.Unlock()

	return token, nil
}

func (c *Client) clearToken() {
	c.tokenLock.Lock()
	c.token = ""
	c.tokenLock.Unlock()
}

func (c *Client) NewRequest(method string, path string, body interface{}, result interface{}, useToken bool) (*http.Response, error) {
	fullPath := apiBaseUrl + path
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}
	req, e := http.NewRequest(method, fullPath, buf)
	if e != nil {
		return nil, e
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if useToken {
		token, err := c.getToken()
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &SendpulseError{http.StatusServiceUnavailable, path, "", err.Error()}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized && useToken {
		c.clearToken()
		respData, err := c.NewRequest(method, path, body, result, useToken)
		if err != nil {
			return nil, err
		}
		return respData, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &SendpulseError{resp.StatusCode, path, string(respBody), err.Error()}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &SendpulseError{resp.StatusCode, path, string(respBody), ""}
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, &SendpulseError{resp.StatusCode, path, string(respBody), err.Error()}
	}

	return resp, nil
}

func (c *Client) NewFormDataRequest(path string, buffer *bytes.Buffer, contentType string, result interface{}, useToken bool) (*http.Response, error) {
	fullPath := apiBaseUrl + path
	req, e := http.NewRequest(http.MethodPost, fullPath, buffer)
	if e != nil {
		return nil, e
	}

	req.Header.Set("Content-Type", contentType)

	if useToken {
		token, err := c.getToken()
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &SendpulseError{http.StatusServiceUnavailable, path, "", err.Error()}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized && useToken {
		c.clearToken()
		respData, err := c.NewFormDataRequest(path, buffer, contentType, result, useToken)
		if err != nil {
			return nil, err
		}
		return respData, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &SendpulseError{resp.StatusCode, path, string(respBody), err.Error()}
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, &SendpulseError{resp.StatusCode, path, string(respBody), ""}
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, &SendpulseError{resp.StatusCode, path, string(respBody), err.Error()}
	}

	return resp, nil
}
