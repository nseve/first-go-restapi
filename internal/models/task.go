package models

type Task struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null"`
	Duration  int    `json:"duration" gorm:"not null"`
	ProjectID uint   `json:"project_id" gorm:"index;not null"`
}
