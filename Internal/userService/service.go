package userService

import (
	"errors"
	"fmt"
	"log"

	"Pet-project/internal/taskService"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(id uint) (User, error)
	GetAllUsers() ([]User, error)
	CreateUser(user User) (User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, payload UserUpdatePayload) (User, error)
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}

type ErrUserNotFound struct{ ID uint }

func (e *ErrUserNotFound) Error() string { return fmt.Sprintf("user with ID %d not found", e.ID) }

type ErrEmailConflict struct{ Email string }

func (e *ErrEmailConflict) Error() string { return fmt.Sprintf("email '%s' is already taken", e.Email) }

type service struct {
	repo UserRepository
}

func NewService(r UserRepository) UserService {
	if r == nil {
		log.Fatal("UserRepository cannot be nil in NewService")
	}
	return &service{repo: r}
}

func (s *service) GetUserByID(id uint) (User, error) {
	userPtr, err := s.repo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, &ErrUserNotFound{ID: id}
		}
		log.Printf("Error fetching user by ID %d from repository: %v", id, err)
		return User{}, fmt.Errorf("database error fetching user %d", id)
	}
	if userPtr == nil {
		return User{}, &ErrUserNotFound{ID: id}
	}
	return *userPtr, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching all users from repository: %v", err)
		return nil, fmt.Errorf("database error fetching all users")
	}
	return users, nil
}

func (s *service) CreateUser(user User) (User, error) {
	_, findErr := s.repo.FindByEmail(user.Email)
	if findErr == nil {
		return User{}, &ErrEmailConflict{Email: user.Email}
	} else if !errors.Is(findErr, gorm.ErrRecordNotFound) {
		log.Printf("Error checking email %s existence: %v", user.Email, findErr)
		return User{}, fmt.Errorf("database error checking email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password for email %s: %v", user.Email, err)
		return user, fmt.Errorf("internal error processing password")
	}
	user.Password = string(hashedPassword)

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user with email %s in repository: %v", user.Email, err)
		return User{}, fmt.Errorf("database error creating user")
	}
	return createdUser, nil
}

func (s *service) DeleteUser(id uint) error {
	err := s.repo.DeleteUserByID(id)
	if err != nil {
		_, getErr := s.repo.GetUserByID(id)
		if getErr != nil {
			if errors.Is(getErr, gorm.ErrRecordNotFound) {
				return &ErrUserNotFound{ID: id}
			}
			log.Printf("Error checking user %d before delete: %v", id, getErr)
			return fmt.Errorf("database error checking user before delete")
		}
		log.Printf("Error deleting user %d from repository: %v", id, err)
		return fmt.Errorf("database error deleting user")
	}
	return nil
}

func (s *service) UpdateUser(id uint, payload UserUpdatePayload) (User, error) {
	currentUserPtr, err := s.repo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, &ErrUserNotFound{ID: id}
		}
		log.Printf("Error fetching user %d for update: %v", id, err)
		return User{}, fmt.Errorf("database error fetching user for update")
	}
	currentUser := *currentUserPtr

	userModified := false
	if payload.Name != nil && *payload.Name != currentUser.Name {
		currentUser.Name = *payload.Name
		userModified = true
	}
	if payload.Role != nil && *payload.Role != currentUser.Role {
		currentUser.Role = *payload.Role
		userModified = true
	}
	if payload.Email != nil && *payload.Email != currentUser.Email {
		_, findErr := s.repo.FindByEmail(*payload.Email)
		if findErr == nil {
			return User{}, &ErrEmailConflict{Email: *payload.Email}
		} else if !errors.Is(findErr, gorm.ErrRecordNotFound) {
			log.Printf("Error checking email conflict for %s: %v", *payload.Email, findErr)
			return User{}, fmt.Errorf("database error checking email conflict")
		}
		currentUser.Email = *payload.Email
		userModified = true
	}

	if !userModified {
		return currentUser, nil
	}

	updatedUserPtr, err := s.repo.UpdateUser(currentUser)
	if err != nil {
		log.Printf("Error updating user %d in repository: %v", id, err)
		return User{}, fmt.Errorf("database error updating user")
	}
	if updatedUserPtr == nil {
		log.Printf("Repository returned nil user after update for ID %d", id)
		return User{}, fmt.Errorf("internal error: repository returned nil after update")
	}
	return *updatedUserPtr, nil
}

func (s *service) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	tasks, err := s.repo.GetTasksForUser(userID)
	if err != nil {
		log.Printf("Error fetching tasks for user %d: %v", userID, err)
		return nil, fmt.Errorf("database error fetching user tasks")
	}
	return tasks, nil
}

var _ UserService = (*service)(nil)
