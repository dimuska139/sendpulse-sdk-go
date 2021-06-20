package models

type Task struct {
	TaskID     int    `json:"task_id"`
	TaskName   string `json:"task_name"`
	TaskStatus int    `json:"task_status"`
}
