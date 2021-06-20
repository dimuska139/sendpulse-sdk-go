package sendpulse

import (
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_TemplatesService_Create() {
	suite.mux.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"result": true,"real_id":1}`)
	})

	tplID, err := suite.client.Emails.Templates.Create("First template", "<h1>Message</h1>", "ru")
	suite.NoError(err)
	suite.Equal(1, tplID)
}

func (suite *SendpulseTestSuite) TestEmailsService_TemplatesService_Update() {
	suite.mux.HandleFunc("/template/edit/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.Templates.Update(1, "<h1>Message</h1>", "ru")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_TemplatesService_Get() {
	suite.mux.HandleFunc("/template/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"id": "1",
			"real_id": 1,
			"name": "Тестовый шаблон 1",
			"name_slug": "testovyy-shablon-1",
			"meta_description": null,
			"full_description": null,
			"category": "",
			"category_info": [],
			"mark": null,
			"mark_count": null,
			"body": "PGgxPtCf0YDQrtCy0LXRgiE8L2gxPg==",
			"tags": {
				"digest": "digest",
				"blog": "blog",
				"exhibition": "exhibition",
				"invite": "invite"
			},
			"created": "2021-06-19 19:18:32",
			"preview": "https://login.sendpulse.com/files/emailservice/userfiles/templates/preview/123/12345.png",
			"owner": "me",
			"is_structure": false
		}`)
	})

	tpl, err := suite.client.Emails.Templates.Get(1)
	suite.NoError(err)
	suite.Equal(1, tpl.RealID)
}

func (suite *SendpulseTestSuite) TestEmailsService_TemplatesService_List() {
	suite.mux.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"id": "f3266876955c9d21e214deed49b97446",
				"real_id": 1153018,
				"lang": "en",
				"name": "Webinar Speakers",
				"name_slug": "",
				"created": "2020-09-04 13:54:30",
				"full_description": "Use this template as a webinar invitation for your subscribers. Specify who is going to host the webinar and what it will be about. Remember to include the date and time of the webinar.",
				"is_structure": true,
				"category": "education",
				"category_info": {
					"id": 109,
					"name": "Education",
					"meta_description": "These “Education” free email templates were developed by SendPulse for all those who wish to make their email communication colorful and unforgettable. You can use these templates to create your email campaigns in SendPulse.",
					"full_description": "",
					"code": "education",
					"sort": 6
				},
				"tags": {
					"webinar": "webinar",
					"study": "study",
					"marketing": "marketing",
					"museum": "museum",
					"exhibition": "exhibition"
				},
				"owner": "sendpulse",
				"preview": "https://login.sendpulse.com/files/emailservice/userfiles/templates/preview/f3266876955c9d21e214deed49b97446_thumbnail_300.png"
			},
			{
				"id": "2a7c59e5bcb0db1dee02c60208fbb498",
				"real_id": 1153017,
				"lang": "en",
				"name": "Offline conference",
				"name_slug": "education-events",
				"created": "2020-09-04 13:54:21",
				"full_description": "Use this template to invite your subscribers to an online or offline event. Indicate its name, date, and time. Add information about the  location if it's an offline event. Upload photos of your speakers and give a short description of their presentation. That’s it! Now you can send your email campaign!",
				"is_structure": true,
				"category": "education",
				"category_info": [],
				"tags": {
					"digest": "digest",
					"blog": "blog",
					"exhibition": "exhibition",
					"invite": "invite"
				},
				"owner": "sendpulse",
				"preview": "https://login.sendpulse.com/files/emailservice/userfiles/templates/preview/2a7c59e5bcb0db1dee02c60208fbb498_thumbnail_300.png"
			}]`)
	})
	templates, err := suite.client.Emails.Templates.List(10, 0, "me")
	suite.NoError(err)
	suite.Equal(2, len(templates))
}
