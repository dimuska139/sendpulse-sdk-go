# SendPulse REST client library (unofficial)
An unofficial SendPulse REST client library for Go (Golang).

API Documentation [https://sendpulse.com/api](https://sendpulse.com/api)

[![Build Status](https://travis-ci.com/dimuska139/sendpulse-sdk-go.svg?branch=master)](https://travis-ci.org/dimuska139/sendpulse-sdk-go)
[![codecov](https://codecov.io/gh/dimuska139/sendpulse-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/dimuska139/sendpulse-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/sendpulse-sdk-go)](https://goreportcard.com/report/github.com/dimuska139/sendpulse-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/sendpulse-sdk-go/blob/master/LICENSE)

### Installation

```shell
go get -u github.com/dimuska139/sendpulse-sdk-go/sendpulse
```

### Usage
```go
package main

import (
	"fmt"
	sendpulse "github.com/dimuska139/sendpulse-sdk-go/v6"
	"net/http"
)

func main() {
	config := &sendpulse.Config{
		UserID: "",
		Secret: "",
	}
	client := sendpulse.NewClient(http.DefaultClient, config)
	
	emails := make([]*sendpulse.EmailToAdd, 0)
	emails = append(emails, &sendpulse.EmailToAdd{
		Email:     "test@test.com",
		Variables: map[string]interface{}{"age": 21, "weight": 99},
	})

	if err := client.Emails.MailingLists.SingleOptIn(1266208, emails); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*emails[0])
}
```

The tests should be considered a part of the documentation.
