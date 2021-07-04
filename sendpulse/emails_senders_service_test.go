package sendpulse

import (
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_SendersService_Create() {
	suite.mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.Senders.CreateSender("Ivan Petrov", "test@sendpulse.com")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_SendersService_GetActivationCode() {
	suite.mux.HandleFunc("/senders/test@sendpulse.com/code", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"result": true,
			"email": "test@sendpulse.com"
		}`)
	})

	err := suite.client.Emails.Senders.GetSenderActivationCode("test@sendpulse.com")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_SendersService_Activate() {
	suite.mux.HandleFunc("/senders/test@sendpulse.com/code", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true,
			"email": "test@sendpulse.com"
		}`)
	})

	err := suite.client.Emails.Senders.ActivateSender("test@sendpulse.com", "code")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_SendersService_List() {
	suite.mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
		{
			"email": "test@sendpulse.com",
			"name": "Dmitriy Petrov",
			"status": "Active"
		},
		{
			"email": "test1@sendpulse.com",
			"name": "Petr Ivanov",
			"status": "Requested activation"
		}]`)
	})

	senders, err := suite.client.Emails.Senders.GetSenders()
	suite.NoError(err)
	suite.Equal(2, len(senders))
}

func (suite *SendpulseTestSuite) TestEmailsService_SendersService_Delete() {
	suite.mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.Senders.DeleteSender("test@sendpulse.com")
	suite.NoError(err)
}
