package sendpulse_sdk_go

import (
	"fmt"
	"net/http"
)

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_ValidateAddressBook() {
	suite.mux.HandleFunc("/verifier-service/send-list-to-verify/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{"result": true}`)
	})

	err := suite.client.Emails.Validator.ValidateMailingList(1)
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_GetAddressBookValidationProgress() {
	suite.mux.HandleFunc("/verifier-service/get-progress/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"result": true,
			"data": {
				"total": 25,
				"processed": 19
			}
		}`)
	})

	progress, err := suite.client.Emails.Validator.GetMailingListValidationProgress(1)
	suite.NoError(err)
	suite.Equal(25, progress.Total)
	suite.Equal(19, progress.Processed)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_GetAddressBookValidationResult() {
	suite.mux.HandleFunc("/verifier-service/check/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"id": 11,
			"address_book_name": "12345 book",
			"all_emails_quantity": 1,
			"status": 1,
			"check_date": "2021-06-27 16:03:42",
			"data": {
				"0": 0,
				"1": 1,
				"2": 0,
				"3": 0
			},
			"is_updated": 0,
			"status_text": "Зелёный",
			"email_addresses": [
				{
					"id": 12345,
					"email_address": "test@sendpulse.com",
					"check_date": "2021-06-27 15:45:18",
					"status": 1,
					"status_text": "Действительный адрес"
				}
			],
			"email_addresses_total": 1
		}`)
	})

	result, err := suite.client.Emails.Validator.GetMailingListValidationResult(1)
	suite.NoError(err)
	suite.Equal(11, result.ID)
	suite.Equal("test@sendpulse.com", result.EmailAddresses[0].EmailAddress)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_GetValidatedAddressBooksList() {
	suite.mux.HandleFunc("/verifier-service/check-list", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"total": 1,
			"list": [
				{
					"id": 1266227,
					"address_book_name": "12345 book",
					"all_emails_quantity": 1,
					"status": 1,
					"check_date": "2021-06-27 16:03:42",
					"data": {
						"0": 0,
						"1": 1,
						"2": 0,
						"3": 0
					},
					"is_updated": 0,
					"status_text": "Зелёный",
					"is_garbage_in_book": true
				}
			]
		}`)
	})

	list, err := suite.client.Emails.Validator.GetValidatedMailingLists(10, 0)
	suite.NoError(err)
	suite.Equal("12345 book", list[0].Name)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_ValidateEmail() {
	suite.mux.HandleFunc("/verifier-service/send-single-to-verify/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.Emails.Validator.ValidateEmail("test@sendpulse.com")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_GetEmailValidationResult() {
	suite.mux.HandleFunc("/verifier-service/get-single-result/", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"result": true,
			"data": {
				"email": "test@sendpulse.com",
				"checks": {
					"status": 1,
					"valid_format": 1,
					"disposable": 0,
					"webmail": 0,
					"gibberish": 0,
					"status_text": "Действительный адрес"
				}
			}
		}`)
	})

	result, err := suite.client.Emails.Validator.GetEmailValidationResult("test@sendpulse.com")
	suite.NoError(err)
	suite.Equal("test@sendpulse.com", result.Email)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_DeleteEmailValidationResult() {
	suite.mux.HandleFunc("/verifier-service/delete-single-result", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.Emails.Validator.DeleteEmailValidationResult("test@sendpulse.com")
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_CreateAddressBookValidationReport() {
	suite.mux.HandleFunc("/verifier-service/make-report", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodPost, r.Method)
		fmt.Fprintf(w, `{
			"result": true
		}`)
	})

	err := suite.client.Emails.Validator.CreateMailingListValidationReport(MailingListReportParams{
		ID:       25,
		Statuses: []int{1, 2, 3},
	})
	suite.NoError(err)
}

func (suite *SendpulseTestSuite) TestEmailsService_ValidatorService_GetAddressBookValidationReport() {
	suite.mux.HandleFunc("/verifier-service/check-report", func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(http.MethodGet, r.Method)
		fmt.Fprintf(w, `{
    		"id": 1,
			"address_book_name": "12345 book",
			"all_emails_quantity": 1,
			"status": 0,
			"check_date": "2021-06-27 15:46:23",
			"data": {},
			"is_updated": 1,
			"status_text": "Идет проверка, обновлён",
			"email_addresses": [
				{
					"id": 1,
					"email_address": "test@sendpulse.com",
					"check_date": "2021-06-27 15:45:18",
					"status": 1,
					"status_text": "Действительный адрес"
				}
			],
			"email_addresses_total": 1
		}`)
	})

	report, err := suite.client.Emails.Validator.GetMailingListValidationReport(1)
	suite.NoError(err)
	suite.Equal("12345 book", report.Name)
	suite.Equal("test@sendpulse.com", report.EmailAddresses[0].EmailAddress)
}
