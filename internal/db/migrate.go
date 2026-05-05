package db

import (
	"github.com/nseve/first-go-restapi/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Project{},
		&models.Task{},
		&models.User{},
	)
}
