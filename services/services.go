package services

import (
	"task-service/models"
	"task-service/repositories"
)

type TaskService struct {
	Repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) GetTasks(status string, page, limit int) ([]models.Task, error) {
	return s.Repo.GetAll(status, page, limit)
}

func (s *TaskService) GetTask(id uint) (models.Task, error) {
	return s.Repo.GetByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.Repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.Repo.Delete(id)
}
