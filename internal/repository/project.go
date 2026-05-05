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

func (r *ProjectRepository) GetByID(id, userID uint) (*models.Project, error) {
	var project models.Project

	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&project, id).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) GetAll(userID uint) ([]models.Project, error) {
	var projects []models.Project

	err := r.db.Where("user_id = ?", userID).Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) Update(project *models.Project, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", project.ID, userID).Updates(project).Error
}

func (r *ProjectRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Project{}, id).Error
}
