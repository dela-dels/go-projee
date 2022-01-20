package models

import "time"

type Task struct {
	BaseModel
	Name string `json:"name"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	IsCompleted bool `gorm:"index" json:"is_completed"`

	User []User `gorm:"many2many:user_tasks"`
}