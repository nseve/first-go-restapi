package service

import (
	"errors"

	"github.com/nseve/first-go-restapi/internal/models"
	"github.com/nseve/first-go-restapi/internal/repository"
)

type TaskService struct {
	taskRepo    *repository.TaskRepository
	projectRepo *repository.ProjectRepository
}

func NewTaskService(
	taskRepo *repository.TaskRepository,
	projectRepo *repository.ProjectRepository,
) *TaskService {
	return &TaskService{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
	}
}

func (s *TaskService) Create(projectID, userID uint, title string, duration int) (*models.Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if duration <= 0 {
		return nil, errors.New("duration must be positive")
	}

	_, err := s.projectRepo.GetByID(projectID, userID)
	if err != nil {
		return nil, errors.New("project not found")
	}

	task := &models.Task{Title: title, Duration: duration, ProjectID: projectID}
	if err := s.taskRepo.Create(task, userID); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetByID(id, userID uint) (*models.Task, error) {
	if id == 0 {
		return nil, errors.New("Invalid task id")
	}

	return s.taskRepo.GetByID(id, userID)
}

func (s *TaskService) GetByProjectID(projectID, userID uint) ([]models.Task, error) {
	if projectID == 0 {
		return nil, errors.New("invaild project id")
	}

	return s.taskRepo.GetByProjectID(projectID, userID)
}

func (s *TaskService) Update(id, userID uint, title string, duration int) (*models.Task, error) {
	if id == 0 {
		return nil, errors.New("invalid task id")
	}

	if title == "" {
		return nil, errors.New("title is required")
	}

	if duration <= 0 {
		return nil, errors.New("duration must be positive")
	}

	task, err := s.taskRepo.GetByID(id, userID)
	if err != nil {
		return nil, errors.New("task not found")
	}

	task.Title = title
	task.Duration = duration

	if err := s.taskRepo.Update(task, userID); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) Delete(id, userID uint) error {
	return s.taskRepo.Delete(id, userID)
}
