package controllers

import (
	"net/http"
	"../structs"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) LoginHandler(c *gin.Context) {
	var (
		user structs.Login
		newUser structs.Login
		err = c.Bind(&user)
	)

	nip := c.PostForm("nip")
	password := c.PostForm("password") 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": http.StatusBadRequest,
			"pesan": "tidak bisa bind Structs",
		})
	}

	NipIsExist := idb.CekNip(nip)
	Pass := GetPwd(password) 
	//EncryptPass := HashAndSalt(Pass)

	idb.DB.Select("password").Where(" nip = ?", nip).Find(&user)
	passToString := user.Password
	 
	matchPass := ComparePasswords(passToString, Pass) 

	if NipIsExist == false {
		c.JSON(http.StatusUnauthorized, gin.H {
			"status": "warning",
			"pesan": "Nip Tidak Terdaftar",
			"passToString": Pass,
		})
	} else if matchPass == false {
		c.JSON(http.StatusUnauthorized, gin.H {
			"status": "warning",
			"pesan": "Password Salah",
		})
	} else if matchPass == true && NipIsExist == true {
		sign := jwt.New(jwt.GetSigningMethod("HS256")) // hs256 
		
		claims := sign.Claims.(jwt.MapClaims)
		claims["user"] = nip

		token, err := sign.SignedString([]byte("secret"))

		newUser.Token = token
		update := idb.DB.Model(&user).Where("nip = ?", nip).Updates(newUser).Error

		if update != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"pesan": err.Error(),
				"status": "error",
			})
			c.Abort()
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"pesan": err.Error(),
			})
			c.Abort()
		}
		
		c.JSON(http.StatusOK, gin.H {
			"token": token,
			"detail_user": claims,
			"status": "success",
			"pesan": "Login Berhasil",
		})
	}
}

func (idb *InDB) GetToken(c *gin.Context) {
	var user structs.Login
	idb.DB.Find(&user, 7)
	c.JSON(http.StatusOK, gin.H {
		"token": user.Token,
	})
}