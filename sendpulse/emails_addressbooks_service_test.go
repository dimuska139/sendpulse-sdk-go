package sendpulse

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Create() {
	suite.mux.HandleFunc("/addressbooks", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"id": 12345}`)
	})

	id, err := suite.client.Emails.AddressBooks.Create("name")
	suite.NoError(err)
	suite.Equal(12345, id)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Update() {
	suite.mux.HandleFunc("/addressbooks/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPut, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.AddressBooks.Update(1, "name")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_List() {
	suite.mux.HandleFunc("/addressbooks", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"id": 1266208,
				"name": "Book 1",
				"all_email_qty": 0,
				"active_email_qty": 0,
				"inactive_email_qty": 0,
				"creationdate": "2021-06-18 19:57:39",
				"status": 0,
				"status_explain": "Active"
			},
			{
				"id": 1266209,
				"name": "Book 2",
				"all_email_qty": 0,
				"active_email_qty": 0,
				"inactive_email_qty": 0,
				"creationdate": "2021-06-19 11:02:14",
				"status": 0,
				"status_explain": "Active"
			}
		]`)
	})

	books, err := suite.client.Emails.AddressBooks.List(10, 0)
	suite.NoError(err)
	suite.Equal(2, len(books))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Get() {
	suite.mux.HandleFunc("/addressbooks/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"id": 1266208,
				"name": "Book 1",
				"all_email_qty": 0,
				"active_email_qty": 0,
				"inactive_email_qty": 0,
				"creationdate": "2021-06-18 19:57:39",
				"status": 0,
				"status_explain": "Active"
			}
		]`)
	})

	book, err := suite.client.Emails.AddressBooks.Get(1)
	suite.NoError(err)
	suite.NotNil(book)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Variables() {
	suite.mux.HandleFunc("/addressbooks/1/variables", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"name": "age",
				"type": "number"
			},
			{
				"name": "weight",
				"type": "number"
			}
		]`)
	})

	variables, err := suite.client.Emails.AddressBooks.Variables(1)
	suite.NoError(err)
	suite.Equal(2, len(variables))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Emails() {
	suite.mux.HandleFunc("/addressbooks/1/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"email": "test@test.com",
				"status": 0,
				"phone": 79312351234,
				"status_explain": "New",
				"variables": {
					"age": 12
				}
			}
		]`)
	})

	emails, err := suite.client.Emails.AddressBooks.Emails(1, 100, 0)
	suite.NoError(err)
	suite.Equal("test@test.com", emails[0].Email)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_EmailsTotal() {
	suite.mux.HandleFunc("/addressbooks/1/emails/total", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"total": 12
		}`)
	})

	total, err := suite.client.Emails.AddressBooks.EmailsTotal(1)
	suite.NoError(err)
	suite.Equal(12, total)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_EmailsByVariable() {
	suite.mux.HandleFunc("/addressbooks/1/variables/age/12", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"email": "test@test.com",
				"status": 0,
				"status_explain": "New"
			}
		]`)
	})

	emails, err := suite.client.Emails.AddressBooks.EmailsByVariable(1, "age", 12)
	suite.NoError(err)
	suite.Equal("test@test.com", (*emails[0]).Email)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_SingleOptIn() {
	suite.mux.HandleFunc("/addressbooks/1/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	emails := make([]*models.EmailToAdd, 0)
	emails = append(emails, &models.EmailToAdd{
		Email:     "test@test.com",
		Variables: map[string]interface{}{"age": 21, "weight": 99},
	})

	suite.NoError(suite.client.Emails.AddressBooks.SingleOptIn(1, emails))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_DoubleOptIn() {
	suite.mux.HandleFunc("/addressbooks/1/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	emails := make([]*models.EmailToAdd, 0)
	emails = append(emails, &models.EmailToAdd{
		Email:     "test@test.com",
		Variables: map[string]interface{}{"age": 21, "weight": 99},
	})
	suite.NoError(suite.client.Emails.AddressBooks.DoubleOptIn(1, emails, "admin@admin.com", "ru", "tpl123"))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_EmailsDelete() {
	suite.mux.HandleFunc("/addressbooks/1/emails", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	emails := []string{"test@test.com"}
	suite.NoError(suite.client.Emails.AddressBooks.DeleteEmails(1, emails))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Delete() {
	suite.mux.HandleFunc("/addressbooks/1", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodDelete, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	suite.NoError(suite.client.Emails.AddressBooks.Delete(1))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_CampaignCost() {
	suite.mux.HandleFunc("/addressbooks/1/cost", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"cur": "RUR",
			"sent_emails_qty": 1,
			"overdraftAllEmailsPrice": 0,
			"addressesDeltaFromBalance": 0,
			"addressesDeltaFromTariff": 1,
			"max_emails_per_task": 500,
			"result": true
		}`)
	})
	cost, err := suite.client.Emails.AddressBooks.CampaignCost(1)
	suite.NoError(err)
	suite.NotNil(cost)
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_EmailsUnsubscribe() {
	suite.mux.HandleFunc("/addressbooks/1/emails/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	emails := []string{"test@test.com"}
	suite.NoError(suite.client.Emails.AddressBooks.Unsubscribe(1, emails))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_Campaigns() {
	suite.mux.HandleFunc("/addressbooks/1/campaigns", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `[
			{
				"task_id": 9147533,
				"task_name": "тест",
				"task_status": 3
			},
			{
				"task_id": 9156025,
				"task_name": "Campaign_6741804_UM99",
				"task_status": 3
			}
		]`)
	})

	tasks, err := suite.client.Emails.AddressBooks.CampaignsList(1, 10, 0)
	suite.NoError(err)
	suite.Equal(2, len(tasks))
}

func (suite *SendpulseTestSuite) TestEmailsService_AddressBooksService_UpdateEmailVariables() {
	suite.mux.HandleFunc("/addressbooks/1/emails/variable", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	variables := []*models.Variable{
		{
			Name:  "age",
			Value: 12,
		},
	}
	suite.NoError(suite.client.Emails.AddressBooks.UpdateEmailVariables(1, "test@test.com", variables))
}
