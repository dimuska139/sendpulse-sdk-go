package emails

import (
	"github.com/dimuska139/sendpulse-sdk-go"
	"github.com/dimuska139/sendpulse-sdk-go/client"
	"net/http"
)

// Emails: addressbooks, campaigns, templates, senders and other (see https://sendpulse.com/ru/integrations/api/bulk-email)
type Emails struct {
	*client.Client
}

func New(httpClient *http.Client, cfg *sendpulse.Config) *Emails {
	c := client.NewClient(httpClient, cfg)
	return &Emails{
		Client: c,
	}
}
