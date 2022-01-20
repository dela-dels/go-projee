package models

import (
	"time"
)

type Project struct {
	BaseModel
	Name string `gorm:"not null"`
	Description string `gorm:"string"`
	StartDate time.Time `gorm:"not null"`
	EndDate time.Time `gorm:"not null"`
	IsCompleted bool `gorm:"index"`

	ProjectLead int
	User User`gorm:"foreignKey:ProjectLead;not null"`

	Sprint []Sprint
}