package repository

import (
	"github.com/nseve/first-go-restapi/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task, userID uint) error {
	var project models.Project

	err := r.db.
		Where("id = ? AND user_id = ?", task.ProjectID, userID).
		First(&project).Error

	if err != nil {
		return err
	}

	return r.db.Create(task).Error
}

func (r *TaskRepository) GetByID(id, userID uint) (*models.Task, error) {
	var task models.Task

	err := r.db.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.user_id = ?", id, userID).
		First(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetByProjectID(projectID, userID uint) ([]models.Task, error) {
	var tasks []models.Task

	err := r.db.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("project_id = ? AND projects.user_id = ?", projectID, userID).
		Find(&tasks).Error

	return tasks, err
}

func (r *TaskRepository) Update(updated *models.Task, userID uint) error {
	var task models.Task

	err := r.db.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.user_id = ?", updated.ID, userID).
		First(&task).Error

	if err != nil {
		return err
	}

	task.Title = updated.Title
	task.Duration = updated.Duration

	return r.db.Save(&task).Error
}

func (r *TaskRepository) Delete(id, userID uint) error {
	var task models.Task

	err := r.db.
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.user_id = ?", id, userID).
		First(&task).Error

	if err != nil {
		return err
	}

	return r.db.Delete(&task).Error
}
