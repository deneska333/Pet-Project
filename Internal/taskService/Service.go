package taskService

import (
	"errors"
)

type Task struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

type TasksRepository interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id uint) (*Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(task Task) (*Task, error)
	DeleteTaskByID(id uint) error
}

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id uint) (*Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	if task.Text == "" {
		return Task{}, errors.New("task text cannot be empty")
	}
	return s.repo.CreateTask(task)
}

func (s *TaskService) UpdateTask(id uint, update Task) (*Task, error) {
	existing, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	if update.Text != "" {
		existing.Text = update.Text
	}

	existing.IsDone = update.IsDone

	return s.repo.UpdateTask(*existing)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
