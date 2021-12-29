package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetEmailInfo() {
	suite.mux.HandleFunc("/emails/test@sendpulse.com", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"book_id": 1,
				"email": "test@sendpulse.com",
				"status": 0,
				"status_explain": "New",
				"variables": [
					{
						"name": "age",
						"type": "number",
						"value": 21
					},
					{
						"name": "weight",
						"type": "string",
						"value": "99"
					}
				]
			}
		]`)
	})

	items, err := suite.client.Emails.Address.GetEmailInfo(context.Background(), "test@sendpulse.com")
	suite.NoError(err)
	suite.Equal(1, items[0].BookID)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetDetails() {
	suite.mux.HandleFunc("/emails/test@sendpulse.com/details", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"list_name": "New name 1",
				"list_id": 1266208,
				"add_date": "2021-06-18 22:01:55",
				"source": "panel"
			},
			{
				"list_name": "12345 book",
				"list_id": 1266227,
				"add_date": "2021-06-26 11:50:53",
				"source": "panel"
			}
		]`)
	})

	items, err := suite.client.Emails.Address.GetDetails(context.Background(), "test@sendpulse.com")
	suite.NoError(err)
	suite.Equal(2, len(items))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetEmailsInfo() {
	suite.mux.HandleFunc("/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"test@sendpulse.com": [
				{
					"book_id": 1,
					"status": 0,
					"variables": [
						{
							"name": "age",
							"type": "number",
							"value": 21
						},
						{
							"name": "weight",
							"type": "string",
							"value": "99"
						}
					]
				},
				{
					"book_id": 2,
					"status": 0,
					"variables": [
						{
							"name": "age",
							"type": "string",
							"value": "21"
						},
						{
							"name": "weight",
							"type": "string",
							"value": "99"
						}
					]
				}
			]
		}`)
	})

	info, err := suite.client.Emails.Address.GetEmailsInfo(context.Background(), []string{"test@sendpulse.com"})
	suite.NoError(err)

	emailInfo, ok := info["test@sendpulse.com"]
	suite.True(ok)
	suite.Equal(1, emailInfo[0].BookID)
	suite.Equal(2, emailInfo[1].BookID)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetStatisticsByCampaign() {
	suite.mux.HandleFunc("/campaigns/1/email/test@sendpulse.com", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"send_date": "2021-03-29 07:46:58",
			"global_status": 1,
			"global_status_explain": "Sent",
			"detail_status": 3,
			"detail_status_explain": "Opened"
		}`)
	})

	info, err := suite.client.Emails.Address.GetStatisticsByCampaign(context.Background(), 1, "test@sendpulse.com")
	suite.NoError(err)
	suite.Equal(1, info.GlobalStatus)
	suite.Equal(3, info.DetailStatus)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetStatisticsByAddressBook() {
	suite.mux.HandleFunc("/addressbooks/12345/emails/test@sendpulse.com", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"email": "test@sendpulse.com",
			"abook_id": "12345",
			"phone": "",
			"status": 0,
			"status_explain": "New",
			"variables": [
				{
					"name": "age",
					"type": "string",
					"value": "21"
				},
				{
					"name": "weight",
					"type": "string",
					"value": "99"
				}
			]
		}`)
	})

	info, err := suite.client.Emails.Address.GetStatisticsByAddressBook(context.Background(), 12345, "test@sendpulse.com")
	suite.NoError(err)
	suite.Equal(12345, info.AddressBookID)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_DeleteFromAllAddressBooks() {
	suite.mux.HandleFunc("/emails/test@sendpulse.com", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.Emails.Address.DeleteFromAllAddressBooks(context.Background(), "test@sendpulse.com")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetEmailStatisticsByCampaignsAndAddressBooks() {
	suite.mux.HandleFunc("/emails/test@sendpulse.com/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"statistic": {
				"sent": 9,
				"open": 2,
				"link": 0
			},
			"blacklist": false,
			"addressbooks": [
				{
					"id": 154441,
					"address_book_name": "Mailing list 1"
				},
				{
					"id": 154472,
					"address_book_name": "Mailing list 2"
				}
			]
		}`)
	})

	stat, err := suite.client.Emails.Address.GetEmailStatisticsByCampaignsAndAddressBooks(context.Background(), "test@sendpulse.com")
	suite.NoError(err)
	suite.Equal(2, len(stat.Addressbooks))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_GetEmailsStatisticsByCampaignsAndAddressBooks() {
	suite.mux.HandleFunc("/emails/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"example@yourdomain.com": {
				"sent": 21,
				"open": 11,
				"link": 3,
				"adressbooks": [
					{
						"id": 1375516,
						"name": "book1"
					},
					{
						"id": 1415158,
						"name": "book3"
					},
					{
						"id": 1649207,
						"name": "book10"
					}
				],
				"blacklist": false
			},
			"example2@yourdomain.com": {
				"sent": 1,
				"open": 1,
				"link": 0,
				"adressbooks": [
					{
						"id": 1734397,
						"name": "тест1"
					}
				],
				"blacklist": true
			}
		}`)
	})

	stat, err := suite.client.Emails.Address.GetEmailsStatisticsByCampaignsAndAddressBooks(context.Background(), []string{"example@yourdomain.com", "example2@yourdomain.com"})
	suite.NoError(err)
	email1data, email1ok := stat["example@yourdomain.com"]
	suite.True(email1ok)
	suite.Equal(3, len(email1data.Addressbooks))

	email2data, email2ok := stat["example2@yourdomain.com"]
	suite.True(email2ok)
	suite.Equal(1, len(email2data.Addressbooks))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressService_ChangeVariables() {
	suite.mux.HandleFunc("/addressbooks/1/emails/variable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	variables := make([]*Variable, 0)
	variables = append(variables, &Variable{
		Name:  "var1",
		Value: "val1",
	}, &Variable{
		Name:  "var2",
		Value: "val2",
	})
	err := suite.client.Emails.Address.ChangeVariables(context.Background(), 1, "test@sendpulse.com", variables)
	suite.NoError(err)
}
