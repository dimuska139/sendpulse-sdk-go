package common

import (
	"github.com/dimuska139/sendpulse-sdk-go"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

type Common struct {
	*client.Client
}

func New(httpClient *http.Client, cfg *sendpulse.Config) *Common {
	c := client.NewClient(httpClient, cfg)
	return &Common{
		Client: c,
	}
}
