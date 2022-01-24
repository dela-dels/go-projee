package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"type:varchar(200);uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null"`
}

//func (user *User) FindByEmail(email string) (*User, error) {
//
//	err := gorm.DB.Where("email = ?", email).First(&user).Error
//
//	if err != nil {
//		return user, nil
//	}
//
//	return nil, err
//}
