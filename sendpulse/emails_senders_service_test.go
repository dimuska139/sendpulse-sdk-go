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

	err := suite.client.Emails.Senders.Create("Ivan Petrov", "test@sendpulse.com")
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

	err := suite.client.Emails.Senders.GetActivationCode("test@sendpulse.com")
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

	err := suite.client.Emails.Senders.Activate("test@sendpulse.com", "code")
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

	senders, err := suite.client.Emails.Senders.List()
	suite.NoError(err)
	suite.Equal(2, len(senders))
}

func (suite *SendpulseTestSuite) TestEmailsService_SendersService_Delete() {
	suite.mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.Senders.Delete("test@sendpulse.com")
	suite.NoError(err)
}
