# SendPulse REST client library
A SendPulse REST client library and example for Go (Golang).

API Documentation [https://sendpulse.com/api](https://sendpulse.com/api)

[![Build Status](https://travis-ci.org/dimuska139/sendpulse-sdk-go.svg?branch=master)](https://travis-ci.org/dimuska139/sendpulse-sdk-go)

### Download

```shell
go get -u github.com/dimuska139/sendpulse-sdk-go
```

### Example
```go
package main

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go"
)
const ApiUserId = "12345"
const ApiSecret = "12345"
const ApiTimeout = 5

func main() {
	addressBookId := 12345

	client, e := sendpulse.ApiClient(ApiUserId, ApiSecret, ApiTimeout)
	if e != nil {
		switch err := e.(type) {
		case *sendpulse.HttpError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Message)
		default: // Another errors
			fmt.Println(e)
		}
	}

	// Get address book info by id
	bookInfo, e := client.Books.Get(uint(addressBookId))
	if e != nil {
		switch err := e.(type) {
		case *sendpulse.HttpError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Message)
		default: // Another errors
			fmt.Println(e)
		}
	} else {
		fmt.Println(*bookInfo)
	}

	// Get address books list
	limit := 10
	offset := 20
	books, err := client.Books.List(uint(limit), uint(offset))
	if err != nil {
		switch err := e.(type) {
		case *sendpulse.HttpError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Message)
		default: // Another errors
			fmt.Println(e)
		}
	} else {
		fmt.Println(*books)
	}

	// Add emails to address book
	emails := []sendpulse.Email{
		sendpulse.Email{
			Email:     "alex@test.net",
			Variables: map[string]string{
				"name": "Alex",
				"age": "25",
			},
		},
		sendpulse.Email{
			Email:     "dima@test.net",
			Variables: make(map[string]string),
		},
	}
	
	extraParams := make(map[string]string)
	
	err = client.Books.AddEmails(uint(addressBookId), emails, extraParams)
	if err != nil {
		switch err := e.(type) {
		case *sendpulse.HttpError: // Http error
			fmt.Println(err.HttpCode)
			fmt.Println(err.Url)
			fmt.Println(err.Message)
		default: // Another errors
			fmt.Println(e)
		}
	}
}
```
