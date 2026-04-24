package repository

import (
	"github.com/nseve/first-go-restapi/internal/models"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepository) GetByID(id uint) (*models.Project, error) {
	var project models.Project

	err := r.db.First(&project, id).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project

	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *ProjectRepository) Delete(id uint) error {
	return r.db.Delete(&models.Project{}, id).Error
}
