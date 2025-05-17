package taskService

import (
	"errors"

	"gorm.io/gorm"
)

var ErrTaskNotFound = errors.New("task not found")
var ErrUserNotFound = errors.New("task not found")

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks(userID uint) ([]Task, error)
	GetTaskByID(id uint, userID uint) (*Task, error)
	UpdateTask(task Task) (*Task, error)
	DeleteTaskByID(id uint, userID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllTasks(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uint, userID uint) (*Task, error) {
	var task Task
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrTaskNotFound
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) UpdateTask(task Task) (*Task, error) {
	err := r.db.Model(&Task{}).
		Where("id = ? AND user_id = ?", task.ID, task.UserID).
		Updates(map[string]interface{}{
			"task":    task.Task,
			"is_done": task.IsDone,
		}).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint, userID uint) error {
	result := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	return nil
}
