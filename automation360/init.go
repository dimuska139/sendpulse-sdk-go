package automation360

import (
	"github.com/dimuska139/sendpulse-sdk-go"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

type Automation360 struct {
	*client.Client
}

func New(httpClient *http.Client, cfg *sendpulse.Config) *Automation360 {
	c := client.NewClient(httpClient, cfg)
	return &Automation360{
		Client: c,
	}
}
