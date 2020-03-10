package main

import (
	"fmt"
	"net/http"
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/login", inDB.LoginHandler)
	router.POST("/register", Auth, inDB.RegisterUser)
	router.GET("/api/token", inDB.GetToken)
	router.Run(":3000")
}

func Auth(c *gin.Context) {
	var result gin.H
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){

		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		result = gin.H{
			"status": "success",
			"pesan": "token verified",
		}
		c.JSON(http.StatusOK, result)
		c.Abort()
	} else {
		result = gin.H {
			"status": "error " + err.Error(),
			"pesan": "not authorized",
			"token": token,
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}