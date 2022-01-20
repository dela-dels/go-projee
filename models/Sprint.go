package models

import "time"

type Sprint struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate time.Time `gorm:"not null" json:"end_date"`
	IsCompleted bool `gorm:"index" json:"is_completed"`

	ProjectID int
	Project Project

	SprintLead int
	User User `gorm:"foreignKey:SprintLead;not null"`
}