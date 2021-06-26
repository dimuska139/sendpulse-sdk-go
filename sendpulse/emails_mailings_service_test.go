package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_CreateMailing() {
	suite.mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
		  "id": 245587,
		  "status": 13, 
		  "count": 1,
		  "tariff_email_qty": 1, 
		  "overdraft_price": "0.0044", 
		  "ovedraft_currency": "RUR"
		}`)
	})

	body := "<h1>Hello!</h1>"
	addressBookID := 12345
	sendDate := models.DateTimeType(time.Now())
	mailing, err := suite.client.Emails.Mailings.CreateMailing(&models.MailingDto{
		SenderName:    "Admin",
		SenderEmail:   "test@sendpulse.com",
		Subject:       "Test message",
		Body:          &body,
		AddressBookID: &addressBookID,
		SendDate:      &sendDate,
	})
	suite.NoError(err)
	suite.Equal(245587, mailing.ID)
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_UpdateMailing() {
	suite.mux.HandleFunc("/campaigns/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPatch, r.Method)
		fmt.Fprintf(w, `{
		  "id": 245587,
		  "status": 13, 
		  "count": 1,
		  "tariff_email_qty": 1, 
		  "overdraft_price": "0.0044", 
		  "overdraft_currency": "RUR"
		}`)
	})

	body := "<h1>Hello!</h1>"
	addressBookID := 12345
	sendDate := models.DateTimeType(time.Now())
	err := suite.client.Emails.Mailings.UpdateMailing(1, &models.MailingDto{
		SenderName:    "Admin",
		SenderEmail:   "test@sendpulse.com",
		Subject:       "Test message",
		Body:          &body,
		AddressBookID: &addressBookID,
		SendDate:      &sendDate,
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_GetMailing() {
	suite.mux.HandleFunc("/campaigns/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"id": 4164892,
			"name": "{{Test}}",
			"message": {
				"sender_name": "Alex",
				"sender_email": "test@sendpulse.com",
				"subject": "{{Test}}",
				"body": "<!DOCTYPE html>\n<html>\n<head>\n<title></title>\n</head>\n<body>\n<h1>Test</h1>\n</body>\n</html>",
				"attachments": "",
				"list_id": 391289
			},
			"status": 26,
			"all_email_qty": 0,
			"tariff_email_qty": 0,
			"paid_email_qty": 0,
			"overdraft_price": 0,
			"overdraft_currency": "RUR",
			"send_date": "2020-09-28 22:22:00"
		}`)
	})

	mailing, err := suite.client.Emails.Mailings.GetMailing(1)
	suite.NoError(err)
	suite.Equal(4164892, mailing.ID)
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_GetMailings() {
	suite.mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"id": 4164892,
				"name": "{{Test}}",
				"message": {
					"sender_name": "Alex",
					"sender_email": "test@sendpulse.com",
					"subject": "{{Test}}",
					"attachments": "",
					"list_id": 391289
				},
				"status": 26,
				"all_email_qty": 0,
				"tariff_email_qty": 0,
				"paid_email_qty": 0,
				"overdraft_price": 0,
				"overdraft_currency": "RUR"
			},
			{
				"id": 7723666,
				"name": "Test",
				"message": {
					"sender_name": "Alex1",
					"sender_email": "test1@sendpulse.com",
					"subject": "Test",
					"attachments": "",
					"list_id": 1266227
				},
				"status": 14,
				"all_email_qty": 1,
				"tariff_email_qty": 1,
				"paid_email_qty": 0,
				"overdraft_price": 0,
				"overdraft_currency": "RUR"
			}
		]`)
	})

	mailing, err := suite.client.Emails.Mailings.List(10, 0)
	suite.NoError(err)
	suite.Equal(4164892, mailing[0].ID)
	suite.Equal(7723666, mailing[1].ID)
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_MailingsByAddressBook() {
	suite.mux.HandleFunc("/addressbooks/1/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"task_id": 9147533,
				"task_name": "тест",
				"task_status": 3
			},
			{
				"task_id": 9156025,
				"task_name": "Campaign_6741804_UM99",
				"task_status": 3
			}
		]`)
	})

	tasks, err := suite.client.Emails.Mailings.MailingsByAddressBook(1, 10, 0)
	suite.NoError(err)
	suite.Equal(2, len(tasks))
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_CountriesStatistics() {
	suite.mux.HandleFunc("/campaigns/1/countries", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "UA": 23,
		  "RU": 34567
		}`)
	})

	statistics, err := suite.client.Emails.Mailings.CountriesStatistics(1)
	suite.NoError(err)
	ua, ok := statistics["UA"]
	suite.True(ok)
	suite.Equal(23, ua)
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_ReferralsStatistics() {
	suite.mux.HandleFunc("/campaigns/1/referrals", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  {
			"link": "http://first_link.com",
			"count": 123454
		  },
		  {
			"link": "http://second_link.com",
			"count": 5463
		  }
		]`)
	})

	statistics, err := suite.client.Emails.Mailings.ReferralsStatistics(1)
	suite.NoError(err)
	suite.Equal(2, len(statistics))
}

func (suite *SendpulseTestSuite) TestEmailsService_MailingsService_Cancel() {
	suite.mux.HandleFunc("/campaigns/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
		   "result": true
		}`)
	})

	err := suite.client.Emails.Mailings.Cancel(1)
	suite.NoError(err)
}
