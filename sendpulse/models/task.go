package models

type Task struct {
	ID     int    `json:"task_id"`
	Name   string `json:"task_name"`
	Status int    `json:"task_status"`
}
