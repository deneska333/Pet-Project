package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)

	GetAllTasks() ([]Task, error)

	PatchTasks(id int, task Task) (Task, error)

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

func (r *taskRepository) PatchTasks(iD int, task Task) (Task, error) {
	var updatedTask Task
	err := r.db.Model(&updatedTask).Where("id = ?", iD).Updates(Task{Task: task.Task, IsDone: task.IsDone}).Error
	r.db.First(&updatedTask, iD)
	return updatedTask, err
}

func (r *taskRepository) DeleteTaskByID(iD int) error {

	err := r.db.Delete(&Task{}, iD).Error
	return err
}
