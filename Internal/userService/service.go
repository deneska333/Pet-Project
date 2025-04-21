package userService

import "gorm.io/gorm"

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id uint, update User) (*User, error) {
	existing, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, gorm.ErrRecordNotFound
	}

	if update.Email != "" {
		existing.Email = update.Email
	}
	if update.Password != "" {
		existing.Password = update.Password
	}

	return s.repo.UpdateUser(*existing)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUserByID(id)
}
