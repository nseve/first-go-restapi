package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
}
