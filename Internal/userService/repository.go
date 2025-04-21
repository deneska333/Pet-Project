package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(user User) (*User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) UpdateUser(user User) (*User, error) {
	err := r.db.Save(&user).Error
	return &user, err
}

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
