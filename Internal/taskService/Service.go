package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	createdTask, err := s.repo.CreateTask(task)
	return createdTask, err
}

func (s *TaskService) UpdateTask(id uint, update Task) (*Task, error) {
	existing, err := s.repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	existing.Text = update.Text
	existing.IsDone = update.IsDone
	return s.repo.UpdateTask(*existing)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
