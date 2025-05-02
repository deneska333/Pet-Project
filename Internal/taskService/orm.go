package taskService

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id" gorm:"not null"`
}
