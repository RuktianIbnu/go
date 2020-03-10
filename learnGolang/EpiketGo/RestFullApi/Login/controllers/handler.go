package controllers

import (
    "../structs"
    "log"
    "fmt"
    "net/http"

    "golang.org/x/crypto/bcrypt"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) CekNip(nip string) bool{
	var user structs.Login
	isExist :=	idb.DB.Where("nip = ?", nip).First(&user).Error
	if isExist == nil {
		return true
	} else {
		return false
	}
}

func GetPwd(pass string) []byte {
    return []byte(pass)
}

func HashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func ComparePasswords(PassFromDb string, EncryptPassInput []byte) bool {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(PassFromDb)
    err := bcrypt.CompareHashAndPassword(byteHash, EncryptPassInput)
    if err != nil {
        log.Println(err)
        return false
    }
    
    return true
}

func Auth(c *gin.Context) {
	var result gin.H
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("hS256") != token.Method {
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
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}