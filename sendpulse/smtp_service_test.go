package sendpulse

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestSmtpService_Send() {
	suite.mux.HandleFunc("/smtp/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true,
			"id": "pzkic9-0afezp-fc"
		}`)
	})

	id, err := suite.client.SMTP.SendMessage(SendEmailParams{
		Html:          "<h1>Hello</h1>",
		Text:          "",
		Template:      nil,
		AutoPlainText: false,
		Subject:       "Notification",
		From: struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			Name:  "Alex",
			Email: "Brown",
		},
		To: struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			Name:  "Andy",
			Email: "Forest",
		},
		Attachments: nil,
	})
	suite.NoError(err)
	suite.Equal("pzkic9-0afezp-fc", id)
}

func (suite *SendpulseTestSuite) TestSmtpService_List() {
	suite.mux.HandleFunc("/smtp/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		{
				"id": "pzkic9-0afezp-fc",
				"sender": "tech@yourdream.com",
				"total_size": 1175,
				"sender_ip": "173.212.198.158",
				"smtp_answer_code": 250,
				"smtp_answer_code_explain": "Delivered",
				"smtp_answer_subcode": "",
				"smtp_answer_data": "anna-maria@gmail.com   H=gmail-smtp-in.l.google.com [64.233.165.27] X=TLSv1.2:ECDHE-RSA-CHACHA20-POLY1305:256 CV=yes K C=\"250 2.0.0 OK s129-v6si13409556lja.72 - gsmtp\"",
				"used_ip": "78.41.200.153",
				"recipient": "anna-maria@gmail.com",
				"subject": "Template test",
				"send_date": "2018-10-10 12:54:45",
				"tracking": {
					"click": 1,
					"open": 1,
					"link": [
						{
							"url": "wikia.com",
							"browser": "Firefox 64.0",
							"os": "Linux",
							"screen_resolution": "1920x1080",
							"ip": "77.222.152.150",
							"action_date": "2018-12-12 11:54:55"
						}
					],
					"client_info": [
						{
							"browser": "Firefox 11.0viaggpht.comGoog",
							"os": "Windows",
							"ip": "66.102.9.17",
							"action_date": "2018-12-12 11:54:54"
						}
					]
				}
			}]`)
	})

	list, err := suite.client.SMTP.GetMessages(SmtpListParams{
		Limit:     100,
		Offset:    0,
		From:      time.Now().Add(-24 * 5 * time.Hour),
		To:        time.Now(),
		Sender:    "Admin",
		Recipient: "Dev",
	})
	suite.NoError(err)
	suite.Equal(1, len(list))
}

func (suite *SendpulseTestSuite) TestSmtpService_Total() {
	suite.mux.HandleFunc("/smtp/emails/total", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"total": 25408
		}`)
	})

	total, err := suite.client.SMTP.CountMessages()
	suite.NoError(err)
	suite.Equal(25408, total)
}

func (suite *SendpulseTestSuite) TestSmtpService_Get() {
	suite.mux.HandleFunc("/smtp/emails/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "id": "pzkic9-0afezp-fc",
		  "sender": "JohnDoe@test.com",
		  "total_size": 1128,
		  "sender_ip": "127.0.0.1",
		  "smtp_answer_code": 250,
		  "smtp_answer_subcode": "0",
		  "smtp_answer_data": "Bad recipients",
		  "used_ip": "5.104.224.87",
		  "recipient": null,
		  "subject": "SendPulse :: Email confirmation",
		  "send_date": "2013-12-17 10:33:53",
		  "tracking": {
			"click": 1,
			"open": 1,
			"link": [
			  {
			"url": "http://some-url.com",
			"browser": "Chrome 29.0.1547.57",
			"os": "Linux",
			"screen_resolution": "1920x1080",
			"ip": "46.149.83.86",
			"country": "Ukraine",
			"action_date": "2013-09-30 11:27:40"
			  }
			],
			"client_info": [
			  {
			"browser": "Thunderbird 17.0.8",
			"os": "Linux",
			"ip": "46.149.83.86",
			"country": "Ukraine",
			"action_date": "2013-09-30 11:27:49"
			  }
			]
		  }
		}`)
	})

	message, err := suite.client.SMTP.GetMessage(1)
	suite.NoError(err)
	suite.Equal("pzkic9-0afezp-fc", message.ID)
	suite.Equal(1, len(message.Tracking.Link))
}

func (suite *SendpulseTestSuite) TestSmtpService_DailyBounces() {
	suite.mux.HandleFunc("/smtp/bounces/day", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"total": 2,
			"emails": [
				{
					"email_to": "gverb2016@yandex.ru",
					"sender": "no-reply@boomstream.com",
					"send_date": "2019-03-25 19:05:02",
					"subject": "[Boomstream] Покупка доступа к видео контенту",
					"smtp_answer_code": 550,
					"smtp_answer_subcode": "5.1.1",
					"smtp_answer_data": "gverb2016@yandex.ru   H=mx.yandex.ru [93.158.134.89] X=TLSv1.2:ECDHE-RSA-AES128-GCM-SHA256:128 CV=yes: SMTP error from remote mail server after RCPT TO:<gverb2016@yandex.ru>: 550 5.7.1 No such user!"
				},
				{
					"email_to": "1934a1a621@mailboxy.fun",
					"sender": "no-reply@boomstream.com",
					"send_date": "2019-03-25 15:58:00",
					"subject": "Your video files has been deleted",
					"smtp_answer_code": 552,
					"smtp_answer_subcode": "5.7.1",
					"smtp_answer_data": "1934a1a621@mailboxy.fun   H=mx5.mailboxy.fun [165.227.245.168] X=TLSv1.2:ECDHE-RSA-AES128-GCM-SHA256:128 CV=no: SMTP error from remote mail server after RCPT TO:<1934a1a621@mailboxy.fun>: 552 Mailbox limit exeeded for this email address"
				}
			],
			"request_limit": 1000,
			"found": 2
		}`)
	})

	bounces, err := suite.client.SMTP.GetDailyBounces(10, 0, time.Now())
	suite.NoError(err)
	suite.Equal(1000, bounces.RequestLimit)
	suite.Equal(2, len(bounces.Emails))
}

func (suite *SendpulseTestSuite) TestSmtpService_TotalBounces() {
	suite.mux.HandleFunc("/smtp/bounces/day/total", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"total": 3
		}`)
	})

	total, err := suite.client.SMTP.CountBounces()
	suite.NoError(err)
	suite.Equal(3, total)
}

func (suite *SendpulseTestSuite) TestSmtpService_Unsubscribe() {
	suite.mux.HandleFunc("/smtp/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	emails := make([]*SmtpUnsubscribeEmail, 0)
	emails = append(emails, &SmtpUnsubscribeEmail{
		Email:   faker.Email(),
		Comment: faker.Word(),
	})

	err := suite.client.SMTP.UnsubscribeEmails(emails)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmtpService_DeleteUnsubscribed() {
	suite.mux.HandleFunc("/smtp/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.SMTP.DeleteUnsubscribedEmails([]string{faker.Email(), faker.Email()})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmtpService_UnsubscribedList() {
	suite.mux.HandleFunc("/smtp/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[  
		   {  
			  "email":"4090797@mail.ru",
			  "unsubscribe_by_link":1,
			  "unsubscribe_by_user":0,
			  "spam_complaint":1,
			  "date":"2018-11-24 19:19:01"
		   },
		   {  
			  "email":"4lik@rambler.ru",
			  "unsubscribe_by_link":1,
			  "unsubscribe_by_user":0,
			  "spam_complaint":1,
			  "date":"2019-03-20 16:47:01"
		   }
		]`)
	})

	unsubscribed, err := suite.client.SMTP.GetUnsubscribedEmails(UnsubscribedListParams{
		Limit:  10,
		Offset: 0,
		Date:   time.Now(),
	})
	suite.NoError(err)
	suite.Equal(2, len(unsubscribed))
}

func (suite *SendpulseTestSuite) TestSmtpService_SendersIPs() {
	suite.mux.HandleFunc("/smtp/ips", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  "127.0.0.1",
		  "192.168.0.1"
		]`)
	})

	ips, err := suite.client.SMTP.GetSendersIPs()
	suite.NoError(err)
	suite.Equal(2, len(ips))
}

func (suite *SendpulseTestSuite) TestSmtpService_SendersEmails() {
	suite.mux.HandleFunc("/smtp/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  "test@sendpulse.com",
		  "dev@sendpulse.com"
		]`)
	})

	emails, err := suite.client.SMTP.GetSendersEmails()
	suite.NoError(err)
	suite.Equal(2, len(emails))
}

func (suite *SendpulseTestSuite) TestSmtpService_AllowedDomains() {
	suite.mux.HandleFunc("/smtp/domains", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  "@sendpulse.com",
		  "@test.com"
		]`)
	})

	domains, err := suite.client.SMTP.GetAllowedDomains()
	suite.NoError(err)
	suite.Equal(2, len(domains))
}

func (suite *SendpulseTestSuite) TestSmtpService_AddDomain() {
	suite.mux.HandleFunc("/smtp/domains", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.SMTP.AddDomain(faker.Email())
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestSmtpService_VerifyDomain() {
	email := faker.Email()

	suite.mux.HandleFunc(fmt.Sprintf("/domains/%s", email), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.SMTP.VerifyDomain(email)
	suite.NoError(err)
}
