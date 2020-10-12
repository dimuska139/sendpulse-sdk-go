# SendPulse REST client library (unofficial)
An unofficial SendPulse REST client library for Go (Golang).

API Documentation [https://sendpulse.com/api](https://sendpulse.com/api)

[![Build Status](https://travis-ci.org/dimuska139/sendpulse-sdk-go.svg?branch=master)](https://travis-ci.org/dimuska139/sendpulse-sdk-go)
[![codecov](https://codecov.io/gh/dimuska139/sendpulse-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/dimuska139/sendpulse-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/sendpulse-sdk-go)](https://goreportcard.com/report/github.com/dimuska139/sendpulse-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/sendpulse-sdk-go/blob/master/LICENSE)

### Installation

```shell
go get -u github.com/dimuska139/sendpulse-sdk-go
```

### Usage
```go
package main

import (
    "fmt"
    "github.com/dimuska139/sendpulse-sdk-go"
    "github.com/dimuska139/sendpulse-sdk-go/emails"
    "net/http"
)
const ApiUserId = "12345"
const ApiSecret = "12345"

func main() {
    config := sendpulse.Config{
        UserID: ApiUserId, 
        Secret: ApiSecret,
	}
	
    client := emails.New(http.DefaultClient, &config)


    // Get address book info by id 
    bookInfo, e := client.GetAddressbook(12345)
    if e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(*bookInfo)
    }
}
```

The tests should be considered a part of the documentation.
