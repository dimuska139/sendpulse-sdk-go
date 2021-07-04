package sendpulse

type BotsService struct {
	client   *Client
	Fb       *BotsFbService
	Vk       *BotsVkService
	Telegram *BotsTelegramService
}

func newBotsService(cl *Client) *BotsService {
	return &BotsService{
		client:   cl,
		Fb:       newBotsFbService(cl),
		Vk:       newBotsVkService(cl),
		Telegram: newBotsTelegramService(cl),
	}
}
