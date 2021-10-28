package sendpulse_sdk_go

import (
	"fmt"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestBotsVkService_GetAccount() {
	suite.mux.HandleFunc("/vk/account", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"plan": {
			  "branding": true,
			  "max_bots": 3,
			  "max_contacts": 1000,
			  "max_messages": 10000,
			  "max_tags": 0,
			  "max_variables": 10,
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

	account, err := suite.client.Bots.Vk.GetAccount()
	suite.NoError(err)
	suite.True(account.Plan.Branding)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetBots() {
	suite.mux.HandleFunc("/vk/bots", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "12345",
			"channel_data": {
			  "access_token": "qwerty",
			  "id": "12345",
			  "name": "Alex",
			  "username": "alex23"
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

	bots, err := suite.client.Bots.Vk.GetBots()
	suite.NoError(err)
	suite.Equal("12345", bots[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetContact() {
	suite.mux.HandleFunc("/vk/contacts/get", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"id": "1234",
			"bot_id": "1234",
			"status": 1,
			"channel_data": {
			  "username": "alex",
			  "name": "Test",
			  "first_name": "Alex",
			  "last_name": "Alex",
			  "language_code": "en"
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

	contact, err := suite.client.Bots.Vk.GetContact("1")
	suite.NoError(err)
	suite.Equal("1234", contact.ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetContactsByTag() {
	suite.mux.HandleFunc("/vk/contacts/getByTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "1234",
			"bot_id": "1234",
			"status": 1,
			"channel_data": {
				"group_id": 0,
				"is_member": true,
				"name": "Test",
				"data": {}
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

	contacts, err := suite.client.Bots.Vk.GetContactsByTag("tag", "bot_id")
	suite.NoError(err)
	suite.Equal("1234", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetContactsByVariable() {
	suite.mux.HandleFunc("/vk/contacts/getByVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "1234",
			"bot_id": "1234",
			"status": 1,
			"channel_data": {
			  "group_id": 0,
				"is_member": true,
				"name": "Test",
				"data": {}
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

	contacts, err := suite.client.Bots.Vk.GetContactsByVariable(BotContactsByVariableParams{
		VariableID:    "var_id",
		VariableName:  "name",
		BotID:         "qwe123",
		VariableValue: "alex",
	})
	suite.NoError(err)
	suite.Equal("1234", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_SendTextByContact() {
	suite.mux.HandleFunc("/vk/contacts/sendText", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.SendTextByContact("qwe12345", "hello")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_SetVariableToContact() {
	suite.mux.HandleFunc("/vk/contacts/setVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.SetVariableToContact("contactId", "variableId", "variableName", 123)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_SetTagsToContact() {
	suite.mux.HandleFunc("/vk/contacts/setTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.SetTagsToContact("contactId", []string{"tag1", "tag2"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_DeleteTagFromContact() {
	suite.mux.HandleFunc("/vk/contacts/deleteTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.DeleteTagFromContact("contactId", "tag1")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_DisableContact() {
	suite.mux.HandleFunc("/vk/contacts/disable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.DisableContact("contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_EnableContact() {
	suite.mux.HandleFunc("/vk/contacts/enable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.EnableContact("contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_DeleteContact() {
	suite.mux.HandleFunc("/vk/contacts/delete", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.DeleteContact("contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetPauseAutomation() {
	suite.mux.HandleFunc("/vk/contacts/getPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
			"data": {
				"minutes": 123
			}
		}`)
	})

	p, err := suite.client.Bots.Vk.GetPauseAutomation("contactId")
	suite.NoError(err)
	suite.Equal(123, p)
}

func (suite *SendpulseTestSuite) TestBotsVkService_SetPauseAutomation() {
	suite.mux.HandleFunc("/vk/contacts/setPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.SetPauseAutomation("contactId", 60)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_DeletePauseAutomation() {
	suite.mux.HandleFunc("/vk/contacts/deletePauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.DeletePauseAutomation("contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetBotVariables() {
	suite.mux.HandleFunc("/vk/variables", func(w http.ResponseWriter, r *http.Request) {
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

	variables, err := suite.client.Bots.Vk.GetBotVariables("contactId")
	suite.NoError(err)
	suite.Equal("qwerty", variables[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetFlows() {
	suite.mux.HandleFunc("/vk/flows", func(w http.ResponseWriter, r *http.Request) {
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

	flows, err := suite.client.Bots.Vk.GetFlows("contactId")
	suite.NoError(err)
	suite.Equal("qwe123", flows[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_RunFlow() {
	suite.mux.HandleFunc("/vk/flows/run", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.RunFlow("contactId", "flowId", map[string]interface{}{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_RunFlowByTrigger() {
	suite.mux.HandleFunc("/vk/flows/runByTrigger", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Vk.RunFlowByTrigger("contactId", "keyword", map[string]interface{}{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetBotTriggers() {
	suite.mux.HandleFunc("/vk/triggers", func(w http.ResponseWriter, r *http.Request) {
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

	triggers, err := suite.client.Bots.Vk.GetBotTriggers("bot")
	suite.NoError(err)
	suite.Equal("qwe1234", triggers[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetBotChats() {
	suite.mux.HandleFunc("/vk/chats", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"contact": {
			  "id": "string",
			  "bot_id": "string",
			  "status": 1,
			  "channel_data": {
				"group_id": 0,
				"is_member": true,
				"name": "Test",
				"data": {}
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

	chats, err := suite.client.Bots.Vk.GetBotChats("bot")
	suite.NoError(err)
	suite.Equal("string", chats[0].InboxLastMessage.CampaignID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_GetContactMessages() {
	suite.mux.HandleFunc("/vk/chats/messages", func(w http.ResponseWriter, r *http.Request) {
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

	messages, err := suite.client.Bots.Vk.GetContactMessages("bot")
	suite.NoError(err)
	suite.Equal("string", messages[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsVkService_SendCampaign() {
	suite.mux.HandleFunc("/vk/campaigns/send", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	messages := make([]VkBotCampaignMessage, 0)
	messages = append(messages, VkBotCampaignMessage{
		Type: "type",
		Message: struct {
			Text string `json:"text"`
		}{
			Text: "text",
		},
	})

	err := suite.client.Bots.Vk.SendCampaign(VkBotSendCampaignParams{
		Title:    "Title",
		BotID:    "qwe123",
		SendAt:   time.Now(),
		Messages: messages,
	})
	suite.NoError(err)
}
