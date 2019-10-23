package sendpulse

import (
	"sync"
)

type SendpulseClient struct {
	client *client
	Emails Emails
}

func ApiClient(config Config) (*SendpulseClient, error) {
	if config.Timeout == 0 {
		config.Timeout = 5
	}

	c := &client{config, "", nil}
	c.tokenLock = new(sync.RWMutex)

	b := books{c}
	automation := automation360{c}

	spClient := &SendpulseClient{
		client: c,
		Emails: Emails{
			Books:         b,
			Automation360: automation,
		},
	}

	return spClient, nil
}
