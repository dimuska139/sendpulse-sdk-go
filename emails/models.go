package emails

import (
	"encoding/json"
	"github.com/dimuska139/sendpulse-sdk-go/types"
	"time"
)

type Book struct {
	ID               *json.Number    `json:"id,omitempty"`
	Name             *string         `json:"name,omitempty"`
	AllEmailQty      *json.Number    `json:"all_email_qty,omitempty"`
	ActiveEmailQty   *json.Number    `json:"active_email_qty,omitempty"`
	InactiveEmailQty *json.Number    `json:"inactive_email_qty,omitempty"`
	CreationDate     *types.DateTime `json:"creationdate,omitempty"`
	Status           *json.Number    `json:"status,omitempty"`
	StatusExplain    *string         `json:"status_explain,omitempty"`
}

type Variable struct {
	Name  *string      `json:"name,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}

type EmailDetailed struct {
	Email         *string                 `json:"email,omitempty"`
	Status        *json.Number            `json:"status,omitempty"`
	StatusExplain *string                 `json:"status_explain,omitempty"`
	Variables     *map[string]interface{} `json:"variables,omitempty"`
}

type Email struct {
	Email     string
	Variables map[string]interface{}
}

type CampaignCost struct {
	Cur                       *string      `json:"email,omitempty"`
	SentEmailsQty             *json.Number `json:"sent_emails_qty,omitempty"`
	OverdraftAllEmailsPrice   *json.Number `json:"overdraft_all_emails_price,omitempty"`
	AddressesDeltaFromBalance *json.Number `json:"address_delta_from_balance,omitempty"`
	AddressesDeltaFromTariff  *json.Number `json:"address_delta_from_tariff,omitempty"`
	MaxEmailsPerTask          *json.Number `json:"max_emails_per_task,omitempty"`
	Result                    *bool        `json:"result,omitempty"`
}

type CreateCampaignDto struct {
	SenderName   string
	SenderEmail  string
	Subject      string
	Body         string
	TemplateID   int
	BodyAMP      string
	ListID       int
	SegmentID    int
	SendTestOnly []string
	SendDate     types.DateTime
	Name         string
	Attachments  map[string]string
	IsDraft      bool
}

type UpdateCampaignDto struct {
	ID          int
	Name        string
	SenderName  string
	SenderEmail string
	Subject     string
	Body        string
	TemplateID  int
	SendDate    time.Time
}

type MessageInfo struct {
	SenderName  string      `json:"sender_name"`
	SenderEmail string      `json:"sender_email"`
	Subject     string      `json:"subject"`
	Body        string      `json:"body"`
	Attachments string      `json:"attachments"`
	ListID      json.Number `json:"list_id"`
}

type Campaign struct {
	ID                int
	Name              string
	Message           MessageInfo
	Status            int
	AllEmailQty       int
	TariffEmailQty    int
	PaidEmailQty      int
	OverdraftPrice    int
	OverdraftCurrency string
}

type CampaignStatisticsCounts struct {
	Code    int
	Count   int
	Explain string
}

type CampaignInfo struct {
	ID                int
	Name              string
	Message           MessageInfo
	Status            int
	AllEmailQty       int
	TariffEmailQty    int
	PaidEmailQty      int
	OverdraftPrice    int
	OverdraftCurrency string
}

type CampaignFullInfo struct {
	CampaignInfo
	Statistics []CampaignStatisticsCounts
	SendDate   time.Time
	Permalink  string
}

type Task struct {
	TaskID     int    `json:"task_id"`
	TaskName   string `json:"task_name"`
	TaskStatus int    `json:"task_status"`
}

type ReferralsStatistics struct {
	Link  string
	Count int
}

type TemplateCategoryInfo struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	MetaDescription string `json:"meta_description,omitempty"`
	FullDescription string `json:"full_description,omitempty"`
	Code            string `json:"code,omitempty"`
	Sort            int    `json:"sort,omitempty"`
}

type Sender struct {
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty"`
	Status *string `json:"status,omitempty"`
}

type Template struct {
	ID              *string `json:"id,omitempty"`
	RealID          *int    `json:"real_id,omitempty"`
	Name            *string `json:"name,omitempty"`
	NameSlug        *string `json:"name_slug,omitempty"`
	MetaDescription *string `json:"meta_description,omitempty"`
	FullDescription *string `json:"full_description,omitempty"`
	Category        *string `json:"category,omitempty"`
	//CategoryInfo    *TemplateCategoryInfo `json:"category_info,omitempty"`
	Mark      *string `json:"mark,omitempty"`
	MarkCount *int    `json:"mark_count,omitempty"`
	Body      *string `json:"body,omitempty"`
	//Tags            *map[string]string `json:"tags,omitempty"`
	Owner   *string         `json:"owner,omitempty"`
	Created *types.DateTime `json:"created,omitempty"`
	Preview *string         `json:"preview,omitempty"`
}

type Balance struct {
	Currency        *string  `json:"currency,omitempty"`
	BalanceCurrency *float32 `json:"balance_currency,omitempty"`
}

type BalanceDetailed struct {
	Balance *struct {
		Main     *json.Number `json:"main,omitempty"`
		Bonus    *json.Number `json:"bonus,omitempty"`
		Currency *string      `json:"currency,omitempty"`
	} `json:"balance,omitempty"`
	Email *struct {
		TariffName         *string         `json:"tariff_name,omitempty"`
		FinishedTime       *types.DateTime `json:"finished_time,omitempty"`
		EmailsLeft         *int            `json:"emails_left,omitempty"`
		MeximumSubscribers *int            `json:"maximum_subscribers,omitempty"`
		CurrentSubscribers *int            `json:"current_subscribers,omitempty"`
	} `json:"email,omitempty"`
	Smtp *struct {
		TariffName *string         `json:"tariff_name,omitempty"`
		EndDate    *types.DateTime `json:"end_date,omitempty"`
		AutoRenew  *int            `json:"auto_renew,omitempty"`
	} `json:"smtp,omitempty"`
	Push *struct {
		TariffName *string     `json:"tariff_name,omitempty"`
		EndDate    *types.Date `json:"end_date,omitempty"`
		AutoRenew  *int        `json:"auto_renew,omitempty"`
	} `json:"push,omitempty"`
}

type EmailInfo struct {
	BookID    *json.Number `json:"book_id,omitempty"`
	Email     *string      `json:"email,omitempty"`
	Status    *json.Number `json:"status,omitempty"`
	Variables []*Variable  `json:"variables,omitempty"`
}

type EmailInfoDetails struct {
	ListName *string         `json:"list_name,omitempty"`
	ListID   *json.Number    `json:"list_id,omitempty"`
	AddDate  *types.DateTime `json:"add_date,omitempty"`
	Source   *string         `json:"source,omitempty"`
}

type EmailCampaignStatistics struct {
	SentDate            *types.DateTime `json:"sent_date,omitempty"`
	GlobalStatus        *int            `json:"global_status,omitempty"`
	GlobalStatusExplain *string         `json:"global_status_explain,omitempty"`
}

type EmailCampaignsStatistics struct {
	Statistic *struct {
		Sent *int `json:"sent,omitempty"`
		Open *int `json:"open,omitempty"`
		Link *int `json:"link,omitempty"`
	} `json:"statistic,omitempty"`
	Blacklist *bool
}

type EmailCampaignsStatisticsDetails struct {
	Sent         *int `json:"sent,omitempty"`
	Open         *int `json:"open,omitempty"`
	Link         *int `json:"link,omitempty"`
	AddressBooks []*struct {
		ID   *int    `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"adressbooks,omitempty"`
}

type EmailAddressbookStatistics struct {
	Email         *string      `json:"email,omitempty"`
	AddressBookID *json.Number `json:"abook_id,omitempty"`
	Status        *json.Number `json:"status,omitempty"`
	StatusExplain *string      `json:"status_explain,omitempty"`
	Variables     []*Variable  `json:"variables,omitempty"`
}

type Webhook struct {
	ID     *int    `json:"id,omitempty"`
	UserID *int    `json:"user_id,omitempty"`
	Url    *string `json:"url,omitempty"`
	Action *string `json:"action,omitempty"`
}
