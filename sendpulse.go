package sendpulse

type SendpulseClient struct {
	client *client
	Emails Emails
}

func ApiClient(config Config) (*SendpulseClient, error) {
	if config.Timeout == 0 {
		config.Timeout = 5
	}

	c := NewClient(config)

	b := books{c}
	automation := automation360{c}
	camp := campaigns{c}
	tpl := templates{c}

	spClient := &SendpulseClient{
		client: c,
		Emails: Emails{
			Books:         b,
			Automation360: automation,
			Campaigns:     camp,
			Templates:     tpl,
		},
	}

	return spClient, nil
}
