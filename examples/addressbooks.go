package main

import (
	"fmt"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse"
	"github.com/dimuska139/sendpulse-sdk-go/sendpulse/models"
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

	emails := make([]*models.EmailToAdd, 0)
	emails = append(emails, &models.EmailToAdd{
		Email:     "dimuska139@yandex.ru",
		Variables: map[string]interface{}{"age": 21, "weight": 99},
	})

	if err := client.Emails.AddressBooks.SingleOptIn(1266208, emails); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*emails[0])
}
