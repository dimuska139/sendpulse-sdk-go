package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetAccount() {
	suite.mux.HandleFunc("/whatsapp/account", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"plan": {
			  "code": "Plan code",
			  "max_bots": 0,
			  "max_contacts": -1,
			  "max_messages": 0,
			  "max_tags": 0,
			  "max_variables": 0,
			  "branding": true,
			  "is_exceeded": true,
			  "is_expired": true,
			  "expired_at": "2020-12-12T00:00:00+03:00"
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

	account, err := suite.client.Bots.WhatsApp.GetAccount(context.Background())
	suite.NoError(err)
	suite.Equal("Plan code", account.Plan.Code)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetBots() {
	suite.mux.HandleFunc("/whatsapp/bots", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "12345",
			"channel_data": {
			  "name": "Alex",
			  "phone": "89221112233"
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

	bots, err := suite.client.Bots.WhatsApp.GetBots(context.Background())
	suite.NoError(err)
	suite.Equal("12345", bots[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_CreateContact() {
	suite.mux.HandleFunc("/whatsapp/contacts", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"id": "12345",
			"bot_id": "6789",
			"status": 1,
			"channel_data": {
			  "username": "alex1981",
			  "first_name": "Alex",
			  "last_name": "Pavlov",
			  "name": "Aleksey",
			  "language_code": "ru"
			},
			"tags": [
			  "tag1",
			  "tag2"
			],
			"variables": {},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }
		}`)
	})

	contact, err := suite.client.Bots.WhatsApp.CreateContact(context.Background(), "6789", "89221112233", "Aleksey")
	suite.NoError(err)
	suite.Equal("12345", contact.ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/get", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"id": "12345",
			"bot_id": "6789",
			"status": 1,
			"channel_data": {
			  "username": "alex1981",
			  "first_name": "Alex",
			  "last_name": "Pavlov",
			  "name": "Aleksey",
			  "language_code": "ru"
			},
			"tags": [
			  "tag1",
			  "tag2"
			],
			"variables": {},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }
		}`)
	})

	contact, err := suite.client.Bots.WhatsApp.GetContact(context.Background(), "1")
	suite.NoError(err)
	suite.Equal("12345", contact.ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetContactsByPhone() {
	suite.mux.HandleFunc("/whatsapp/contacts/getByPhone", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "12345",
			"bot_id": "6789",
			"status": 1,
			"channel_data": {
			  "username": "alex1981",
			  "first_name": "Alex",
			  "last_name": "Pavlov",
			  "name": "Aleksey",
			  "language_code": "ru"
			},
			"tags": [
			  "tag1",
			  "tag2"
			],
			"variables": {},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	contacts, err := suite.client.Bots.WhatsApp.GetContactsByPhone(context.Background(), "89991112233", "6789")
	suite.NoError(err)
	suite.Equal("12345", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetContactsByTag() {
	suite.mux.HandleFunc("/whatsapp/contacts/getByTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "12345",
			"bot_id": "6789",
			"status": 1,
			"channel_data": {
			  "username": "alex1981",
			  "first_name": "Alex",
			  "last_name": "Pavlov",
			  "name": "Aleksey",
			  "language_code": "ru"
			},
			"tags": [
			  "tag1",
			  "tag2"
			],
			"variables": {},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	contacts, err := suite.client.Bots.WhatsApp.GetContactsByTag(context.Background(), "tag1", "6789")
	suite.NoError(err)
	suite.Equal("12345", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetContactsByVariable() {
	suite.mux.HandleFunc("/whatsapp/contacts/getByVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "12345",
			"bot_id": "6789",
			"status": 1,
			"channel_data": {
			  "username": "alex1981",
			  "first_name": "Alex",
			  "last_name": "Pavlov",
			  "name": "Aleksey",
			  "language_code": "ru"
			},
			"tags": [
			  "tag1",
			  "tag2"
			],
			"variables": {},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	contacts, err := suite.client.Bots.WhatsApp.GetContactsByVariable(context.Background(), BotContactsByVariableParams{
		VariableName:  "name",
		VariableValue: "Aleksey",
		VariableID:    "12345",
		BotID:         "6789",
	})
	suite.NoError(err)
	suite.Equal("12345", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendByContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/send", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	msg := WhatsAppMessage{
		Type: "image",
		Image: &struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		}{
			Link:    faker.URL(),
			Caption: faker.Word(),
		},
	}
	err := suite.client.Bots.WhatsApp.SendByContact(context.Background(), "12345", &msg)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendByPhone() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendByPhone", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	msg := WhatsAppMessage{
		Type: "image",
		Image: &struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		}{
			Link:    faker.URL(),
			Caption: faker.Word(),
		},
	}
	err := suite.client.Bots.WhatsApp.SendByPhone(context.Background(), "12345", "89221112345", &msg)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendTemplate() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendTemplate", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SendTemplate(context.Background(), "12345", "first_template", "ru")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendTemplateWithVariables() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendTemplate", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SendTemplateWithVariables(context.Background(), "12345", "first_template", "ru", []string{
		"текст переменной",
		"{{last_name}}",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendTemplateWithImage() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendTemplate", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SendTemplateWithImage(context.Background(), "12345", "first_template", "ru", faker.Word())
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendTemplateByPhone() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendTemplateByPhone", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SendTemplateByPhone(context.Background(), "12345", "89991112233", "first_template", "ru")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendTemplateByPhoneWithVariables() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendTemplateByPhone", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SendTemplateByPhoneWithVariables(context.Background(), "12345",
		"89991112233",
		"first_template",
		"ru",
		[]string{
			"текст переменной",
			"{{last_name}}",
		})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendTemplateByPhoneWithImage() {
	suite.mux.HandleFunc("/whatsapp/contacts/sendTemplateByPhone", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SendTemplateByPhoneWithImage(context.Background(), "12345",
		"89991112233",
		"first_template",
		"ru",
		faker.URL())
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SetVariableToContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/setVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SetVariableToContact(context.Background(), "contactId", "variableId", "variableName", 123)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SetTagsToContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/setTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SetTagsToContact(context.Background(), "contactId", []string{"tag1", "tag2"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_DeleteTagFromContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/deleteTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.DeleteTagFromContact(context.Background(), "contactId", "tag1")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_DisableContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/disable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.DisableContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_EnableContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/enable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.EnableContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_DeleteContact() {
	suite.mux.HandleFunc("/whatsapp/contacts/delete", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.DeleteContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetPauseAutomation() {
	suite.mux.HandleFunc("/whatsapp/contacts/getPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
			"data": {
				"minutes": 123
			}
		}`)
	})

	p, err := suite.client.Bots.WhatsApp.GetPauseAutomation(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal(123, p)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SetPauseAutomation() {
	suite.mux.HandleFunc("/whatsapp/contacts/setPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.SetPauseAutomation(context.Background(), "contactId", 60)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_DeletePauseAutomation() {
	suite.mux.HandleFunc("/whatsapp/contacts/deletePauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.DeletePauseAutomation(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetBotVariables() {
	suite.mux.HandleFunc("/whatsapp/variables", func(w http.ResponseWriter, r *http.Request) {
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

	variables, err := suite.client.Bots.WhatsApp.GetBotVariables(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal("qwerty", variables[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetFlows() {
	suite.mux.HandleFunc("/whatsapp/flows", func(w http.ResponseWriter, r *http.Request) {
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

	flows, err := suite.client.Bots.WhatsApp.GetFlows(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal("qwe123", flows[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_RunFlow() {
	suite.mux.HandleFunc("/whatsapp/flows/run", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.RunFlow(context.Background(), "contactId", "flowId", map[string]interface{}{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_RunFlowByTrigger() {
	suite.mux.HandleFunc("/whatsapp/flows/runByTrigger", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.WhatsApp.RunFlowByTrigger(context.Background(), "contactId", "keyword", map[string]interface{}{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetBotTriggers() {
	suite.mux.HandleFunc("/whatsapp/triggers", func(w http.ResponseWriter, r *http.Request) {
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

	triggers, err := suite.client.Bots.WhatsApp.GetBotTriggers(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("qwe1234", triggers[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetBotChats() {
	suite.mux.HandleFunc("/whatsapp/chats", func(w http.ResponseWriter, r *http.Request) {
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

	chats, err := suite.client.Bots.WhatsApp.GetBotChats(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("string", chats[0].InboxLastMessage.CampaignID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetContactMessages() {
	suite.mux.HandleFunc("/whatsapp/chats/messages", func(w http.ResponseWriter, r *http.Request) {
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

	messages, err := suite.client.Bots.WhatsApp.GetContactMessages(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("string", messages[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendCampaign() {
	suite.mux.HandleFunc("/whatsapp/campaigns/send", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	messages := make([]WhatsAppMessage, 1)
	messages[0] = WhatsAppMessage{
		Type: "image",
		Image: &struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		}{
			Link:    faker.URL(),
			Caption: faker.Word(),
		},
	}

	err := suite.client.Bots.WhatsApp.SendCampaign(context.Background(), WhatsAppBotSendCampaignParams{
		Title:    "Title",
		BotID:    "qwe123",
		SendAt:   time.Now(),
		Messages: messages,
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_SendCampaignByTemplate() {
	suite.mux.HandleFunc("/whatsapp/campaigns/sendTemplate", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	messages := make([]WhatsAppMessage, 1)
	messages[0] = WhatsAppMessage{
		Type: "image",
		Image: &struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		}{
			Link:    faker.URL(),
			Caption: faker.Word(),
		},
	}

	err := suite.client.Bots.WhatsApp.SendCampaignByTemplate(context.Background(), WhatsAppBotSendCampaignByTemplateParams{
		TemplateName: "first_template",
		LanguageCode: "en",
		Title:        "Title",
		BotID:        "qwe123",
		SendAt:       time.Now(),
		Messages:     messages,
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsWhatsAppService_GetTemplates() {
	suite.mux.HandleFunc("/whatsapp/templates", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "123",
			"bot_id": "456",
			"namespace": "tpls",
			"category": "category",
			"components": [
			  {}
			],
			"language": "ru",
			"name": "first_template",
			"rejected_reason": "",
			"status": "APPROVED",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	messages := make([]WhatsAppMessage, 1)
	messages[0] = WhatsAppMessage{
		Type: "image",
		Image: &struct {
			Link    string `json:"link"`
			Caption string `json:"caption"`
		}{
			Link:    faker.URL(),
			Caption: faker.Word(),
		},
	}

	templates, err := suite.client.Bots.WhatsApp.GetTemplates(context.Background())
	suite.NoError(err)
	suite.Equal("456", templates[0].BotID)
}
