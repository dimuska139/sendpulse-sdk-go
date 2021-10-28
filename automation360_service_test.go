package sendpulse_sdk_go

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"net/http"
)

func (suite *SendpulseTestSuite) TestAutomation360Service_GetAutoresponderStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/autoresponders/%d", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"autoresponder": {
				"id": 281013,
				"name": "Новая авторассылка 2020-06-23",
				"status": 1,
				"created": "2020-06-23 11:06:25",
				"changed": "2020-06-23 11:06:25"
			},
			"flows": [
				{
					"id": 467510,
					"main_id": 281013,
					"af_type": "maintrigger",
					"created": "2020-06-23 14:06:25",
					"last_send": "2020-06-25 17:05:40",
					"task": null
				},
				{
					"id": 467511, 
					"main_id": 281013,
					"af_type": "email",
					"created": "2020-06-23 14:06:25",
					"last_send": "2020-06-25 17:05:41",
					"task": {
						"id": 12018943,
						"address_book_id": 0,
						"message_title": "Спасибо за заказ, 7136062",
						"sender_mail_address": "smtp_test@e.cn.ua",
						"sender_mail_name": "smtp_test",
						"created": "2020-06-23 14:06:25"
					}
				},
				{
					"id": 486649,
					"main_id": 281013,
					"af_type": "push",
					"created": "2020-07-09 15:40:12",
					"last_send": null,
					"task": {
						"id": 3016415,
						"website_id": 14699,
						"title": "заголовок для пуш", 
						"body": "текст пуш уведомления",
						"icon": "/files/push/7023639/tasks/3016415/icons/577416720c24.png",
						"url_for_click": "https://google.com",
						"created": "2020-07-09 15:40:12"
					}
				},
				{
					"id": 486657,
					"main_id": 281013,
					"af_type": "sms",
					"created": "2020-07-09 15:48:22",
					"last_send": null,
					"task": {
						"id": 12110431,
						"address_book_id": 0,
						"sms_body": "текст смс уведомления",
						"sms_sender_name": "Bakler",
						"created": "2020-07-09 15:48:22"
					}
				}
			],
			"starts": 18,
			"in_queue": 0,
			"end_count": 18,
			"send_messages": 16,
			"conversions": 0
		}`)
	})

	statistics, err := suite.client.Automation360.GetAutoresponderStatistics(id)
	suite.NoError(err)
	suite.Equal(4, len(statistics.Flows))
}

func (suite *SendpulseTestSuite) TestAutomation360Service_StartEvent() {
	eventName := "eventname"

	suite.mux.HandleFunc(fmt.Sprintf("/events/name/%s", eventName), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.Automation360.StartEvent(eventName, map[string]interface{}{
		"email": faker.Email(),
		"name":  faker.Name(),
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetStartBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/main-trigger/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"data":{
			  "flow_id": 123,
			  "executed": 456,
			  "deleted": 9
		   }
		}`)
	})

	statistics, err := suite.client.Automation360.GetStartBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(456, statistics.Executed)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetEmailBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/email/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"data": {
				"flow_id": 467511,
				"task": {
					"id": 12018943,
					"address_book_id": 0,
					"message_title": "Спасибо за заказ, 7136062",
					"sender_mail_address": "smtp_test@e.cn.ua",
					"sender_mail_name": "smtp_test",
					"created": "2020-06-23 14:06:25"
				},
				"sent": 16,
				"delivered": 16,
				"opened": 16,
				"clicked": 0,
				"errors": 0,
				"unsubscribed": 0,
				"marked_as_spam": 0,
				"last_send": "2020-06-25 17:05:41"
			}
		}`)
	})

	statistics, err := suite.client.Automation360.GetEmailBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(467511, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetPushBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/push/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 467511,
			"sent": 9,
			"delivered": 3,
			"clicked": 6,
			"errors": 0,
			"last_send": "2020-02-28 12:18:23"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetPushBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(467511, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetSmsBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/sms/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 467511,
			"executed": 2,
			"sent": 2,
			"delivered": 2,
			"opened": null,
			"clicked": null,
			"errors": 0,
			"last_send": "2020-04-09 11:17:44"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetSmsBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(467511, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetMessengerBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/messenger/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 234728,
			"executed": 7,
			"sent": 7,
			"last_send": "2021-02-02 14:41:51"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetMessengerBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(234728, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetFilterBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/filter/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 234728,
			"executed": 7,
			"last_send": "2021-02-02 14:41:51"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetFilterBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(234728, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetTriggerBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/trigger/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 234728,
			"executed": 7,
			"last_send": "2021-02-02 14:41:51"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetTriggerBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(234728, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetGoalBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/goal/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 234728,
			"executed": 2,
			"sent": 2,
			"delivered": 2,
			"opened": null,
			"clicked": null,
			"errors": 0,
			"last_send": "2020-04-09 11:17:44"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetGoalBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(234728, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetActionBlockStatistics() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/stats/action/%d/group-stat", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "data": {
			"flow_id": 234728,
			"executed": 2,
			"last_send": "2020-04-09 11:17:44"
		  }
		}`)
	})

	statistics, err := suite.client.Automation360.GetActionBlockStatistics(id)
	suite.NoError(err)
	suite.Equal(234728, statistics.FlowID)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetAutoresponderConversions() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/autoresponders/%d/conversions", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"data": {
				"total_conversions": 5,
				"maintrigger_conversions": 1,
				"goal_conversions": 4,
				"maintrigger": {
					"id": 233500,
					"main_id": 127820,
					"af_type": "maintrigger",
					"created": "2020-04-28 18:00:09",
					"last_send": "2020-04-28 18:13:44",
					"conversions": 1
				},
				"goals": [
					{
						"id": 233502,
						"name": "left",
						"main_id": 127820,
						"af_type": "goal",
						"created": "2020-04-28 18:00:09",
						"conversions": 3
					},
					{
						"id": 233503,
						"name": "right",
						"main_id": 127820,
						"af_type": "goal",
						"created": "2020-04-28 18:00:09",
						"conversions": 1
					}
				]
			}
		}`)
	})

	conversions, err := suite.client.Automation360.GetAutoresponderConversions(id)
	suite.NoError(err)
	suite.Equal(5, conversions.TotalConversions)
}

func (suite *SendpulseTestSuite) TestAutomation360Service_GetAutoresponderContacts() {
	id := 12345

	suite.mux.HandleFunc(fmt.Sprintf("/a360/autoresponders/%d/conversions/list/all", id), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"total": 5,
			"items": [
				{
					"id": 40941,
					"conversion_type": "maintrigger",
					"flow_id": 233500,
					"email": "m.jim@sendpulse.com",
					"phone": null,
					"conversion_date": "2020-04-28 18:13:53",
					"start_date": "2020-04-28 18:13:44"
				},
				{
					"id": 40940,
					"conversion_type": "goal",
					"flow_id": 233502,
					"email": "m.jim@sendpulse.com",
					"phone": null,
					"conversion_date": "2020-04-28 18:03:44",
					"start_date": "2020-04-28 18:03:42"
				},
				{
					"id": 40939,
					"conversion_type": "goal",
					"flow_id": 233502,
					"email": "m.jim@sendpulse.com",
					"phone": null,
					"conversion_date": "2020-04-28 18:02:52",
					"start_date": "2020-04-28 18:02:50"
				},
				{
					"id": 40938,
					"conversion_type": "goal",
					"flow_id": 233503,
					"email": "m.jim@sendpulse.com",
					"phone": null,
					"conversion_date": "2020-04-28 18:02:43",
					"start_date": "2020-04-28 18:02:40"
				},
				{
					"id": 40937,
					"conversion_type": "goal",
					"flow_id": 233502,
					"email": "m.jim@sendpulse.com",
					"phone": null,
					"conversion_date": "2020-04-28 18:01:59",
					"start_date": "2020-04-28 18:01:56"
				}
			]
		}`)
	})

	contacts, err := suite.client.Automation360.GetAutoresponderContacts(id)
	suite.NoError(err)
	suite.Equal(5, len(contacts))
}
