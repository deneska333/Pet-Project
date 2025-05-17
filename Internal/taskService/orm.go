package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task" gorm:"column:task;not null;size:255"`
	Text   string `json:"text" gorm:"column:text"`
	IsDone bool   `json:"is_done" gorm:"default:false"`
	UserID uint   `json:"user_id" gorm:"not null"`
}
