package models

type MailingDto struct {
	SenderName    string            `json:"sender_name"`
	SenderEmail   string            `json:"sender_email"`
	Subject       string            `json:"subject"`
	Body          *string           `json:"body,omitempty"`
	TemplateID    *string           `json:"template_id,omitempty"`
	AddressBookID *int              `json:"list_id,omitempty"`
	SegmentID     *int              `json:"segment_id,omitempty"`
	IsTest        *bool             `json:"is_test,omitempty"`
	SendDate      *DateTimeType     `json:"send_date,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Attachments   map[string]string `json:"attachments"`
	Type          *string           `json:"type,omitempty"`
	BodyAMP       *string           `json:"body_amp,omitempty"`
}
