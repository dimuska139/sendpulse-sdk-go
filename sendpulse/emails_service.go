package sendpulse

type EmailsService struct {
	client       *Client
	AddressBooks *AddressBooksService
	Templates    *TemplatesService
	Senders      *SendersService
	Blacklist    *BlacklistService
}

func newEmailsService(cl *Client) *EmailsService {
	return &EmailsService{
		client:       cl,
		AddressBooks: newAddressBooksService(cl),
		Templates:    newTemplatesService(cl),
		Senders:      newSendersService(cl),
		Blacklist:    newBlacklistService(cl),
	}
}
