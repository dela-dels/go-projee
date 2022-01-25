package utils

import (
	"os"
	"time"

	"github.com/dela-dels/go-projee/models"
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Email       string
	TokenString string
}

// type JWT struct {
// 	//this is the secret key for signing the jwt
// 	key []byte

// 	//jwt algorithm to use. Most likely to be HMAC SHA 256
// 	algorithm jwt.SigningMethod

// 	//Period for which the token is a valid one
// 	//ttl time.Duration
// }

//func (j JWT) New() JWT {
//	return JWT{
//		[]byte(os.Getenv("JWT_SECRET")),
//		jwt.GetSigningMethod("HS256"),
//	}
//}

func GenerateToken(user models.User) (string, error) {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"exp":        time.Now().Add(time.Minute * 15).Unix(),
		},
	).SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
