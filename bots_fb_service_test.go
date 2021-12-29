package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestBotsFbService_GetAccount() {
	suite.mux.HandleFunc("/messenger/account", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"tariff": {
			  "branding": true,
			  "max_bots": 3,
			  "max_contacts": 1000,
			  "max_messages": 10000,
			  "max_tags": 0,
			  "max_variables": 10,
			  "max_rss": -1,
			  "code": "messengersFree10Km",
			  "is_exceeded": false,
			  "is_expired": false,
			  "expired_at": "2021-07-26T11:10:12+00:00"
			},
			"statistics": {
			  "messages": 0,
			  "bots": 0,
			  "contacts": 0,
			  "variables": 0
			}
		  }
		}`)
	})

	account, err := suite.client.Bots.Fb.GetAccount(context.Background())
	suite.NoError(err)
	suite.True(account.Tariff.Branding)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetBots() {
	suite.mux.HandleFunc("/messenger/bots", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "12345",
			"channel_data": {
			  "access_token": "qwerty",
			  "id": "12345",
			  "name": "Alex",
			  "photo": "https://vk.com/img.jpg"
			},
			"inbox": {
			  "total": 100,
			  "unread": 20
			},
			"status": 3,
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	bots, err := suite.client.Bots.Fb.GetBots(context.Background())
	suite.NoError(err)
	suite.Equal("12345", bots[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetContact() {
	suite.mux.HandleFunc("/messenger/contacts/get", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"id": "1234",
			"bot_id": "1234",
			"status": 1,
			"channel_data": {
			  "id": "1234",
			  "name": "Test",
			  "first_name": "Alex",
			  "last_name": "Alex",
			  "profile_pic": "https://vk.com/img.jpg",
			  "locale": "en",
			  "gender": "male"
			},
			"tags": [
			  "test"
			],
			"variables": {"name":"Alex"},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"unsubscribed_at": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }
		}`)
	})

	contact, err := suite.client.Bots.Fb.GetContact(context.Background(), "1")
	suite.NoError(err)
	suite.Equal("1234", contact.ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetContactsByTag() {
	suite.mux.HandleFunc("/messenger/contacts/getByTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "1234",
			"bot_id": "1234",
			"status": 1,
			"channel_data": {
			  "id": "1234",
			  "name": "Test",
			  "first_name": "Alex",
			  "last_name": "Alex",
			  "profile_pic": "https://vk.com/img.jpg",
			  "locale": "en",
			  "gender": "male"
			},
			"tags": [
			  "test"
			],
			"variables": {"name":"Alex"},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"unsubscribed_at": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	contacts, err := suite.client.Bots.Fb.GetContactsByTag(context.Background(), "tag", "bot_id")
	suite.NoError(err)
	suite.Equal("1234", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetContactsByVariable() {
	suite.mux.HandleFunc("/messenger/contacts/getByVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "1234",
			"bot_id": "1234",
			"status": 1,
			"channel_data": {
			  "id": "1234",
			  "name": "Test",
			  "first_name": "Alex",
			  "last_name": "Alex",
			  "profile_pic": "https://vk.com/img.jpg",
			  "locale": "en",
			  "gender": "male"
			},
			"tags": [
			  "test"
			],
			"variables": {"name":"Alex"},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"unsubscribed_at": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	contacts, err := suite.client.Bots.Fb.GetContactsByVariable(context.Background(), BotContactsByVariableParams{
		VariableID:    "var_id",
		VariableName:  "name",
		BotID:         "qwe123",
		VariableValue: "alex",
	})
	suite.NoError(err)
	suite.Equal("1234", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_SendTextByContact() {
	suite.mux.HandleFunc("/messenger/contacts/sendText", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.SendTextByContact(context.Background(), FbBotSendTextParams{
		ContactID:   "qwe12345",
		MessageType: "RESPONSE",
		MessageTag:  "ACCOUNT_UPDATE",
		Text:        "Hello",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_SetVariableToContact() {
	suite.mux.HandleFunc("/messenger/contacts/setVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.SetVariableToContact(context.Background(), "contactId", "variableId", "variableName", 123)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_SetTagsToContact() {
	suite.mux.HandleFunc("/messenger/contacts/setTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.SetTagsToContact(context.Background(), "contactId", []string{"tag1", "tag2"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_DeleteTagFromContact() {
	suite.mux.HandleFunc("/messenger/contacts/deleteTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.DeleteTagFromContact(context.Background(), "contactId", "tag1")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_DisableContact() {
	suite.mux.HandleFunc("/messenger/contacts/disable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.DisableContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_EnableContact() {
	suite.mux.HandleFunc("/messenger/contacts/enable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.EnableContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_DeleteContact() {
	suite.mux.HandleFunc("/messenger/contacts/delete", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.DeleteContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetPauseAutomation() {
	suite.mux.HandleFunc("/messenger/contacts/getPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
			"data": {
				"minutes": 123
			}
		}`)
	})

	p, err := suite.client.Bots.Fb.GetPauseAutomation(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal(123, p)
}

func (suite *SendpulseTestSuite) TestBotsFbService_SetPauseAutomation() {
	suite.mux.HandleFunc("/messenger/contacts/setPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.SetPauseAutomation(context.Background(), "contactId", 60)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_DeletePauseAutomation() {
	suite.mux.HandleFunc("/messenger/contacts/deletePauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.DeletePauseAutomation(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetBotVariables() {
	suite.mux.HandleFunc("/messenger/variables", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "qwerty",
			"bot_id": "qwerty",
			"name": "Alex",
			"description": "Alex qwerty",
			"type": 1,
			"value_type": 1,
			"status": 1,
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	variables, err := suite.client.Bots.Fb.GetBotVariables(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal("qwerty", variables[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetFlows() {
	suite.mux.HandleFunc("/messenger/flows", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "qwe123",
			"bot_id": "qwe12345",
			"name": "Alex",
			"status": 1,
			"triggers": [
			  {
				"id": "string",
				"name": "string"
			  }
			],
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	flows, err := suite.client.Bots.Fb.GetFlows(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal("qwe123", flows[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_RunFlow() {
	suite.mux.HandleFunc("/messenger/flows/run", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.RunFlow(context.Background(), "contactId", "flowId", map[string]interface{}{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_RunFlowByTrigger() {
	suite.mux.HandleFunc("/messenger/flows/runByTrigger", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Fb.RunFlowByTrigger(context.Background(), "contactId", "keyword", map[string]interface{}{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetBotTriggers() {
	suite.mux.HandleFunc("/messenger/triggers", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "qwe1234",
			"bot_id": "bot1",
			"flow_id": "flow1",
			"name": "string",
			"type": 1,
			"status": 1,
			"keywords": [
			  "bot1"
			],
			"execution": {
			  "interval": 0,
			  "units": 1
			},
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	triggers, err := suite.client.Bots.Fb.GetBotTriggers(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("qwe1234", triggers[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetBotChats() {
	suite.mux.HandleFunc("/messenger/chats", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"contact": {
			  "id": "string",
			  "bot_id": "string",
			  "status": 1,
			  "channel_data": {
				"id": "string",
				"name": "string",
				"first_name": "string",
				"last_name": "string",
				"profile_pic": "string",
				"locale": "string",
				"gender": "string"
			  },
			  "tags": [
				"string"
			  ],
			  "variables": {},
			  "is_chat_opened": true,
			  "last_activity_at": "2020-12-12T00:00:00+03:00",
			  "automation_paused_until": "2020-12-12T00:00:00+03:00",
			  "unsubscribed_at": "2020-12-12T00:00:00+03:00",
			  "created_at": "2020-12-12T00:00:00+03:00"
			},
			"inbox_last_message": {
			  "id": "string",
			  "contact_id": "string",
			  "bot_id": "string",
			  "campaign_id": "string",
			  "data": {
				"text": "hello"
			  },
			  "direction": 1,
			  "status": 1,
			  "delivered_at": "2020-12-12T00:00:00+03:00",
			  "opened_at": "2020-12-12T00:00:00+03:00",
			  "redirected_at": "2020-12-12T00:00:00+03:00",
			  "created_at": "2020-12-12T00:00:00+03:00"
			},
			"inbox_unread": 0
		  }]
		}`)
	})

	chats, err := suite.client.Bots.Fb.GetBotChats(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("string", chats[0].InboxLastMessage.CampaignID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_GetContactMessages() {
	suite.mux.HandleFunc("/messenger/chats/messages", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "string",
			"contact_id": "string",
			"bot_id": "string",
			"campaign_id": "string",
			"data": {
			  "text": "hello"
			},
			"direction": 1,
			"status": 1,
			"delivered_at": "2020-12-12T00:00:00+03:00",
			"opened_at": "2020-12-12T00:00:00+03:00",
			"redirected_at": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	messages, err := suite.client.Bots.Fb.GetContactMessages(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("string", messages[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsFbService_SendCampaign() {
	suite.mux.HandleFunc("/messenger/campaigns/send", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	messages := make([]FbBotCampaignMessage, 0)
	messages = append(messages, FbBotCampaignMessage{
		Type: "type",
		Data: struct {
			Text string `json:"text"`
		}{
			Text: "text",
		},
	})

	err := suite.client.Bots.Fb.SendCampaign(context.Background(), FbBotSendCampaignParams{
		Title:                   "Title",
		BotID:                   "qwe123",
		MessageTag:              "ACCOUNT_UPDATE",
		MessageNotificationType: "REGULAR",
		SendAt:                  time.Now(),
		Messages:                messages,
	})
	suite.NoError(err)
}
