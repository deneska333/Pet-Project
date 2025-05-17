package userService

import (
	"Pet-project/internal/taskService"
)

type User struct {
	ID       int64              `json:"id" gorm:"primaryKey"`
	Email    string             `json:"email" gorm:"uniqueIndex;size:191"`
	Password string             `json:"-" gorm:"size:255"`
	Name     string             `json:"name" gorm:"size:100"`
	Role     string             `json:"role" gorm:"size:50"`
	Tasks    []taskService.Task `json:"tasks" gorm:"foreignKey:UserID"`
}
type UserUpdatePayload struct {
	Name  *string
	Email *string
	Role  *string
}
