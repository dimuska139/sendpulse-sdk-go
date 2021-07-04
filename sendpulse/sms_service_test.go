package sendpulse

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (suite *SendpulseTestSuite) TestSmsService_AddPhones() {
	suite.mux.HandleFunc("/sms/numbers", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		   "result": true,
		   "counters": {
			 "added": 0,
			 "exceptions": 0,
			 "exists": 83
		   }
		 }`)
	})

	statistics, err := suite.client.SMS.AddPhones(123, []string{"380632631234", "38063333333"})
	suite.NoError(err)
	suite.Equal(83, statistics.Exists)
}

func (suite *SendpulseTestSuite) TestSmsService_AddPhonesWithVariables() {
	suite.mux.HandleFunc("/sms/numbers/variables", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		b, _ := io.ReadAll(r.Body)
		suite.Equal(`{"addressBookId":123,"phones":{"380632631234":[[{"name":"vvv","type":"type","value":"value"}]]}}`, strings.TrimSuffix(string(b), "\n"))

		fmt.Fprintf(w, `{
		   "result": true,
		   "counters": {
			 "added": 0,
			 "exceptions": 0,
			 "exists": 83
		   }
		 }`)
	})

	items := make([]*PhoneWithVariable, 0)
	items = append(items, &PhoneWithVariable{
		Phone: "380632631234",
		Variables: []SmsVariable{
			{
				Name:  "vvv",
				Type:  "type",
				Value: "value",
			},
		},
	})
	statistics, err := suite.client.SMS.AddPhonesWithVariables(123, items)
	suite.NoError(err)
	suite.Equal(83, statistics.Exists)
}

func (suite *SendpulseTestSuite) TestSmsService_UpdateVariablesSingle() {
	addressBookID := 24
	suite.mux.HandleFunc(fmt.Sprintf("/addressbooks/%d/phones/variable", addressBookID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		   "result": true
		 }`)
	})
	err := suite.client.SMS.UpdateVariablesSingle(addressBookID, "380632631234", []SmsVariable{{
		Name:  "vvv",
		Value: "value",
	}})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmsService_UpdateVariablesMultiple() {
	suite.mux.HandleFunc("/sms/numbers", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPut, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"counters": {
				"updated": 4
			}
		}`)
	})
	err := suite.client.SMS.UpdateVariablesMultiple(12345, []string{"380632631234"}, []SmsVariable{{
		Name:  "vvv",
		Value: "value",
	}})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmsService_DeletePhones() {
	suite.mux.HandleFunc("/sms/numbers", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"counters": {
				"added": 0,
     			"exists": 3
			}
		}`)
	})
	err := suite.client.SMS.DeletePhones(12345, []string{"380632631234"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmsService_GetPhoneInfo() {
	addressBookID := 12345
	phone := "380632631234"
	suite.mux.HandleFunc(fmt.Sprintf("/sms/numbers/info/%d/%s", addressBookID, phone), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"data": {
				"email": false,
				"status": 0,
				"variables": {
					"Vvv": "test"
				},
				"added": "2021-06-30 18:08:07"
			}
		}`)
	})
	info, err := suite.client.SMS.GetPhoneInfo(addressBookID, phone)
	suite.NoError(err)
	suite.Equal("test", info.Variables["Vvv"])
}

func (suite *SendpulseTestSuite) TestSmsService_AddToBlacklist() {
	suite.mux.HandleFunc("/sms/black_list", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"counters": {
				"added": 2,
				"exists": 3 
			}
		}`)
	})
	err := suite.client.SMS.AddToBlacklist([]string{"380632631234"}, "Invalid phone numbers")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmsService_RemoveFromBlacklist() {
	suite.mux.HandleFunc("/sms/black_list", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"counters": {
				"removed":3
			}
		}`)
	})
	err := suite.client.SMS.RemoveFromBlacklist([]string{"380632631234"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmsService_BlacklistPhones() {
	suite.mux.HandleFunc("/sms/black_list/by_numbers", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"data": [
				{
					"phone": 79217451232,
					"description": "test",
					"add_date": "2021-07-01 08:10:08"
				}
			]
		}`)
	})
	items, err := suite.client.SMS.GetBlacklistedPhones([]string{"380632631234"})
	suite.NoError(err)
	suite.Equal("79217451232", items[0].Phone)
}

func (suite *SendpulseTestSuite) TestSmsService_CreateCampaignByAddressBook() {
	suite.mux.HandleFunc("/sms/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"campaign_id": 2623084
		}`)
	})
	campaignID, err := suite.client.SMS.CreateCampaignByMailingList(CreateSmsCampaignByAddressBookParams{
		Sender:        "Alex",
		MailingListID: 12345,
		Body:          "Hello!",
		Route:         nil,
		Date:          DateTimeType(time.Now()),
	})
	suite.NoError(err)
	suite.Equal(2623084, campaignID)
}

func (suite *SendpulseTestSuite) TestSmsService_CreateCampaignByPhones() {
	suite.mux.HandleFunc("/sms/send", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"campaign_id": 2623085,
			"counters": {
				"exceptions": 0,
				"sends": 3
			}
		}`)
	})

	campaignID, err := suite.client.SMS.CreateCampaignByPhones(CreateSmsCampaignByPhonesParams{
		Sender: "Alex",
		Phones: []string{"79217451232"},
		Body:   "Hello",
		Route:  nil,
		Date:   DateTimeType(time.Now()),
	})
	suite.NoError(err)
	suite.Equal(2623085, campaignID)
}

func (suite *SendpulseTestSuite) TestSmsService_CampaignsList() {
	dateFrom := time.Now()
	dateTo := time.Now()

	suite.mux.HandleFunc("/sms/campaigns/list", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		   "result": true,
		   "data": [ {
			 "id": 2136035,
			 "address_book_id": 0,
			 "company_price": 0.81,
			 "company_currency": "UAH",
			 "send_date": "2017-01-18 08:15:18",
			 "date_created": "2017-01-18 08:15:18",
			 "sender_mail_address": "",
			 "sender_mail_name": "",
			 "external_stat": []
		   }, 
		   {
		   "id": 2136036,
		   "address_book_id": 0,
		   "company_price": 0.27,
		   "company_currency": "UAH",
		   "send_date": "2017-01-18 11:59:52",
		   "date_created": "2017-01-18 11:59:52",
		   "sender_mail_address": "",
		   "sender_mail_name": "",
		   "external_stat": []
		   }
		   ]
		 }`)
	})

	items, err := suite.client.SMS.GetCampaigns(dateFrom, dateTo)
	suite.NoError(err)
	suite.Equal(2, len(items))
}

func (suite *SendpulseTestSuite) TestSmsService_CampaignInfo() {
	campaignID := 23
	suite.mux.HandleFunc(fmt.Sprintf("/sms/campaigns/info/%d", campaignID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{  
			"result":true,
			"data":{  
				"id":7520226,
				"address_book_id":12345,
				"currency":"UAH",
				"company_price":0.54,
				"send_date":"2018-09-13 11:03:22",
				"date_created":"2018-09-13 11:03:22",
				"sender_name":"AFK",
				"task_phones_info":[  
					{  
						"phone":380956045455,
						"status":2,
						"status_explain":"Delivered",
						"сountry_code":"UA",
						"money_spent":0.27
					},
					{  
						"phone":380985587288,
						"status":2,
						"status_explain":"Delivered",
						"сountry_code":"UA",
						"money_spent":0.27
					}
				]
			}
		}`)
	})

	info, err := suite.client.SMS.GetCampaignInfo(campaignID)
	suite.NoError(err)
	suite.Equal(7520226, info.ID)
}

func (suite *SendpulseTestSuite) TestSmsService_CancelCampaign() {
	campaignID := 23
	suite.mux.HandleFunc(fmt.Sprintf("/sms/campaigns/cancel/%d", campaignID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPut, r.Method)
		fmt.Fprintf(w, `{  
			"result":true
		}`)
	})
	err := suite.client.SMS.CancelCampaign(campaignID)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmsService_GetCampaignCost() {
	suite.mux.HandleFunc("/sms/campaigns/cost", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		   "result": true,
		   "data": {
			 "price": 14.679,
			 "currency": "UAH" 
		   }
		}`)
	})

	cost, err := suite.client.SMS.GetCampaignCost(SmsCampaignCostParams{
		AddressBookID: 12345,
		Phones:        []string{"79217451232"},
		Body:          "Hello",
		Sender:        "Alex",
		Route:         map[string]string{"UA": "national", "BY": "internationa;"},
	})
	suite.NoError(err)
	suite.Equal("UAH", cost.Currency)
}

func (suite *SendpulseTestSuite) TestSmsService_GetSenders() {
	suite.mux.HandleFunc("/sms/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"id": 18021,
				"sender": "apicheck",
				"country": "Ukraine",
				"country_code": "UA",
				"status": 2,
				"status_explain": "Rejected"
			},
			{
				"id": 4786,
				"sender": "iQA",
				"country": "Ukraine",
				"country_code": "UA",
				"status": 1,
				"status_explain": "Active"
			},
			{
				"id": 4787,
				"sender": "iQA",
				"country": "Ukraine",
				"country_code": "UA",
				"status": 1,
				"status_explain": "Active"
			},
			{
				"id": 18027,
				"sender": "rejectit",
				"country": "Ukraine",
				"country_code": "UA",
				"status": 0,
				"status_explain": "On moderation"
			}
		]`)
	})

	senders, err := suite.client.SMS.GetSenders()
	suite.NoError(err)
	suite.Equal(4, len(senders))
}

func (suite *SendpulseTestSuite) TestSmsService_DeleteCampaign() {
	suite.mux.HandleFunc("/sms/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{  
			"result":true
		}`)
	})

	err := suite.client.SMS.DeleteCampaign(2)
	suite.NoError(err)
}
