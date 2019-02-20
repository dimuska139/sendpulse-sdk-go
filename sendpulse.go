package sendpulse

import (
	"errors"
	"sync"
)

type SendpulseClient struct {
	Books         *books
	Automation360 *automation360
}

func ApiClient(apiUserId string, apiSecret string, timeout int) (*SendpulseClient, error) {
	if len(apiUserId) == 0 || len(apiSecret) == 0 {
		return nil, errors.New("client ID or Secret is empty")
	}

	c := &client{apiUserId, apiSecret, nil, timeout, nil}
	c.tokenLock = new(sync.RWMutex)

	b := &books{c}
	automation := &automation360{c}
	spClient := &SendpulseClient{b, automation}

	return spClient, nil
}
