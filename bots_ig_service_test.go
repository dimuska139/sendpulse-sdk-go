package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestBotsIgService_GetAccount() {
	suite.mux.HandleFunc("/instagram/account", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"tariff": {
			  "code": "code",
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

	account, err := suite.client.Bots.Ig.GetAccount(context.Background())
	suite.NoError(err)
	suite.True(account.Tariff.Branding)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetBots() {
	suite.mux.HandleFunc("/instagram/bots", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [
			{
			  "id": "12345",
			  "channel_data": {
				"fb_user": {
				  "id": "string",
				  "first_name": "string",
				  "last_name": "string",
				  "name": "string",
				  "name_format": "string",
				  "short_name": "string",
				  "picture": {
					"data": {
					  "height": 0,
					  "is_silhouette": true,
					  "url": "string",
					  "width": 0
					}
				  }
				},
				"ig_user": {
				  "id": 0,
				  "ig_id": 0,
				  "followers_count": 0,
				  "follows_count": 0,
				  "media_count": 0,
				  "profile_picture_url": "string",
				  "username": "string",
				  "website": "string"
				},
				"ig_page": {
				  "instagram_business_account": {
					"id": 0,
					"ig_id": 0,
					"name": "string",
					"biography": "string",
					"followers_count": 0,
					"follows_count": 0,
					"media_count": 0,
					"profile_picture_url": "string",
					"website": "string",
					"username": "string"
				  },
				  "id": 0,
				  "category": "string",
				  "category_list": [
					{
					  "id": 0,
					  "name": "string"
					}
				  ],
				  "name": "string",
				  "picture": {
					"data": {
					  "height": 0,
					  "is_silhouette": true,
					  "url": "string",
					  "width": 0
					}
				  }
				}
			  },
			  "inbox": {
				"total": 0,
				"unread": 0
			  },
			  "status": 3,
			  "created_at": "2020-12-12T00:00:00+03:00"
			}
		  ]
		}`)
	})

	bots, err := suite.client.Bots.Ig.GetBots(context.Background())
	suite.NoError(err)
	suite.Equal("12345", bots[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetContact() {
	suite.mux.HandleFunc("/instagram/contacts/get", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": {
			"id": "12345",
			"bot_id": "string",
			"status": 1,
			"channel_data": {
			  "id": "2344",
			  "user_name": "string",
			  "first_name": "string",
			  "last_name": "string",
			  "name": "string",
			  "profile_pic": "string"
			},
			"tags": [
			  "string"
			],
			"variables": {},
			"is_chat_opened": true,
			"last_activity_at": "2020-12-12T00:00:00+03:00",
			"automation_paused_until": "2020-12-12T00:00:00+03:00",
			"created_at": "2020-12-12T00:00:00+03:00"
		  }
		}`)
	})

	contact, err := suite.client.Bots.Ig.GetContact(context.Background(), "1")
	suite.NoError(err)
	suite.Equal("12345", contact.ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetContactsByTag() {
	suite.mux.HandleFunc("/instagram/contacts/getByTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [
			{
			  "id": "12345",
			  "bot_id": "string",
			  "status": 1,
			  "channel_data": {
				"id": "54321",
				"user_name": "string",
				"first_name": "string",
				"last_name": "string",
				"name": "string",
				"profile_pic": "string"
			  },
			  "tags": [
				"string"
			  ],
			  "variables": {},
			  "is_chat_opened": true,
			  "last_activity_at": "2020-12-12T00:00:00+03:00",
			  "automation_paused_until": "2020-12-12T00:00:00+03:00",
			  "created_at": "2020-12-12T00:00:00+03:00"
			}
		  ]
		}`)
	})

	contacts, err := suite.client.Bots.Ig.GetContactsByTag(context.Background(), "tag", "bot_id")
	suite.NoError(err)
	suite.Equal("12345", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetContactsByVariable() {
	suite.mux.HandleFunc("/instagram/contacts/getByVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [
			{
			  "id": "12345",
			  "bot_id": "string",
			  "status": 1,
			  "channel_data": {
				"id": "54321",
				"user_name": "string",
				"first_name": "string",
				"last_name": "string",
				"name": "string",
				"profile_pic": "string"
			  },
			  "tags": [
				"string"
			  ],
			  "variables": {},
			  "is_chat_opened": true,
			  "last_activity_at": "2020-12-12T00:00:00+03:00",
			  "automation_paused_until": "2020-12-12T00:00:00+03:00",
			  "created_at": "2020-12-12T00:00:00+03:00"
			}
		  ]
		}`)
	})

	contacts, err := suite.client.Bots.Ig.GetContactsByVariable(context.Background(), BotContactsByVariableParams{
		VariableID:    "var_id",
		VariableName:  "name",
		BotID:         "qwe123",
		VariableValue: "alex",
	})
	suite.NoError(err)
	suite.Equal("12345", contacts[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_SendTextByContact() {
	suite.mux.HandleFunc("/instagram/contacts/sendText", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.SendTextByContact(context.Background(), IgBotSendMessagesParams{
		ContactID: "qwe12345",
		Messages: []struct {
			Type    string `json:"type"`
			Message struct {
				Text string `json:"text"`
			} `json:"message"`
		}{{
			Type: "text",
			Message: struct {
				Text string `json:"text"`
			}{
				Text: "text",
			},
		}},
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_SetVariableToContact() {
	suite.mux.HandleFunc("/instagram/contacts/setVariable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.SetVariableToContact(context.Background(), "contactId", "variableId", "variableName", 123)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_SetTagsToContact() {
	suite.mux.HandleFunc("/instagram/contacts/setTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.SetTagsToContact(context.Background(), "contactId", []string{"tag1", "tag2"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_DeleteTagFromContact() {
	suite.mux.HandleFunc("/instagram/contacts/deleteTag", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.DeleteTagFromContact(context.Background(), "contactId", "tag1")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_DisableContact() {
	suite.mux.HandleFunc("/instagram/contacts/disable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.DisableContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_EnableContact() {
	suite.mux.HandleFunc("/instagram/contacts/enable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.EnableContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_DeleteContact() {
	suite.mux.HandleFunc("/instagram/contacts/delete", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.DeleteContact(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetPauseAutomation() {
	suite.mux.HandleFunc("/instagram/contacts/getPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
			"data": {
				"minutes": 123
			}
		}`)
	})

	p, err := suite.client.Bots.Ig.GetPauseAutomation(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal(123, p)
}

func (suite *SendpulseTestSuite) TestBotsIgService_SetPauseAutomation() {
	suite.mux.HandleFunc("/instagram/contacts/setPauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.SetPauseAutomation(context.Background(), "contactId", 60)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_DeletePauseAutomation() {
	suite.mux.HandleFunc("/instagram/contacts/deletePauseAutomation", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.DeletePauseAutomation(context.Background(), "contactId")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetBotVariables() {
	suite.mux.HandleFunc("/instagram/variables", func(w http.ResponseWriter, r *http.Request) {
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

	variables, err := suite.client.Bots.Ig.GetBotVariables(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal("qwerty", variables[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetFlows() {
	suite.mux.HandleFunc("/instagram/flows", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "success": true,
		  "data": [{
			"id": "qwe123",
			"bot_id": "qwe12345",
			"name": "Alex",
			"status": {
			  "ACTIVE": 1,
			  "INACTIVE": 2,
			  "DRAFT": 4
			},
			"triggers": [
			  {
				"id": "string",
				"name": "string",
				"type": 1
			  }
			],
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	flows, err := suite.client.Bots.Ig.GetFlows(context.Background(), "contactId")
	suite.NoError(err)
	suite.Equal("qwe123", flows[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_RunFlow() {
	suite.mux.HandleFunc("/instagram/flows/run", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.RunFlow(context.Background(), "contactId", "flowId", map[string]any{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_RunFlowByTrigger() {
	suite.mux.HandleFunc("/instagram/flows/runByTrigger", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	err := suite.client.Bots.Ig.RunFlowByTrigger(context.Background(), "contactId", "keyword", map[string]any{
		"tracking_number": "1234-0987-5678-9012",
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetBotTriggers() {
	suite.mux.HandleFunc("/instagram/triggers", func(w http.ResponseWriter, r *http.Request) {
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

	triggers, err := suite.client.Bots.Ig.GetBotTriggers(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("qwe1234", triggers[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetBotChats() {
	suite.mux.HandleFunc("/instagram/chats", func(w http.ResponseWriter, r *http.Request) {
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

	chats, err := suite.client.Bots.Ig.GetBotChats(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("string", chats[0].InboxLastMessage.CampaignID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_GetContactMessages() {
	suite.mux.HandleFunc("/instagram/chats/messages", func(w http.ResponseWriter, r *http.Request) {
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
			"created_at": "2020-12-12T00:00:00+03:00"
		  }]
		}`)
	})

	messages, err := suite.client.Bots.Ig.GetContactMessages(context.Background(), "bot")
	suite.NoError(err)
	suite.Equal("string", messages[0].ID)
}

func (suite *SendpulseTestSuite) TestBotsIgService_SendCampaign() {
	suite.mux.HandleFunc("/instagram/campaigns/send", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "success": true
		}`)
	})

	messages := make([]IgBotCampaignMessage, 0)
	messages = append(messages, IgBotCampaignMessage{
		Type: "type",
		Message: struct {
			Text string `json:"text"`
		}{
			Text: "text",
		},
	})

	err := suite.client.Bots.Ig.SendCampaign(context.Background(), IgBotSendCampaignParams{
		Title:    "Title",
		BotID:    "qwe123",
		SendAt:   time.Now(),
		Messages: messages,
	})
	suite.NoError(err)
}
