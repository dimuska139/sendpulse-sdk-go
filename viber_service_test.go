package sendpulse_sdk_go

import (
	"fmt"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestViberService_CreateCampaign() {
	suite.mux.HandleFunc("/viber", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"data": {
				"address_book_id": null,
				"button_link": null,
				"button_text": null,
				"image_link": null,
				"message": "Ciao! Вас вітає офіційний viber-канал бренду Yamamay та нагадує, що Ви - найчарівніша.",
				"message_live_time": "1000",
				"message_type": 3,
				"resend_sms": 0,
				"send_date": "2019-03-26 12:40:05",
				"sender_id": 4501,
				"sms_sender_name": null,
				"sms_text": null,
				"task_id": 90241,
				"task_name": "Viber campaign for the personal list on 2019-03-26 12:40"
			}
		}`)
	})

	taskID, err := suite.client.Viber.CreateCampaign(CreateViberCampaignParams{
		TaskName:        "Viber task",
		MessageType:     2,
		SenderID:        2222,
		MessageLiveTime: 1000,
		SendDate:        DateTimeType(time.Now()),
		MailingListID:   12345,
		Recipients:      []int{380931111111, 380931111112, 380931111113},
		Message:         "Ciao! Вас вітає офіційний viber-канал бренду Yamamay та нагадує, що Ви - найчарівніша.",
		Additional:      nil,
	})
	suite.NoError(err)
	suite.Equal(90241, taskID)
}

func (suite *SendpulseTestSuite) TestViberService_UpdateCampaign() {
	suite.mux.HandleFunc("/viber/update", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
			"result": true,
			"data": {
				"address_book_id": null,
				"button_link": null,
				"button_text": null,
				"image_link": null,
				"message": "Ciao! Ciao Ciao Ciao Ciao Вас вітає офіційний viber-канал бренду Yamamay та нагадує, що Ви - найчарівніша.",
				"message_live_time": "1000",
				"message_type": "3",
				"send_date": "2019-03-26 15:16:00",
				"sender_id": "4495",
				"task_id": 9380939,
				"task_name": "Viber campaign for the personal list on 2019-03-26 15:15"
			}
		}`)
	})

	err := suite.client.Viber.UpdateCampaign(UpdateViberCampaignParams{
		TaskID:          12345,
		TaskName:        "Task name",
		Message:         "New viber message",
		MessageType:     2,
		ButtonText:      "",
		ButtonLink:      "",
		ImageLink:       "",
		AddressBookID:   12345,
		SenderID:        222,
		MessageLiveTime: 1000,
		SendDate:        DateTimeType(time.Now()),
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestViberService_Campaigns() {
	suite.mux.HandleFunc("/viber/task", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `[
			{
				"id": 9380939,
				"name": "Viber campaign for the personal list on 2019-03-26 15:02",
				"message": "Ciao! Ciao Ciao Ciao Ciao Вас вітає офіційний viber-канал бренду Yamamay та нагадує, що Ви - найчарівніша.",
				"button_text": null,
				"button_link": null,
				"image_link": null,
				"address_book": null,
				"sender_name": "YAMAMAY",
				"sender_id": 4495,
				"message_live_time": 1000,
				"send_date": "2019-03-29 10:00:00",
				"status": "moderation",
				"created": "2019-03-26 12:50:02"
			},
			{
				"id": 9380926,
				"name": "Viber campaign for the personal list on 2019-03-26 14:48",
				"message": "Ciao! Вас вітає офіційний viber-канал бренду Yamamay та нагадує, що Ви - найчарівніша.",
				"button_text": null,
				"button_link": null,
				"image_link": null,
				"address_book": 0,
				"sender_name": "YAMAMAY",
				"sender_id": 4495,
				"message_live_time": 1000,
				"send_date": "2019-03-29 10:00:00",
				"status": null,
				"created": "2019-03-26 12:48:23"
			}
		]`)
	})

	campaigns, err := suite.client.Viber.GetCampaigns(10, 0)
	suite.NoError(err)
	suite.Equal(2, len(campaigns))
}

func (suite *SendpulseTestSuite) TestViberService_GetStatistics() {
	campaignID := 1223
	suite.mux.HandleFunc(fmt.Sprintf("/viber/task/%d", campaignID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
				"id": 38,
				"name": "Viber_Campaign_38",
				"message": "Это текст для вайбер сообщения",
				"button_text": "Кнопка",
				"button_link": "https://sendpulse.com",
				"image_link": null,
				"address_book": null,
				"sender_name": "infoservice",
				"send_date": "2017-06-22 09:51:35",
				"status": "send",
				"statistic": {
					"sent": 1,
					"delivered": 1,
					"read": 0,
					"redirected": 0,
					"undelivered": 0,
					"errors": 0
				},
					"created": "2017-06-22 09:51:22" 
		}`)
	})

	statistics, err := suite.client.Viber.GetStatistics(campaignID)
	suite.NoError(err)
	suite.Equal(38, statistics.ID)
}

func (suite *SendpulseTestSuite) TestViberService_GetSenders() {
	suite.mux.HandleFunc("/viber/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `[
			{
				"id": 2222,
				"status": "verified",
				"name": "Foxkids",
				"service_type": "Магазин iграшок",
				"web_site": "www.foxkids.com",
				"description": "Магазин Foxkids –«Королівство іграшок» де знайдете багато речей, необхідних для комфорту і розвитку вашого малюка",
				"countries": [
					"UA"
				],
				"traffic_type": "Рекламные сообщения",
				"admin_comment": null,
				"owner": "you"
			}
		]`)
	})

	senders, err := suite.client.Viber.GetSenders()
	suite.NoError(err)
	suite.Equal(2222, senders[0].ID)
}

func (suite *SendpulseTestSuite) TestViberService_GetSender() {
	senderID := 12345
	suite.mux.HandleFunc(fmt.Sprintf("/viber/senders/%d", senderID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"id": 1,
			"status": "verified",
			"name": "infoservice",
			"service_type": "Тестовый сервис",
			"web_site": "https://www.sendpulse.com",
			"description": "Мы тестируем финальную отправку сообщений",
			"country": "UA",
			"traffic_type": "Публичная информация",
			"admin_comment": "Ваше имя одобрено, спасибо что выбрали наш сервис для отправки вайбер сообщений. Команда Sendpulse" 
		}`)
	})

	sender, err := suite.client.Viber.GetSender(senderID)
	suite.NoError(err)
	suite.Equal("infoservice", sender.Name)
}

func (suite *SendpulseTestSuite) TestViberService_GetRecipients() {
	taskID := 12345
	suite.mux.HandleFunc(fmt.Sprintf("/viber/task/%d/recipients", taskID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"task_id": 44,
			"recipients": [
				{
					"phone": 380934760182,
					"address_book_id": 850852,
					"status": "send",
					"send_date": "2017-06-23 08:54:01",
					"price": 0.74,
					"currency": "RUR",
					"last_update": "2017-06-23 08:53:38" 
				}
			]
		}`)
	})

	recipients, err := suite.client.Viber.GetRecipients(taskID)
	suite.NoError(err)
	suite.Equal("RUR", recipients[0].Currency)
}
