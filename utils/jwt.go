package utils

import "github.com/dela-dels/go-projee/models"

type Token struct {
	Email       string
	TokenString string
}

func CreateToken(user models.User) {}
