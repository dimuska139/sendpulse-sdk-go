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

	//realID, err := client.Templates.Create("Тестовый шаблон 190621", "<h1>Привет!</h1>", "ru")

	/*err := client.Templates.Update(606652, "<h1>ПрЮвет!</h1>", "ru")
	if err != nil {
		fmt.Println(err)
	}*/
	tpl, _ := client.Emails.Templates.Get(606652)
	fmt.Println(tpl.Tags)
}
