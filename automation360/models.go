package automation360

import (
	"github.com/dimuska139/sendpulse-sdk-go/types"
)

type Autoresponder struct {
	Autoresponder *struct {
		ID      *int            `json:"id,omitempty"`
		Name    *string         `json:"name,omitempty"`
		Status  *int            `json:"status,omitempty"`
		Created *types.DateTime `json:"created,omitempty"`
		Changed *types.DateTime `json:"changed,omitempty"`
	} `json:"autoresponder,omitempty"`
	Flows []*struct {
		ID       *int                    `json:"id,omitempty"`
		MainID   *int                    `json:"main_id,omitempty"`
		AfType   *string                 `json:"af_type,omitempty"`
		Created  *types.DateTime         `json:"created,omitempty"`
		LastSend *types.DateTime         `json:"created,omitempty"`
		Task     *map[string]interface{} `json:"task,omitempty"`
	} `json:"flows,omitempty"`
	Starts       *int `json:"starts,omitempty"`
	InQueue      *int `json:"in_queue,omitempty"`
	EndCount     *int `json:"end_count,omitempty"`
	SendMessages *int `json:"send_messages,omitempty"`
	Conversions  *int `json:"conversions,omitempty"`
}

type MainTriggerBlockStat struct {
	FlowID   *int `json:"flow_id,omitempty"`
	Executed *int `json:"executed,omitempty"`
	Deleted  *int `json:"deleted,omitempty"`
}

type MainTriggerBlockRecipient struct {
	FlowID        *int            `json:"flow_id,omitempty"`
	Email         *string         `json:"email,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	EventID       *string         `json:"event_id,omitempty"`
	EmailBase64   *string         `json:"email_b64,omitempty"`
	ExecutionDate *types.DateTime `json:"execution_date,omitempty"`
}

type EmailBlockStat struct {
	FlowID *int `json:"flow_id,omitempty"`
	Task   *struct {
		ID                *int            `json:"id,omitempty"`
		AddressBookID     *int            `json:"address_book_id,omitempty"`
		MessageTitle      *string         `json:"message_title,omitempty"`
		SenderMailAddress *string         `json:"sender_mail_address,omitempty"`
		SenderMailName    *string         `json:"sender_mail_name,omitempty"`
		Created           *types.DateTime `json:"created,omitempty"`
	} `json:"task,omitempty"`
	Sent         *int            `json:"sent,omitempty"`
	Delivered    *int            `json:"delivered,omitempty"`
	Opened       *int            `json:"opened,omitempty"`
	Clicked      *int            `json:"clicked,omitempty"`
	Errors       *int            `json:"errors,omitempty"`
	Unsubscribed *int            `json:"unsubscribed,omitempty"`
	MarkedAsSpam *int            `json:"marked_as_spam,omitempty"`
	LastSend     *types.DateTime `json:"last_send,omitempty"`
}

type EmailBlockRecipient struct {
	ID                         *int            `json:"id,omitempty"`
	Email                      *string         `json:"email,omitempty"`
	EmailBase64                *string         `json:"email_b64,omitempty"`
	EventID                    *string         `json:"event_id,omitempty"`
	DeliveredStatus            *int            `json:"delivered_status,omitempty"`
	DeliveredStatusDescription *string         `json:"delivered_status_description,omitempty"`
	IsSpam                     *int            `json:"is_spam,omitempty"`
	IsUnsubscribe              *int            `json:"is_unsubscribe,omitempty"`
	Phone                      *string         `json:"phone,omitempty"`
	SentDate                   *types.DateTime `json:"sent_date,omitempty"`
	DeliveredDate              *types.DateTime `json:"delivered_date,omitempty"`
	OpenDate                   *types.DateTime `json:"open_date,omitempty"`
	RedirectDate               *types.DateTime `json:"redirect_date,omitempty"`
	Updated                    *types.DateTime `json:"updated,omitempty"`
}

type PushBlockStat struct {
	FlowID    *int            `json:"flow_id,omitempty"`
	Sent      *int            `json:"sent,omitempty"`
	Delivered *int            `json:"delivered,omitempty"`
	Clicked   *int            `json:"clicked,omitempty"`
	Errors    *int            `json:"errors,omitempty"`
	LastSend  *types.DateTime `json:"last_send,omitempty"`
}

type PushBlockRecipient struct {
	ID           *int            `json:"id,omitempty"`
	Email        *string         `json:"email,omitempty"`
	EmailBase64  *string         `json:"email_b64,omitempty"`
	EventID      *string         `json:"event_id,omitempty"`
	Phone        *string         `json:"phone,omitempty"`
	Status       *int            `json:"status,omitempty"`
	IsSent       *int            `json:"is_spam,omitempty"`
	IsDelivered  *int            `json:"is_delivered,omitempty"`
	IsRedirected *int            `json:"is_redirected,omitempty"`
	SentDate     *types.DateTime `json:"sent_date,omitempty"`
}

type SmsBlockStat struct {
	FlowID    *int            `json:"flow_id,omitempty"`
	Executed  *int            `json:"executed,omitempty"`
	Sent      *int            `json:"sent,omitempty"`
	Delivered *int            `json:"delivered,omitempty"`
	Opened    *int            `json:"opened,omitempty"`
	Clicked   *int            `json:"clicked,omitempty"`
	Errors    *int            `json:"errors,omitempty"`
	LastSend  *types.DateTime `json:"last_send,omitempty"`
}

type SmsBlockRecipient struct {
	ID       *int                      `json:"id,omitempty"`
	Phone    *string                   `json:"phone,omitempty"`
	Status   *int                      `json:"status,omitempty"`
	EventID  *string                   `json:"event_id,omitempty"`
	Sender   *string                   `json:"sender,omitempty"`
	Body     *string                   `json:"body,omitempty"`
	Price    *map[string]types.Float32 `json:"price,omitempty"`
	Cur      *string                   `json:"cur,omitempty"`
	Email    *string                   `json:"email,omitempty"`
	SentDate *types.DateTime           `json:"sent_date,omitempty"`
}

type FilterBlockStat struct {
	FlowID   *int            `json:"flow_id,omitempty"`
	Executed *int            `json:"executed,omitempty"`
	LastSend *types.DateTime `json:"last_send,omitempty"`
}

type FilterBlockRecipient struct {
	Email         *string         `json:"email,omitempty"`
	EmailBase64   *string         `json:"email_b64,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	EventID       *string         `json:"event_id,omitempty"`
	ExecutionDate *types.DateTime `json:"execution_date,omitempty"`
}

type ConditionBlockStat struct {
	FlowID   *int            `json:"flow_id,omitempty"`
	Executed *int            `json:"executed,omitempty"`
	LastSend *types.DateTime `json:"last_send,omitempty"`
}

type ConditionBlockRecipient struct {
	Email         *string         `json:"email,omitempty"`
	EmailBase64   *string         `json:"email_b64,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	EventID       *string         `json:"event_id,omitempty"`
	ExecutionDate *types.DateTime `json:"execution_date,omitempty"`
}

type GoalBlockStat struct {
	FlowID    *int            `json:"flow_id,omitempty"`
	Executed  *int            `json:"executed,omitempty"`
	Sent      *int            `json:"sent,omitempty"`
	Delivered *int            `json:"delivered,omitempty"`
	Opened    *int            `json:"opened,omitempty"`
	Clicked   *int            `json:"clicked,omitempty"`
	Errors    *int            `json:"errors,omitempty"`
	LastSend  *types.DateTime `json:"last_send,omitempty"`
}

type GoalBlockRecipient struct {
	Email         *string         `json:"email,omitempty"`
	EmailBase64   *string         `json:"email_b64,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	EventID       *string         `json:"event_id,omitempty"`
	ExecutionDate *types.DateTime `json:"execution_date,omitempty"`
}

type ActionBlockStat struct {
	FlowID   *int            `json:"flow_id,omitempty"`
	Executed *int            `json:"executed,omitempty"`
	LastSend *types.DateTime `json:"last_send,omitempty"`
}

type ActionBlockRecipient struct {
	Email         *string         `json:"email,omitempty"`
	EmailBase64   *string         `json:"email_b64,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	EventID       *string         `json:"event_id,omitempty"`
	ExecutionDate *types.DateTime `json:"execution_date,omitempty"`
}

type Conversion struct {
	TotalConversions        *int `json:"total_conversions,omitempty"`
	MainTrigggerConversions *int `json:"maintrigger_conversions,omitempty"`
	GoalConversions         *int `json:"goal_conversions,omitempty"`
	Maintrigger             *struct {
		ID          *int            `json:"id,omitempty"`
		Name        *string         `json:"name,omitempty"`
		MainID      *int            `json:"main_id,omitempty"`
		AfType      *string         `json:"af_type,omitempty"`
		Created     *types.DateTime `json:"created,omitempty"`
		LastSend    *types.DateTime `json:"last_send,omitempty"`
		Conversions *int            `json:"conversions,omitempty"`
	} `json:"maintrigger,omitempty"`
	Goals []*struct {
		ID          *int            `json:"id,omitempty"`
		Name        *string         `json:"name,omitempty"`
		MainID      *int            `json:"main_id,omitempty"`
		AfType      *string         `json:"af_type,omitempty"`
		Created     *types.DateTime `json:"created,omitempty"`
		LastSend    *types.DateTime `json:"last_send,omitempty"`
		Conversions *int            `json:"conversions,omitempty"`
	} `json:"goals,omitempty"`
}

type ConversionContact struct {
	ID             *int            `json:"id,omitempty"`
	ConversionType *string         `json:"conversion_type,omitempty"`
	FlowID         *int            `json:"flow_id,omitempty"`
	Email          *string         `json:"email,omitempty"`
	Phone          *string         `json:"phone,omitempty"`
	ConversionDate *types.DateTime `json:"conversion_date,omitempty"`
	StartDate      *types.DateTime `json:"start_date,omitempty"`
}
