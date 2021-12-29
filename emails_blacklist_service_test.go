package sendpulse_sdk_go

import (
	"context"
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_BlacklistService_Add() {
	suite.mux.HandleFunc("/blacklist", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.Blacklist.AddToBlacklist(context.Background(), []string{"test@sendpulse.com"}, "Added to blacklist")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_BlacklistService_Remove() {
	suite.mux.HandleFunc("/blacklist", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.Emails.Blacklist.RemoveFromBlacklist(context.Background(), []string{"test@sendpulse.com"})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_BlacklistService_List() {
	suite.mux.HandleFunc("/blacklist", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			"test@sendpulse.com",
			"test1@sendpulse.com"
		]`)
	})

	blacklist, err := suite.client.Emails.Blacklist.GetEmails(context.Background())
	suite.NoError(err)
	suite.Equal(2, len(blacklist))
}
