package main

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"Title"`
	Status      string    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}
type TaskList []Task
