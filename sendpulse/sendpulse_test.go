package sendpulse

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type SendpulseTestSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
	mux    *http.ServeMux
}

// Rewrites scheme to http to avoid TLS cert issues
type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func (suite *SendpulseTestSuite) BeforeTest(suiteName, testName string) {
	suite.mux = http.NewServeMux()
	suite.mux.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"access_token": "12345"}`)
	})

	suite.server = httptest.NewServer(suite.mux)

	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(suite.server.URL)
		},
	}}

	httpClient := &http.Client{
		Transport: transport,
	}
	config := &Config{
		UserID: "uid",
		Secret: "secret",
	}
	suite.client = NewClient(httpClient, config)
}

func (suite *SendpulseTestSuite) AfterTest(suiteName, testName string) {
	suite.server.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SendpulseTestSuite))
}
