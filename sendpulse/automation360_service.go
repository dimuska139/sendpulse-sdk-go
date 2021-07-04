package sendpulse

import (
	"fmt"
	"net/http"
)

type Automation360Service struct {
	client *Client
}

func newAutomation360Service(cl *Client) *Automation360Service {
	return &Automation360Service{client: cl}
}

type Autoresponder struct {
	Autoresponder struct {
		ID      int          `json:"id"`
		Name    string       `json:"name"`
		Status  int          `json:"status"`
		Created DateTimeType `json:"created"`
		Changed DateTimeType `json:"changed"`
	} `json:"autoresponder"`
	Flows []*struct {
		ID       int                    `json:"id"`
		MainID   int                    `json:"main_id"`
		AfType   string                 `json:"af_type"`
		Created  DateTimeType           `json:"created"`
		LastSend DateTimeType           `json:"created"`
		Task     map[string]interface{} `json:"task"`
	} `json:"flows"`
	Starts       int `json:"starts"`
	InQueue      int `json:"in_queue"`
	EndCount     int `json:"end_count"`
	SendMessages int `json:"send_messages"`
	Conversions  int `json:"conversions"`
}

func (service *Automation360Service) GetAutoresponderStatistics(id int) (*Autoresponder, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d", id)

	var respData *Autoresponder
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData, err
}

func (service *Automation360Service) StartEvent(eventName string, variables map[string]interface{}) error {
	path := fmt.Sprintf("/events/name/%s", eventName)

	var respData struct {
		Result bool `json:"result"`
	}
	_, err := service.client.newRequest(http.MethodPost, fmt.Sprintf(path), variables, &respData, true)
	return err
}

type MainTriggerBlockStat struct {
	FlowID   int `json:"flow_id"`
	Executed int `json:"executed"`
	Deleted  int `json:"deleted"`
}

func (service *Automation360Service) GetStartBlockStatistics(id int) (*MainTriggerBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/main-trigger/%d/group-stat", id)

	var respData struct {
		Data *MainTriggerBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type EmailBlockStat struct {
	FlowID int `json:"flow_id"`
	Task   struct {
		ID                int          `json:"id"`
		AddressBookID     int          `json:"address_book_id"`
		MessageTitle      string       `json:"message_title"`
		SenderMailAddress string       `json:"sender_mail_address"`
		SenderMailName    string       `json:"sender_mail_name"`
		Created           DateTimeType `json:"created"`
	} `json:"task"`
	Sent         int          `json:"sent"`
	Delivered    int          `json:"delivered"`
	Opened       int          `json:"opened"`
	Clicked      int          `json:"clicked"`
	Errors       int          `json:"errors"`
	Unsubscribed int          `json:"unsubscribed"`
	MarkedAsSpam int          `json:"marked_as_spam"`
	LastSend     DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetEmailBlockStatistics(id int) (*EmailBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/email/%d/group-stat", id)

	var respData struct {
		Data *EmailBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type PushBlockStat struct {
	FlowID    int          `json:"flow_id"`
	Sent      int          `json:"sent"`
	Delivered int          `json:"delivered"`
	Clicked   int          `json:"clicked"`
	Errors    int          `json:"errors"`
	LastSend  DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetPushBlockStatistics(id int) (*PushBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/push/%d/group-stat", id)

	var respData struct {
		Data *PushBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type SmsBlockStat struct {
	FlowID    int          `json:"flow_id"`
	Executed  int          `json:"executed"`
	Sent      int          `json:"sent"`
	Delivered int          `json:"delivered"`
	Opened    int          `json:"opened"`
	Clicked   int          `json:"clicked"`
	Errors    int          `json:"errors"`
	LastSend  DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetSmsBlockStatistics(id int) (*SmsBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/sms/%d/group-stat", id)

	var respData struct {
		Data *SmsBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type MessengerBlockStat struct {
	FlowID   int          `json:"flow_id"`
	Executed int          `json:"executed"`
	Sent     int          `json:"sent"`
	LastSend DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetMessengerBlockStatistics(id int) (*MessengerBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/messenger/%d/group-stat", id)

	var respData struct {
		Data *MessengerBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type FilterBlockStat struct {
	FlowID   int          `json:"flow_id"`
	Executed int          `json:"executed"`
	LastSend DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetFilterBlockStatistics(id int) (*FilterBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/filter/%d/group-stat", id)

	var respData struct {
		Data *FilterBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type TriggerBlockStat struct {
	FlowID   int          `json:"flow_id"`
	Executed int          `json:"executed"`
	LastSend DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetTriggerBlockStatistics(id int) (*TriggerBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/trigger/%d/group-stat", id)

	var respData struct {
		Data *TriggerBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type GoalBlockStat struct {
	FlowID    int          `json:"flow_id"`
	Executed  int          `json:"executed"`
	Sent      int          `json:"sent"`
	Delivered int          `json:"delivered"`
	Opened    int          `json:"opened"`
	Clicked   int          `json:"clicked"`
	Errors    int          `json:"errors"`
	LastSend  DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetGoalBlockStatistics(id int) (*GoalBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/goal/%d/group-stat", id)

	var respData struct {
		Data *GoalBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type ActionBlockStat struct {
	FlowID   int          `json:"flow_id"`
	Executed int          `json:"executed"`
	LastSend DateTimeType `json:"last_send"`
}

func (service *Automation360Service) GetActionBlockStatistics(id int) (*ActionBlockStat, error) {
	path := fmt.Sprintf("/a360/stats/action/%d/group-stat", id)

	var respData struct {
		Data *ActionBlockStat `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type AutoresponderConversion struct {
	TotalConversions       int `json:"total_conversions"`
	MaintriggerConversions int `json:"maintrigger_conversions"`
	GoalConversions        int `json:"goal_conversions"`
	Maintrigger            struct {
		ID          int          `json:"id"`
		MainID      int          `json:"main_id"`
		AfType      string       `json:"af_type"`
		Created     DateTimeType `json:"created"`
		LastSend    DateTimeType `json:"last_send"`
		Conversions int          `json:"conversions"`
	} `json:"maintrigger"`
	Goals []struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		MainID      int          `json:"main_id"`
		AfType      string       `json:"af_type"`
		Created     DateTimeType `json:"created"`
		Conversions int          `json:"conversions"`
	} `json:"goals"`
}

func (service *Automation360Service) GetAutoresponderConversions(id int) (*AutoresponderConversion, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d/conversions", id)

	var respData struct {
		Data *AutoresponderConversion `json:"data"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Data, err
}

type AutoresponderContact struct {
	ID             int          `json:"id"`
	ConversionType string       `json:"conversion_type"`
	FlowID         int          `json:"flow_id"`
	Email          string       `json:"email"`
	Phone          string       `json:"phone"`
	ConversionDate DateTimeType `json:"conversion_date"`
	StartDate      DateTimeType `json:"start_date"`
}

func (service *Automation360Service) GetAutoresponderContacts(id int) ([]*AutoresponderContact, error) {
	path := fmt.Sprintf("/a360/autoresponders/%d/conversions/list/all", id)

	var respData struct {
		Items []*AutoresponderContact `json:"items"`
	}
	_, err := service.client.newRequest(http.MethodGet, fmt.Sprintf(path), nil, &respData, true)
	return respData.Items, err
}
