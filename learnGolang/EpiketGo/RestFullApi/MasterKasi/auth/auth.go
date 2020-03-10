package auth

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

func Auth(c *gin.Context) bool{
	var status bool
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		status = true
	} else {
		status = false
	}
	return status
}