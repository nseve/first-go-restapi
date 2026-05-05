package service

import (
	"errors"

	"github.com/nseve/first-go-restapi/internal/models"
	"github.com/nseve/first-go-restapi/internal/repository"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) Create(title string, userID uint) (*models.Project, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	project := &models.Project{Title: title, UserID: userID}

	if err := s.repo.Create(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetByID(id, userID uint) (*models.Project, error) {
	return s.repo.GetByID(id, userID)
}

func (s *ProjectService) GetAll(userID uint) ([]models.Project, error) {
	return s.repo.GetAll(userID)
}

func (s *ProjectService) Update(id, userID uint, title string) (*models.Project, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	project, err := s.repo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	project.Title = title

	if err := s.repo.Update(project, userID); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) Delete(id, userID uint) error {
	return s.repo.Delete(id, userID)
}
