package controllers

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"
	"strconv"
)

var x string

func (idb *InDB) RegisterUser(c *gin.Context) {
	var (
		user structs.Login
		result gin.H
	)

	nip 					:= c.PostForm("nip")
	password 				:= c.PostForm("password")
	name 					:= c.PostForm("name")
	no_hp 					:= c.PostForm("no_hp")
	kode_subdirekotrat 		:= c.PostForm("kode_subdirektorat")
	kode_seksi 				:= c.PostForm("kode_seksi")
	aktif 					:= c.PostForm("aktif")
	level_pengguna 			:= c.PostForm("level_pengguna")

	user.Nip				= nip	
	//user.Password			= password
	user.Name				= name
	user.No_hp				= no_hp
	user.Kode_subdirekotrat	= kode_subdirekotrat
	user.Kode_seksi			= kode_seksi
	user.Aktif, _		    = strconv.Atoi(aktif)
	user.Level_pengguna		= level_pengguna

	pwd := GetPwd(password)
	EncryptPass := HashAndSalt(pwd)
	user.Password = EncryptPass

	if nip == "" || password == "" || name == "" || no_hp == "" || kode_subdirekotrat == "" || kode_seksi == "" || aktif == "" || level_pengguna == ""{
		result = gin.H{
			"status": "Warning",
			"pesan": "Silahkan lengkapi lembar isian",
			"data-user": user,
		}
	} else {
		NipIsExist := idb.CekNip(nip)

		if NipIsExist == true {
			result = gin.H {
				"status": "error",
				"pesan": "Nip sudah terdaftar",
			}
		} else {
			status := idb.DB.Create(&user).Error
			if status != nil {
				result = gin.H {
					"status": "error",
					"pesan": status,
				}
			} else {
				result = gin.H {
					"status": "Success",
					"pesan": "Data berhasil ditambahkan",
				}
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
