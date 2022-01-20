package models

type User struct {
	BaseModel
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"type:varchar(200);uniqueIndex;not null" json:"email"`
}
