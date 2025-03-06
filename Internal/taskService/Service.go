package taskService

import (
	"Project/Internal/request"
)

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTask() ([]Task, error) {
	task, err := s.repo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(id int, task request.MessageRequest) (Task, error) {
	return s.repo.PatchTasks(id, task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTaskByID(id)
}
