package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Project struct {
	BaseModel
	Name string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate time.Time `gorm:"not null" json:"end_date"`
	IsCompleted bool `gorm:"index" json:"is_completed"`

	ProjectLead int
	User User`gorm:"foreignKey:ProjectLead;not null"`

	Sprint []Sprint
}

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

type Task struct {
	BaseModel
	Name string `json:"name"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	IsCompleted bool `gorm:"index" json:"is_completed"`

	User []User `gorm:"many2many:user_tasks"`
}

type User struct {
	BaseModel
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"type:varchar(200);uniqueIndex;not null" json:"email"`
}

type UserTasks struct {
	UserID int `gorm:"primaryKey"`
	TaskID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}