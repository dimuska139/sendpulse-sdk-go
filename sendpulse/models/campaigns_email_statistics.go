package models

type CampaignsEmailStatistics struct {
	Statistic *struct {
		Sent int `json:"sent"`
		Open int `json:"open"`
		Link int `json:"link"`
	} `json:"statistic"`
	Addressbooks []*struct {
		Id   int    `json:"int"`
		Name string `json:"address_book_name"`
	}
	Blacklist bool
}
