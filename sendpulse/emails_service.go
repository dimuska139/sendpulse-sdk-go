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
	Validator    *ValidatorService
}

func newEmailsService(cl *Client) *EmailsService {
	return &EmailsService{
		client:       cl,
		AddressBooks: newAddressBooksService(cl),
		Mailings:     newMailingsService(cl),
		Templates:    newTemplatesService(cl),
		Senders:      newSendersService(cl),
		Address:      newAddressService(cl),
		Blacklist:    newBlacklistService(cl),
		Webhooks:     newWebhooksService(cl),
		Validator:    newValidatorService(cl),
	}
}
