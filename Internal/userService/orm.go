package userService

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex;size:191"`
	Password string `json:"-" gorm:"column:password;size:255"`
	Name     string `json:"name" gorm:"size:100"`
	Role     string `json:"role" gorm:"size:50"`
}

type UserUpdatePayload struct {
	Name  *string
	Email *string
	Role  *string
}
