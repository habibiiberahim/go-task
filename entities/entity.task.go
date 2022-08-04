package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	Assignn string
	IsDone bool
	Deadline time.Time
}