package main

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse"
	"net/http"
)

func main() {
	config := &sendpulse.Config{
		UserID: "",
		Secret: "",
	}
	client := sendpulse.NewClient(http.DefaultClient, config)
	/*id, err := client.AddressBooks.Create("12345 book")
	if err != nil {
		fmt.Println(err)
	}

	if err := client.AddressBooks.Update(id, "new name"); err != nil {
		fmt.Println(err)
	}*/

	webhooks, err := client.Emails.Webhooks.Create([]string{"spam", "redirect"}, "https://test.ru")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*webhooks[0])
	fmt.Println(*webhooks[1])
}
