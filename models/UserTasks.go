package models

import (
	"gorm.io/gorm"
	"time"
)

type UserTasks struct {
	UserID int `gorm:"primaryKey"`
	TaskID int `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
