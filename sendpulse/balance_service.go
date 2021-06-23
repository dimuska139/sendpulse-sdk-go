package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
	"strings"
)

type BalanceService struct {
	client *Client
}

func newBalanceService(cl *Client) *BalanceService {
	return &BalanceService{client: cl}
}

func (service *BalanceService) GetCommon(currency string) (*models.Balance, error) {
	path := "/balance"
	if currency != "" {
		path += "/" + strings.ToLower(currency)
	}

	var respData models.Balance
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}

func (service *BalanceService) GetDetailed() (*models.BalanceDetailed, error) {
	path := "/user/balance/detail"

	var respData models.BalanceDetailed
	_, err := service.client.NewRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return &respData, err
}
