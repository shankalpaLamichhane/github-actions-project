package models

import "time"

// declaring models
type Task struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Task) TableName() string {
	return "tasks"
}
