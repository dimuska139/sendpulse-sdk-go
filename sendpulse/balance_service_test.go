package sendpulse

import (
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestBalanceService_GetCommon() {
	suite.mux.HandleFunc("/balance/rub", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"currency": "RUR",
			"balance_currency": 0
		}`)
	})

	balance, err := suite.client.Balance.GetCommon("RUB")
	suite.NoError(err)
	suite.Equal("RUR", balance.Currency)
}

func (suite *SendpulseTestSuite) TestBalanceService_GetDetailed() {
	suite.mux.HandleFunc("/user/balance/detail", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"balance": {
				"main": "0.00",
				"bonus": "0.00",
				"currency": "RUR"
			},
			"email": {
				"tariff_name": "Подписка 500",
				"finished_time": "2021-07-15 15:39:02",
				"emails_left": 15000,
				"maximum_subscribers": 500,
				"current_subscribers": 0
			},
			"smtp": {
				"tariff_name": "SMTP Free",
				"end_date": "2020-11-08 09:33:52",
				"auto_renew": 0
			},
			"push": {
				"tariff_name": "Бесплатно 10 000",
				"end_date": "2021-07-16 16:46:21",
				"auto_renew": 1
			}
		}`)
	})

	balance, err := suite.client.Balance.GetDetailed()
	suite.NoError(err)
	suite.Equal("RUR", balance.Balance.Currency)
}
