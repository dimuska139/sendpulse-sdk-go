package sendpulse

type EmailsService struct {
	client       *Client
	AddressBooks *AddressBooksService
	Templates    *TemplatesService
	Senders      *SendersService
	Blacklist    *BlacklistService
	Webhooks     *WebhooksService
	Address      *AddressService
	Mailings     *MailingsService
}

func newEmailsService(cl *Client) *EmailsService {
	return &EmailsService{
		client:       cl,
		AddressBooks: newAddressBooksService(cl),
		Templates:    newTemplatesService(cl),
		Senders:      newSendersService(cl),
		Blacklist:    newBlacklistService(cl),
		Webhooks:     newWebhooksService(cl),
		Address:      newAddressService(cl),
		Mailings:     newMailingsService(cl),
	}
}
