package sendpulse_sdk_go

import (
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_WebhooksService_List() {
	suite.mux.HandleFunc("/v2/email-service/webhook", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"success": true,
			"data": [
				{
					"id": 162242,
					"user_id": 7043663,
					"url": "https://webhook.site/577e5242-bb27-4baf-8ea7-70baa6344f68",
					"action": "unsubscribe"
				},
				{
					"id": 162241,
					"user_id": 7043663,
					"url": "https://webhook.site/577e5242-bb27-4baf-8ea7-70baa6344f68",
					"action": "open"
				}
			]
		}`)
	})

	webhooks, err := suite.client.Emails.Webhooks.GetWebhooks()
	suite.NoError(err)
	suite.Equal(2, len(webhooks))
}

func (suite *SendpulseTestSuite) TestEmailsService_WebhooksService_Get() {
	suite.mux.HandleFunc("/v2/email-service/webhook/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"success": true,
			"data": {
				"id": 162242,
				"user_id": 7043663,
				"url": "https://webhook.site/577e5242-bb27-4baf-8ea7-70baa6344f68",
				"action": "unsubscribe"
			}
		}`)
	})

	webhook, err := suite.client.Emails.Webhooks.GetWebhook(1)
	suite.NoError(err)
	suite.Equal(162242, webhook.ID)
}

func (suite *SendpulseTestSuite) TestEmailsService_WebhooksService_Create() {
	suite.mux.HandleFunc("/v2/email-service/webhook/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"success": true,
			"data": [
				{
					"id": 162242,
					"user_id": 7043663,
					"url": "https://sendpulse.com",
					"action": "unsubscribe"
				},
				{
					"id": 162241,
					"user_id": 7043663,
					"url": "https://sendpulse.com",
					"action": "open"
				}
			]
		}`)
	})

	webhooks, err := suite.client.Emails.Webhooks.CreateWebhook([]string{"unsubscribe", "open"}, "https://sendpulse.com")
	suite.NoError(err)
	suite.Equal(2, len(webhooks))
}

func (suite *SendpulseTestSuite) TestEmailsService_WebhooksService_Update() {
	suite.mux.HandleFunc("/v2/email-service/webhook/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPut, r.Method)
		fmt.Fprintf(w, `{
			"success": true,
			"data": [
				true
			]
		}`)
	})

	err := suite.client.Emails.Webhooks.UpdateWebhook(1, "https://sendpulse.com")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_WebhooksService_Delete() {
	suite.mux.HandleFunc("/v2/email-service/webhook/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
			"success": true,
			"data": [
				true
			]
		}`)
	})

	err := suite.client.Emails.Webhooks.DeleteWebhook(1)
	suite.NoError(err)
}
