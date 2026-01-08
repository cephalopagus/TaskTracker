package main

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"Title"`
	Status      string    `json:"Status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type TaskList []Task
