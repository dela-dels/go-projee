package models

type User struct {
	//gorm.Model
	BaseModel
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"type:varchar(200);uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null"`
}
