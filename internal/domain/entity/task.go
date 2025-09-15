package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TaskId   string    `gorm:"primaryKey;" json:"task_id"`
	Title    string    `json:"title"`
	Details  string    `json:"details"`
	Deadline time.Time `json:"deadline"`
	Status   string    `gorm:"default:'pending'" json:"status"`
	UserId   string    `json:"user_id"`
	gorm.Model
}
