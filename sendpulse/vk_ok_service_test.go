package sendpulse

import (
	"encoding/base64"
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
	"os"
	"time"
)

func (suite *SendpulseTestSuite) TestVkOkService_CreateSender() {
	suite.mux.HandleFunc("/vk-ok/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"id": 586,
			"user_id": 12345,
			"name": "Alex",
			"vk_url": "https://vk.com/vk",
			"ok_url": null,
			"created_at": "2021-07-03T13:19:10.000000Z",
			"update_at": "2021-07-03T13:19:10.000000Z"
		}`)
	})

	b64 := "/9j/4AAQSkZJRgABAQEBLAEsAAD//gATQ3JlYXRlZCB3aXRoIEdJTVD/4gKwSUNDX1BST0ZJTEUAAQEAAAKgbGNtcwQwAABtbnRyUkdCIFhZWiAH5QAHAAMADQA1ACxhY3NwQVBQTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA9tYAAQAAAADTLWxjbXMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA1kZXNjAAABIAAAAEBjcHJ0AAABYAAAADZ3dHB0AAABmAAAABRjaGFkAAABrAAAACxyWFlaAAAB2AAAABRiWFlaAAAB7AAAABRnWFlaAAACAAAAABRyVFJDAAACFAAAACBnVFJDAAACFAAAACBiVFJDAAACFAAAACBjaHJtAAACNAAAACRkbW5kAAACWAAAACRkbWRkAAACfAAAACRtbHVjAAAAAAAAAAEAAAAMZW5VUwAAACQAAAAcAEcASQBNAFAAIABiAHUAaQBsAHQALQBpAG4AIABzAFIARwBCbWx1YwAAAAAAAAABAAAADGVuVVMAAAAaAAAAHABQAHUAYgBsAGkAYwAgAEQAbwBtAGEAaQBuAABYWVogAAAAAAAA9tYAAQAAAADTLXNmMzIAAAAAAAEMQgAABd7///MlAAAHkwAA/ZD///uh///9ogAAA9wAAMBuWFlaIAAAAAAAAG+gAAA49QAAA5BYWVogAAAAAAAAJJ8AAA+EAAC2xFhZWiAAAAAAAABilwAAt4cAABjZcGFyYQAAAAAAAwAAAAJmZgAA8qcAAA1ZAAAT0AAACltjaHJtAAAAAAADAAAAAKPXAABUfAAATM0AAJmaAAAmZwAAD1xtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAAgAAAAcAEcASQBNAFBtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAAgAAAAcAHMAUgBHAEL/2wBDAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/2wBDAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQH/wgARCAAFAAUDAREAAhEBAxEB/8QAFAABAAAAAAAAAAAAAAAAAAAACf/EABQBAQAAAAAAAAAAAAAAAAAAAAD/2gAMAwEAAhADEAAAAX8P/8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABBQJ//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPwF//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAgEBPwF//8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQAGPwJ//8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABPyF//9oADAMBAAIAAwAAABAf/8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAwEBPxB//8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAgBAgEBPxB//8QAFBABAAAAAAAAAAAAAAAAAAAAAP/aAAgBAQABPxB//9k="
	dec, err := base64.StdEncoding.DecodeString(b64)

	f, err := os.Create("/tmp/img.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(dec)

	id, err := suite.client.VkOk.CreateSender(CreateVkOkSenderParams{
		Name:        "Test",
		VkUrl:       "https://vk.com/vk",
		OkUrl:       "https://ok.com/ok",
		CoverLetter: f,
	})

	suite.NoError(err)
	suite.Equal(586, id)
}

func (suite *SendpulseTestSuite) TestVkOkService_CreateTemplate() {
	suite.mux.HandleFunc("/vk-ok/templates", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"total": 1,
			"data": {
				"id": 123,
				"user_id": 6615360,
				"sender_id": 1,
				"name": "Отправка накладной",
				"vk_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"ok_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"sender": {
					"id": 1,
					"user_id": 6615360,
					"name": "Healthy Lifestyle",
					"vk_url": "https://vk.com/healthy_lifestyle",
					"ok_url": "https://ok.ru/group/56949408137424",
					"created_at": "2020-05-06T12:09:04.000000Z",
					"update_at": "2020-05-06T12:09:04.000000Z"
				},
				"status": 3,
				"status_detail": {
					"id": 1,
					"name": "new"
				}
			}
		}`)
	})

	tplID, err := suite.client.VkOk.CreateTemplate(CreateVkOkTemplateParams{
		Name:      "Test",
		VkMessage: "Test message",
		SenderID:  12345,
	})
	suite.NoError(err)
	suite.Equal(123, tplID)
}

func (suite *SendpulseTestSuite) TestVkOkService_Templates() {
	suite.mux.HandleFunc("/vk-ok/templates", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "total": 2,
		  "data": [
			{
			  "id": 1,
			  "user_id": 6615360,
			  "sender_id": 1,
			  "name": "block campaign",
			  "vk_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
			  "ok_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
			  "sender": {
				"id": 1,
				"user_id": 6615360,
				"name": "Healthy Lifestyle",
				"vk_url": "https://vk.com/healthy_lifestyle",
				"ok_url": "https://ok.ru/group/56949408137424",
				"created_at": "2020-05-06T12:09:04.000000Z",
				"update_at": "2020-05-06T12:09:04.000000Z"
			  },
			  "status": 3,
			  "status_detail": {
				"id": 3,
				"name": "moderation_by_provider"
			  }
			},
			{
			  "id": 15,
			  "user_id": 6615360,
			  "sender_id": 2,
			  "name": "test",
			  "vk_message": "Привет! Через 15 минут мы начинаем вебинар по правильному питанию! Ждем вас в вебинарной комнате, старт встречи в 15:00",
			  "ok_message": "Привет! Через 15 минут мы начинаем вебинар по правильному питанию! Ждем вас в вебинарной комнате, старт встречи в 15:00",
			  "sender": {
				"id": 2,
				"user_id": 6615360,
				"name": "Healthy Lifestyle",
				"vk_url": "https://vk.com/healthy_lifestyle",
				"ok_url": "https://vk.com/sendpulse",
				"created_at": "2020-05-20T12:59:01.000000Z",
				"update_at": "2020-05-20T12:59:01.000000Z"
			  },
			  "status": 1,
			  "status_detail": {
				"id": 1,
				"name": "new"
			  }
			}
		  ]
		}`)
	})

	templates, err := suite.client.VkOk.Templates()
	suite.NoError(err)
	suite.Equal(6615360, templates[0].UserID)
}

func (suite *SendpulseTestSuite) TestVkOkService_Template() {
	tplID := 12345
	suite.mux.HandleFunc(fmt.Sprintf("/vk-ok/templates/%d", tplID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"total": 1,
			"data": {
				"id": 1,
				"user_id": 6615360,
				"sender_id": 1,
				"name": "block campaign",
				"vk_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"ok_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"sender": {
					"id": 1,
					"user_id": 6615360,
					"name": "Healthy Lifestyle",
					"vk_url": "https://vk.com/healthy_lifestyle",
					"ok_url": "https://ok.ru/group/56949408137424",
					"created_at": "2020-05-06T12:09:04.000000Z",
					"update_at": "2020-05-06T12:09:04.000000Z"
				},
				"status": 3,
				"status_detail": {
					"id": 3,
					"name": "moderation_by_provider"
				}
			}
		}`)
	})

	template, err := suite.client.VkOk.Template(tplID)
	suite.NoError(err)
	suite.Equal(6615360, template.UserID)
}

func (suite *SendpulseTestSuite) TestVkOkService_Send() {
	suite.mux.HandleFunc("/vk-ok/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"total": 1,
			"data": {
				"id": 12345
			}
		}`)
	})

	campaignID, err := suite.client.VkOk.Send(SendVkOkTemplateParams{
		AddressBooks: []int{123, 456},
		Recipients: []struct {
			Phone     string                 `json:"phone"`
			Variables map[string]interface{} `json:"variables"`
		}{
			{
				Phone:     "89221123344",
				Variables: map[string]interface{}{"name": "Alex"},
			},
		},
		LifeTime:   1000,
		LifeType:   "min",
		Name:       "Test",
		Routes:     nil,
		SendDate:   models.DateTimeType(time.Now()),
		TemplateID: 12345,
	})
	suite.NoError(err)
	suite.Equal(12345, campaignID)
}

func (suite *SendpulseTestSuite) TestVkOkService_CampaignsStatistics() {
	suite.mux.HandleFunc("/vk-ok/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
		  "total": 1,
		  "data": [
			{
			  "id": 1,
			  "user_id": 6615360,
			  "name": "test",
			  "total_price": 0,
			  "price_rate": 100000,
			  "currency": {
				"id": 3,
				"currency_name": "Ukrainian hryvnias",
				"currency_abbr": "UAH",
				"currency_sign": "грн."
			  },
			  "life_time": 24,
			  "life_type": "hour",
			  "send_date": "2020-05-06T15:42:12.000000Z",
			  "created_at": "2020-05-06T15:42:12.000000Z",
			  "template": {
				"id": 1,
				"user_id": 6615360,
				"sender_id": 1,
				"name": "block campaign",
				"vk_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"ok_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"sender": {
				  "id": 1,
				  "user_id": 6615360,
				  "name": "sendpulse",
				  "vk_url": "https://vk.com/healthy_lifestyle",
				  "ok_url": "https://ok.ru/group/56949408137424",
				  "created_at": "2020-05-06T12:09:04.000000Z",
				  "update_at": "2020-05-06T12:09:04.000000Z"
				},
				"status": 3
			  },
			  "status": 7,
			  "status_detail": {
				"id": 7,
				"name": "in_progress"
			  },
			  "group_stat": []
			}
		  ]
		}`)
	})

	statistics, err := suite.client.VkOk.CampaignsStatistics()
	suite.NoError(err)
	suite.Equal(6615360, statistics[0].UserID)
}

func (suite *SendpulseTestSuite) TestVkOkService_CampaignStatistics() {
	campaignID := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/vk-ok/campaigns/%d", campaignID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"id": 1,
			"user_id": 6615360,
			"name": "test",
			"address_books": [
				{
					"id": 1,
					"user_id": 6615360,
					"campaign_id": 1,
					"address_book_id": 88890330,
					"created_at": "2020-05-06 12:42:12",
					"updated_at": "2020-05-06 12:42:12"
				}
			],
			"total_price": 0,
			"price_rate": 100000,
			"currency": {
				"id": 3,
				"currency_name": "Ukrainian hryvnias",
				"currency_abbr": "UAH",
				"currency_sign": "грн."
			},
			"life_time": 24,
			"life_type": "hour",
			"send_date": "2020-05-06T15:42:12.000000Z",
			"created_at": "2020-05-06T15:42:12.000000Z",
			"template": {
				"id": 1,
				"user_id": 6615360,
				"sender_id": 1,
				"name": "block campaign",
				"vk_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"ok_message": "Ваш заказ отправлен. TTN {{ttn_number}}. Срок хранения 5 рабочих дней.",
				"sender": {
					"id": 1,
					"user_id": 6615360,
					"name": "sendpulse",
					"vk_url": "https://vk.com/healthy_lifestyle",
					"ok_url": "https://ok.ru/group/56949408137424",
					"created_at": "2020-05-06T12:09:04.000000Z",
					"update_at": "2020-05-06T12:09:04.000000Z"
				},
				"status": 3
			},
			"status": 7,
			"status_detail": {
				"id": 7,
				"name": "in_progress"
			},
			"group_stat": [
					{
						"id": 5,
						"user_id": 6615360,
						"route": null,
						"sent": 1,
						"delivered": 0,
						"not_delivered": 1,
						"opened": 0
					}
			]
		}`)
	})

	statistics, err := suite.client.VkOk.CampaignStatistics(campaignID)
	suite.NoError(err)
	suite.Equal(6615360, statistics.UserID)
}

func (suite *SendpulseTestSuite) TestVkOkService_CampaignPhones() {
	campaignID := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/vk-ok/campaigns/%d/phones", campaignID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"total": 2,
			"data": [
				{
					"id": 1,
					"user_id": 6615360,
					"campaign_id": 1,
					"template_id": null,
					"phone": 380931112233,
					"phone_cost": 41000,
					"currency_id": 3,
					"price_rate": 100000,
					"currency": {
						"id": 3,
						"currency_name": "Ukrainian hryvnias",
						"currency_abbr": "UAH",
						"currency_sign": "грн."
					},
					"created_at": "2020-05-06T15:45:58.000000Z",
					"status": 1,
					"status_detail": {
						"id": 1,
						"name": "to_send"
					}
				},
				{
					"id": 2,
					"user_id": 6615360,
					"campaign_id": 1,
					"template_id": null,
					"phone": 380931112234,
					"phone_cost": 41000,
					"currency_id": 3,
					"price_rate": 100000,
					"currency": {
						"id": 3,
						"currency_name": "Ukrainian hryvnias",
						"currency_abbr": "UAH",
						"currency_sign": "грн."
					},
					"created_at": "2020-05-06T15:45:58.000000Z",
					"status": 1,
					"status_detail": {
						"id": 1,
						"name": "to_send"
					}
				}
			]
		}`)
	})

	phones, err := suite.client.VkOk.CampaignPhones(campaignID)
	suite.NoError(err)
	suite.Equal(2, len(phones))
}
