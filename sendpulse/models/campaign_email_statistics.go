package models

type CampaignEmailStatistics struct {
	SentDate            DateTimeType `json:"sent_date"`
	GlobalStatus        int          `json:"global_status"`
	GlobalStatusExplain string       `json:"global_status_explain"`
	DetailStatus        int          `json:"detail_status"`
	DetailStatusExplain string       `json:"detail_status_explain"`
}
