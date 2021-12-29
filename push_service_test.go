package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (suite *SendpulseTestSuite) TestPushService_List() {
	suite.mux.HandleFunc("/push/tasks/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  {
			"id": 121,
			"title": "push title",
			"body": "push text",
			"website_id": 53,
			"from": "2015-11-17 14:44:47",
			"to": "2015-12-23 19:42:27",
			"status": 13
		  }
		]`)
	})

	list, err := suite.client.Push.GetMessages(context.Background(), PushListParams{
		Limit:     10,
		Offset:    0,
		From:      time.Now(),
		To:        time.Now(),
		WebsiteID: 10,
	})
	suite.NoError(err)
	suite.Equal(1, len(list))
}

func (suite *SendpulseTestSuite) TestPushService_WebsitesTotal() {
	suite.mux.HandleFunc("/push/websites/total", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "total": 2
		}`)
	})

	total, err := suite.client.Push.CountWebsites(context.Background())
	suite.NoError(err)
	suite.Equal(2, total)
}

func (suite *SendpulseTestSuite) TestPushService_WebsitesList() {
	suite.mux.HandleFunc("/push/websites/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  {
			"id": 53,
			"url": "www.test-site.com",
			"add_date": "2015-11-23 14:42:37",
			"status": 1
		  }
		]`)
	})

	list, err := suite.client.Push.GetWebsites(context.Background(), 10, 0)
	suite.NoError(err)
	suite.Equal(53, list[0].ID)
}

func (suite *SendpulseTestSuite) TestPushService_WebsiteVariables() {
	websiteID := 23
	suite.mux.HandleFunc(fmt.Sprintf("/push/websites/%d/variables", websiteID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		  {
			"id": 97,
			"name": "uname",
			"type": "string"
		  }
		]`)
	})

	variables, err := suite.client.Push.GetWebsiteVariables(context.Background(), websiteID)
	suite.NoError(err)
	suite.Equal("uname", variables[0].Name)
}

func (suite *SendpulseTestSuite) TestPushService_WebsiteSubscriptions() {
	websiteID := 23
	suite.mux.HandleFunc(fmt.Sprintf("/push/websites/%d/subscriptions", websiteID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"id": 311003743,
				"browser": "Firefox",
				"lang": "en",
				"os": "Linux",
				"country_code": "UA",
				"city": "Dnipro",
				"variables": [],
				"subscription_date": "2018-08-13 14:27:11",
				"status": 1
			},
			{
				"id": 311008277,
				"browser": "Opera",
				"lang": "en",
				"os": "Linux",
				"country_code": "UA",
				"city": "Dnipro",
				"variables": [],
				"subscription_date": "2018-08-13 14:33:51",
				"status": 1
			}
		]`)
	})

	subscriptions, err := suite.client.Push.GetWebsiteSubscriptions(context.Background(), websiteID, WebsiteSubscriptionsParams{
		Limit:  10,
		Offset: 0,
		From:   time.Now(),
		To:     time.Now(),
	})
	suite.NoError(err)
	suite.Equal(2, len(subscriptions))
}

func (suite *SendpulseTestSuite) TestPushService_SubscriptionsTotal() {
	websiteID := 23
	suite.mux.HandleFunc(fmt.Sprintf("/push/websites/%d/subscriptions/total", websiteID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
		  "total": 2
		}`)
	})

	total, err := suite.client.Push.CountWebsiteSubscriptions(context.Background(), websiteID)
	suite.NoError(err)
	suite.Equal(2, total)
}

func (suite *SendpulseTestSuite) TestPushService_WebsiteInfo() {
	websiteID := 23
	suite.mux.HandleFunc(fmt.Sprintf("/push/websites/info/%d", websiteID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"id": 111111,
			"url": "yoursite.com",
			"status": "active",
			"icon": "https://login.sendpulse.com/img/my/push/push-default-icons/icon.png",
			"add_date": "2017-11-09 13:08:37",
			"total_subscribers": 1081,
			"unsubscribed": 30,
			"subscribers_today": 10,
			"active_subscribers": 1051
		}`)
	})

	info, err := suite.client.Push.GetWebsiteInfo(context.Background(), websiteID)
	suite.NoError(err)
	suite.Equal("https://login.sendpulse.com/img/my/push/push-default-icons/icon.png", info.Icon)
}

func (suite *SendpulseTestSuite) TestPushService_ActivateSubscription() {
	subscrID := 12345
	suite.mux.HandleFunc("/push/subscriptions/state", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		b, _ := io.ReadAll(r.Body)
		suite.Equal(fmt.Sprintf("{\"id\":%d,\"state\":1}\n", subscrID), string(b))

		fmt.Fprintf(w, `{
		  "result": true
		}`)
	})

	err := suite.client.Push.ActivateSubscription(context.Background(), subscrID)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestPushService_DeactivateSubscription() {
	subscrID := 12345
	suite.mux.HandleFunc("/push/subscriptions/state", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		b, _ := io.ReadAll(r.Body)
		suite.Equal(fmt.Sprintf("{\"id\":%d,\"state\":0}\n", subscrID), string(b))

		fmt.Fprintf(w, `{
		  "result": true
		}`)
	})

	err := suite.client.Push.DeactivateSubscription(context.Background(), subscrID)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestPushService_CreatePushTask() {
	suite.mux.HandleFunc("/push/tasks", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)

		fmt.Fprintf(w, `{
		  "result": true,
		  "id": 1
		}`)
	})

	taskID, err := suite.client.Push.CreatePushCampaign(context.Background(), PushMessageParams{
		Title:                "Title",
		WebsiteID:            10,
		Body:                 "Hello",
		TtlSec:               10,
		Link:                 "sendpulse.com",
		FilterLang:           "ru",
		FilterBrowser:        "opera",
		FilterRegion:         "europe",
		FilterUrl:            "sendpulse.com",
		SubscriptionDateFrom: time.Now(),
		SubscriptionDateTo:   time.Now(),
		Filter:               nil,
		StretchTimeSec:       10,
		SendDate:             DateTimeType(time.Now()),
		Buttons:              nil,
		Image:                nil,
		Icon:                 nil,
	})
	suite.NoError(err)
	suite.Equal(1, taskID)
}

func (suite *SendpulseTestSuite) TestPushService_PushTaskStatistics() {
	taskID := 234
	suite.mux.HandleFunc(fmt.Sprintf("/push/tasks/%d", taskID), func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)

		fmt.Fprintf(w, `{
			"id": 36,
			"message": {
				"title": "s",
				"text": "s",
				"link": "http://aaa.aaa" 
			},
			"website": "www.google.com",
			"website_id": 53,
			"status": 3,
			"send": "21",
			"delivered": 14,
			"redirect": 13 
		}`)
	})

	stat, err := suite.client.Push.GetPushMessagesStatistics(context.Background(), taskID)
	suite.NoError(err)
	suite.Equal(36, stat.ID)
}
