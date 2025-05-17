package taskService

import "gorm.io/gorm"

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return TaskService{repo: repo}
}

func (s *TaskService) CreateTask(taskText string, userID uint) (Task, error) {
	task := Task{
		Task:   taskText,
		UserID: userID,
	}
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks(userID uint) ([]Task, error) {
	return s.repo.GetAllTasks(userID)
}

func (s *TaskService) GetTaskByID(id, userID uint) (*Task, error) {
	return s.repo.GetTaskByID(id, userID)
}

func (s *TaskService) UpdateTask(id, userID uint, taskText string, isDone bool) (*Task, error) {
	task := Task{
		Model:  gorm.Model{ID: id},
		Task:   taskText,
		IsDone: isDone,
		UserID: userID,
	}
	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id, userID uint) error {
	return s.repo.DeleteTaskByID(id, userID)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]Task, error) {
	return s.repo.GetAllTasks(userID)
}
