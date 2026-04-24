package models

type Project struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`

	Tasks []Task `json:"tasks,omitempty", gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
