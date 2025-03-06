package taskService

import (
	"Project/Internal/request"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)

	GetAllTasks() ([]Task, error)

	PatchTasks(id int, task request.MessageRequest) (Task, error)

	DeleteTaskByID(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}
func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}
func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) PatchTasks(iD int, msg request.MessageRequest) (Task, error) {
	var message Task
	err := r.db.Model(&message).Where("id = ?", iD).Updates(Task{Task: msg.Message}).Error
	r.db.First(&message, iD)
	return message, err
}

func (r *taskRepository) DeleteTaskByID(iD int) error {

	err := r.db.Delete(&Task{}, iD).Error
	return err
}
