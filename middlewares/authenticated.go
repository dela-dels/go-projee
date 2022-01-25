package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

//Here, we define an interface to aid in extracting the jwt token
//from the request header easily. The interface has one method that returns
//the token extracted as a jwt token or an error should the extraction process fail.
type ExtractToken interface {
	Extract(string) (*jwt.Token, error)
}

func Authenticated() gin.HandlerFunc {
	//var tokenExtractor ExtractToken

	return func(context *gin.Context) {
		log.Printf("token: %s", context.Request.Header.Get("Authorization"))
		//token, err := tokenExtractor.Extract(context.Request.Header.Get("Authorization"))
		//if err != nil || !token.Valid {
		//	context.AbortWithStatus(http.StatusUnauthorized)
		//}
		//
		//claims := token.Claims.(jwt.MapClaims)
		//email := claims["email"].(string)
		//
		//context.Set("email", email)
		context.Next()
	}
}