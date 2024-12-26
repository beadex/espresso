package models

import (
	"time"
)

type Task struct {
	ID          int
	Name        string
	DueDate     time.Time
	IsRecurring bool
}
