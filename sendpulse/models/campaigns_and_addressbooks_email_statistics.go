package models

type CampaignsAndAddressBooksEmailStatistics struct {
	Sent         int `json:"sent"`
	Open         int `json:"open"`
	Link         int `json:"link"`
	Addressbooks []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"adressbooks"`
	Blacklist bool
}
