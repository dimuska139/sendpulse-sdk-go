package automation360

import (
	"github.com/dimuska139/sendpulse-sdk-go/types"
)

type Autoresponder struct {
	Autoresponder *struct {
		ID      *types.Int      `json:"id,omitempty"`
		Name    *string         `json:"name,omitempty"`
		Status  *types.Int      `json:"status,omitempty"`
		Created *types.DateTime `json:"created,omitempty"`
		Changed *types.DateTime `json:"changed,omitempty"`
	} `json:"autoresponder,omitempty"`
	Flows []*struct {
		ID       *types.Int              `json:"id,omitempty"`
		MainID   *types.Int              `json:"main_id,omitempty"`
		AfType   *string                 `json:"af_type,omitempty"`
		Created  *types.DateTime         `json:"created,omitempty"`
		LastSend *types.DateTime         `json:"created,omitempty"`
		Task     *map[string]interface{} `json:"task,omitempty"`
	} `json:"flows,omitempty"`
	Starts       *types.Int `json:"starts,omitempty"`
	InQueue      *types.Int `json:"in_queue,omitempty"`
	EndCount     *types.Int `json:"end_count,omitempty"`
	SendMessages *types.Int `json:"send_messages,omitempty"`
	Conversions  *types.Int `json:"conversions,omitempty"`
}

type MainTriggerBlockStat struct {
	FlowID   *types.Int `json:"flow_id,omitempty"`
	Executed *types.Int `json:"executed,omitempty"`
	Deleted  *types.Int `json:"deleted,omitempty"`
}

type MainTriggerBlockRecipient struct {
	FlowID        *types.Int      `json:"flow_id,omitempty"`
	Email         *string         `json:"email,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	EventID       *string         `json:"event_id,omitempty"`
	EmailBase64   *string         `json:"email_b64,omitempty"`
	ExecutionDate *types.DateTime `json:"execution_date,omitempty"`
}

type EmailBlockStat struct {
	FlowID *types.Int `json:"flow_id,omitempty"`
	Task   *struct {
		ID                *types.Int      `json:"id,omitempty"`
		AddressBookID     *types.Int      `json:"address_book_id,omitempty"`
		MessageTitle      *string         `json:"message_title,omitempty"`
		SenderMailAddress *string         `json:"sender_mail_address,omitempty"`
		SenderMailName    *string         `json:"sender_mail_name,omitempty"`
		Created           *types.DateTime `json:"created,omitempty"`
	} `json:"task,omitempty"`
	Sent         *types.Int      `json:"sent,omitempty"`
	Delivered    *types.Int      `json:"delivered,omitempty"`
	Opened       *types.Int      `json:"opened,omitempty"`
	Clicked      *types.Int      `json:"clicked,omitempty"`
	Errors       *types.Int      `json:"errors,omitempty"`
	Unsubscribed *types.Int      `json:"unsubscribed,omitempty"`
	MarkedAsSpam *types.Int      `json:"marked_as_spam,omitempty"`
	LastSend     *types.DateTime `json:"last_send,omitempty"`
}

type EmailBlockRecipient struct {
	ID                         *types.Int      `json:"id,omitempty"`
	Email                      *string         `json:"email,omitempty"`
	EmailBase64                *string         `json:"email_b64,omitempty"`
	EventID                    *string         `json:"event_id,omitempty"`
	DeliveredStatus            *types.Int      `json:"delivered_status,omitempty"`
	DeliveredStatusDescription *string         `json:"delivered_status_description,omitempty"`
	IsSpam                     *types.Int      `json:"is_spam,omitempty"`
	IsUnsubscribe              *types.Int      `json:"is_unsubscribe,omitempty"`
	Phone                      *string         `json:"phone,omitempty"`
	SentDate                   *types.DateTime `json:"sent_date,omitempty"`
	DeliveredDate              *types.DateTime `json:"delivered_date,omitempty"`
	OpenDate                   *types.DateTime `json:"open_date,omitempty"`
	RedirectDate               *types.DateTime `json:"redirect_date,omitempty"`
	Updated                    *types.DateTime `json:"updated,omitempty"`
}

type PushBlockStat struct {
	FlowID    *types.Int      `json:"flow_id,omitempty"`
	Sent      *types.Int      `json:"sent,omitempty"`
	Delivered *types.Int      `json:"delivered,omitempty"`
	Clicked   *types.Int      `json:"clicked,omitempty"`
	Errors    *types.Int      `json:"errors,omitempty"`
	LastSend  *types.DateTime `json:"last_send,omitempty"`
}

type PushBlockRecipient struct {
	ID           *types.Int      `json:"id,omitempty"`
	Email        *string         `json:"email,omitempty"`
	EmailBase64  *string         `json:"email_b64,omitempty"`
	EventID      *string         `json:"event_id,omitempty"`
	Phone        *string         `json:"phone,omitempty"`
	Status       *types.Int      `json:"status,omitempty"`
	IsSent       *types.Int      `json:"is_spam,omitempty"`
	IsDelivered  *types.Int      `json:"is_delivered,omitempty"`
	IsRedirected *types.Int      `json:"is_redirected,omitempty"`
	SentDate     *types.DateTime `json:"sent_date,omitempty"`
}

type SmsBlockStat struct {
	FlowID    *types.Int      `json:"flow_id,omitempty"`
	Executed  *types.Int      `json:"executed,omitempty"`
	Sent      *types.Int      `json:"sent,omitempty"`
	Delivered *types.Int      `json:"delivered,omitempty"`
	Opened    *types.Int      `json:"opened,omitempty"`
	Clicked   *types.Int      `json:"clicked,omitempty"`
	Errors    *types.Int      `json:"errors,omitempty"`
	LastSend  *types.DateTime `json:"last_send,omitempty"`
}

type SmsBlockRecipient struct {
	ID       *types.Int                `json:"id,omitempty"`
	Phone    *string                   `json:"phone,omitempty"`
	Status   *types.Int                `json:"status,omitempty"`
	EventID  *string                   `json:"event_id,omitempty"`
	Sender   *string                   `json:"sender,omitempty"`
	Body     *string                   `json:"body,omitempty"`
	Price    *map[string]types.Float32 `json:"price,omitempty"`
	Cur      *string                   `json:"cur,omitempty"`
	Email    *string                   `json:"email,omitempty"`
	SentDate *types.DateTime           `json:"sent_date,omitempty"`
}

type FilterBlockStat struct {
	FlowID   *types.Int      `json:"flow_id,omitempty"`
	Executed *types.Int      `json:"executed,omitempty"`
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
	FlowID   *types.Int      `json:"flow_id,omitempty"`
	Executed *types.Int      `json:"executed,omitempty"`
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
	FlowID    *types.Int      `json:"flow_id,omitempty"`
	Executed  *types.Int      `json:"executed,omitempty"`
	Sent      *types.Int      `json:"sent,omitempty"`
	Delivered *types.Int      `json:"delivered,omitempty"`
	Opened    *types.Int      `json:"opened,omitempty"`
	Clicked   *types.Int      `json:"clicked,omitempty"`
	Errors    *types.Int      `json:"errors,omitempty"`
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
	FlowID   *types.Int      `json:"flow_id,omitempty"`
	Executed *types.Int      `json:"executed,omitempty"`
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
	TotalConversions        *types.Int `json:"total_conversions,omitempty"`
	MainTrigggerConversions *types.Int `json:"maintrigger_conversions,omitempty"`
	GoalConversions         *types.Int `json:"goal_conversions,omitempty"`
	Maintrigger             *struct {
		ID          *types.Int      `json:"id,omitempty"`
		Name        *string         `json:"name,omitempty"`
		MainID      *types.Int      `json:"main_id,omitempty"`
		AfType      *string         `json:"af_type,omitempty"`
		Created     *types.DateTime `json:"created,omitempty"`
		LastSend    *types.DateTime `json:"last_send,omitempty"`
		Conversions *types.Int      `json:"conversions,omitempty"`
	} `json:"maintrigger,omitempty"`
	Goals []*struct {
		ID          *types.Int      `json:"id,omitempty"`
		Name        *string         `json:"name,omitempty"`
		MainID      *types.Int      `json:"main_id,omitempty"`
		AfType      *string         `json:"af_type,omitempty"`
		Created     *types.DateTime `json:"created,omitempty"`
		LastSend    *types.DateTime `json:"last_send,omitempty"`
		Conversions *types.Int      `json:"conversions,omitempty"`
	} `json:"goals,omitempty"`
}

type ConversionContact struct {
	ID             *types.Int      `json:"id,omitempty"`
	ConversionType *string         `json:"conversion_type,omitempty"`
	FlowID         *types.Int      `json:"flow_id,omitempty"`
	Email          *string         `json:"email,omitempty"`
	Phone          *string         `json:"phone,omitempty"`
	ConversionDate *types.DateTime `json:"conversion_date,omitempty"`
	StartDate      *types.DateTime `json:"start_date,omitempty"`
}
