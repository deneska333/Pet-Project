package userService

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"size:50"`
	Password string `json:"-" gorm:"column:users_password;size:255"`
}
